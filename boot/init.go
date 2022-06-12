package boot

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

func init() {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).AddPath("/manifest/config")
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.yaml")
}
