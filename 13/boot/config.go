package boot

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

func ConfigSetup() {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).AddPath("/manifest/config")
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.yaml")
}
