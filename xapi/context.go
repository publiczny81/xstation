package xapi

import (
	"github.com/publiczny81/xstation/xapi/model"
	"golang.org/x/net/context"
	"sync"
)

type Context struct {
	connection
	*contextData
}

type contextKey struct{}

type contextData struct {
	mu              sync.RWMutex
	streamSessionId string
}

func (c *contextData) SetStreamSessionId(streamSessionId string) {
	c.mu.Lock()
	c.streamSessionId = streamSessionId
	c.mu.Unlock()
}

func (c *contextData) GetStreamSessionId() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.streamSessionId
}

func WithContext(ctx context.Context, appCtx Context) context.Context {
	return context.WithValue(ctx, contextKey{}, appCtx)
}

func FromContext(ctx context.Context) (appCtx Context, ok bool) {
	appCtx, ok = ctx.Value(contextKey{}).(Context)
	return
}

func (c Context) Receive(response any) (err error) {
	if err = c.connection.Receive(response); err != nil {
		return
	}
	switch v := response.(type) {
	case nil:
	case *model.LoginResponse:
		if v.Status {
			c.SetStreamSessionId(v.StreamSessionId)
		}
	case *model.LogoutResponse:
		if v.Status {
			c.SetStreamSessionId("")
		}
	default:
	}
	return
}
