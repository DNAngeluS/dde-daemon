package desktop

import (
	"pkg.deepin.io/dde/daemon/loader"
)

func init() {
	loader.Register(NewDaemon())
}