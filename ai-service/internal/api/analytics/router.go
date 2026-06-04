package analytics

import (
	"diaxel/internal/app"

	"github.com/gin-gonic/gin"
)

func AnalyticsRoutes(r *gin.Engine, application *app.App) {
	api := r.Group("/analytics")
	{
		api.GET("", GetAnalytics(application))
	}
}
