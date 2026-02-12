package analytics

import (
	analytics2 "diaxel/internal/modules/analytics"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, analyticsService *analytics2.AnalyticsService) {
	api := r.Group("/api/analytics")
	{
		api.GET("/metrics", func(c *gin.Context) {
			filter := analytics2.AnalyticsFilter{
				AssistantID: c.Query("assistant_id"),
				Platform:    c.Query("platform"),
			}

			if startDateStr := c.Query("start_date"); startDateStr != "" {
				if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
					filter.StartDate = startDate
				}
			}

			if endDateStr := c.Query("end_date"); endDateStr != "" {
				if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
					filter.EndDate = endDate
				}
			}

			result, err := analyticsService.GetAnalytics(c.Request.Context(), filter)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, result)
		})

		api.GET("/metrics/assistant/:assistant_id", func(c *gin.Context) {
			assistantID := c.Param("assistant_id")

			filter := analytics2.AnalyticsFilter{
				Platform: c.Query("platform"),
			}

			if startDateStr := c.Query("start_date"); startDateStr != "" {
				if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
					filter.StartDate = startDate
				}
			}

			if endDateStr := c.Query("end_date"); endDateStr != "" {
				if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
					filter.EndDate = endDate
				}
			}

			result, err := analyticsService.GetAnalyticsByAssistant(c.Request.Context(), assistantID, filter)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, result)
		})

		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})
	}
}
