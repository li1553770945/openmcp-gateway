package cache

import (
	"sync"
	"time"

	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
)

type CacheTarget struct {
	targetBaseUrl string
	expiry        time.Time
}

type ProxyCache struct {
	cache sync.Map // map[string]CacheTarget
	cfg   *config.ProxyConfig
}

func NewProxyCache(cfg *config.ProxyConfig) IProxyCache {
	return &ProxyCache{
		cache: sync.Map{},
		cfg:   cfg,
	}
}

func (pc *ProxyCache) GetTargetBaseUrl(key string) (string, bool) {
	value, ok := pc.cache.Load(key)
	if !ok {
		return "", false
	}
	cacheTarget := value.(CacheTarget)
	if time.Now().After(cacheTarget.expiry) {
		// 已过期
		pc.cache.Delete(key)
		return "", false
	}
	return cacheTarget.targetBaseUrl, true
}

func (pc *ProxyCache) SetTargetBaseUrl(key string, targetBaseUrl string) {
	expiration := time.Now().Add(time.Duration(pc.cfg.CacheExpirationSeconds) * time.Second)
	pc.cache.Store(key, CacheTarget{
		targetBaseUrl: targetBaseUrl,
		expiry:        expiration,
	})
}

func (pc *ProxyCache) InvalidateByToken(token string) error {
	pc.cache.Delete(token)
	return nil
}
