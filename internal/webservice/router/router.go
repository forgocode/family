package router

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/webservice/middleware"
	"github.com/forgocode/family/internal/webservice/router/manager"
	"github.com/forgocode/family/pkg/bininfo"
)

func Start() {
	gin.SetMode(bininfo.GinLogMode)
	bininfo.StartTime = time.Now().UnixMilli()
	engine := gin.New()
	engine.Use(middleware.Logger(), middleware.Recovery())

	manager.RegisterRouter(engine)

	engine.Run(":8800")
}
