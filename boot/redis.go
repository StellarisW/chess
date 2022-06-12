package boot

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func RedisSetup() {
	ctx := gctx.New()
	_, err := g.Redis().Do(ctx, "SET", "test", "test")
	if err != nil {
		g.Log().Fatalf(gctx.New(), "Connect to Redis server failed, err: %v\n", err)
	}
	g.Redis().Do(ctx, "EXPIRE", "test", "30")
	g.Log().Info(ctx, "Initialize Redis instance successfully")
}
