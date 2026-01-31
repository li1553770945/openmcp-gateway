package proxy

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type IProxyService interface {
	ForwardRequest(ctx context.Context, c *app.RequestContext)
}
