package boot

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func ORMSetup() {
	ctx := gctx.New()
	err := g.DB().PingMaster()
	if err != nil {
		g.Log().Fatalf(ctx, "Connect to Mysql server failed, err: %v", err)
	} else {
		g.Log().Info(ctx, "Initialize MySQL server successfully")
	}
}
