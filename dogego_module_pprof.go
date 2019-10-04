package dogego_module_pprof

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"os"
)

func NewPprofModule(router *gin.Engine) {
	if gin.Mode() != gin.ReleaseMode {
		usePprof(router)
		return
	} else {
		if os.Getenv("ENABLE_PPROF") == "enable" {
			usePprof(router)
			return
		}
	}
}

func usePprof(router *gin.Engine) {
	pprof.Register(router, "/debug/pprof")
}
