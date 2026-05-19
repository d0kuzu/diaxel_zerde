package campuslogin

import (
	appModule "diaxel/internal/app"

	"github.com/gin-gonic/gin"
)

func CampusLoginRoutes(router *gin.Engine, app *appModule.App) {
	h := NewCampusLoginHandler(app.Cfg, app.Db)

	group := router.Group("campuslogin")
	{
		group.Any("/test", h.HandleTest)
	}
}
