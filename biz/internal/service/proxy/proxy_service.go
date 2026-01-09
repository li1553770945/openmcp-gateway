package proxy

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/li1553770945/openmcp-gateway/biz/infra/cache"
	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
	"github.com/li1553770945/openmcp-gateway/biz/internal/repo/mcpserver"
)

type ProxyServiceImpl struct {
	mcpRepo    mcpserver.IMCPServerRepository
	client     *client.Client
	cfg        *config.ProxyConfig
	proxyCache cache.IProxyCache
}

func NewProxyService(repo mcpserver.IMCPServerRepository, proxyCache cache.IProxyCache) *ProxyServiceImpl {
	c, _ := client.NewClient(
		client.WithMaxIdleConnDuration(30*time.Second),
		client.WithMaxConnsPerHost(1000),
		client.WithDialTimeout(5*time.Second),
	)
	return &ProxyServiceImpl{
		mcpRepo:    repo,
		client:     c,
		proxyCache: proxyCache,
	}
}
func (s *ProxyServiceImpl) resolveTarget(token string) (string, error) {
	// 1. 尝试从缓存获取
	if target, ok := s.proxyCache.GetTargetBaseUrl(token); ok {
		return target, nil
	}

	// 2. 缓存未命中，查询数据库
	tokenEntity, err := s.mcpRepo.FindTokenByToken(token)
	if err != nil {
		return "", err
	}
	if tokenEntity == nil {
		return "", errors.New("token不存在") // Token 不存在
	}
	mcpServer, err := s.mcpRepo.FindMCPServerById(tokenEntity.MCPServerID)
	if err != nil {
		return "", err
	}
	if mcpServer == nil {
		return "", errors.New("关联的MCPServer不存在") // 关联的 MCPServer 不存在
	}
	if mcpServer.OpenProxy == false {
		return "", errors.New("该MCPServer未开启代理功能") // MCPServer 未开启代理功能
	}

	// 3. 更新缓存
	s.proxyCache.SetTargetBaseUrl(token, mcpServer.Url)

	return mcpServer.Url, nil
}
func (s *ProxyServiceImpl) ForwardRequest(ctx context.Context, c *app.RequestContext) {
	// 1. 获取 Token (Header优先, 其次Query)
	token := string(c.GetHeader("X-OpenMCP-Token"))

	if token == "" {
		c.String(http.StatusUnauthorized, "Missing Token")
		return
	}

	// 2. 解析目标地址
	targetBaseURL, err := s.resolveTarget(token)
	if err != nil {
		hlog.CtxErrorf(ctx, "Resolve token failed: %v", err)
		c.String(http.StatusBadRequest, "解析URL失败: "+err.Error())
		return
	}

	if !strings.HasPrefix(targetBaseURL, "http://") && !strings.HasPrefix(targetBaseURL, "https://") {
		targetBaseURL = "http://" + targetBaseURL
	}

	// 3. 构造转发请求
	req := &protocol.Request{}
	res := &protocol.Response{}

	// 复制原始请求 (Header, Body, Method)
	c.Request.CopyTo(req)

	// URL 重写逻辑
	// 假设注册路由是 /proxy/*path
	// 原始: /proxy/v1/chat/completions -> 目标: {targetBaseURL}/v1/chat/completions
	requestPath := string(c.Request.URI().Path())

	// 去除 /proxy 前缀
	finalPath := requestPath
	if strings.HasPrefix(requestPath, "/proxy") {
		finalPath = strings.TrimPrefix(requestPath, "/proxy")
	}

	// 处理 targetBaseURL 可能带 path 的情况 (如 http://example.com/mcp)
	// 简单的拼接逻辑：TrimRight(base, "/") + "/" + TrimLeft(path, "/")
	target := strings.TrimRight(targetBaseURL, "/")
	relativePath := strings.TrimLeft(finalPath, "/")
	if relativePath != "" {
		target = target + "/" + relativePath
	}

	// 同时保留 QueryString
	queryString := string(c.Request.URI().QueryString())
	if queryString != "" {
		target += "?" + queryString
	}
	hlog.CtxInfof(ctx, "Forwarding request to: %s", target)

	req.SetRequestURI(target)

	// 必须重置 Host 头，否则反向代理可能失败
	req.Header.SetHostBytes(req.URI().Host())

	// 4. 发送请求
	err = s.client.Do(ctx, req, res)
	if err != nil {
		hlog.CtxErrorf(ctx, "Forward request failed: %v", err)
		c.String(http.StatusBadGateway, "Upstream Error")
		return
	}

	// 5. 写回响应
	res.CopyTo(&c.Response)
}
