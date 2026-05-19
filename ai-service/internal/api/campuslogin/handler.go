package campuslogin

import (
	"bytes"
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CampusLoginHandler struct {
	cfg *config.Settings
	db  *db.Client
}

func NewCampusLoginHandler(cfg *config.Settings, db *db.Client) *CampusLoginHandler {
	return &CampusLoginHandler{cfg: cfg, db: db}
}

func (h *CampusLoginHandler) HandleTest(c *gin.Context) {
	// Read request body
	var bodyBytes []byte
	if c.Request.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("[CAMPUSLOGIN TEST] Failed to read request body: %v", err)
		} else {
			// Restore the body so it can be read again if needed
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	// Format headers for logging
	headersLog := ""
	for name, values := range c.Request.Header {
		for _, value := range values {
			headersLog += "  " + name + ": " + value + "\n"
		}
	}

	// Format body for logging
	bodyLog := ""
	if len(bodyBytes) > 0 {
		// Try to pretty-print JSON
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, bodyBytes, "", "  "); err == nil {
			bodyLog = prettyJSON.String()
		} else {
			bodyLog = string(bodyBytes)
		}
	} else {
		bodyLog = "<empty body>"
	}

	// Log all request details clearly
	log.Printf(`
============================================================
[CAMPUSLOGIN WEBHOOK TEST] New Request Received
------------------------------------------------------------
Method:      %s
URI:         %s
Client IP:   %s
Headers:
%s
Query Params: %s
Raw Body Length: %d bytes
Body Content:
%s
============================================================`,
		c.Request.Method,
		c.Request.RequestURI,
		c.ClientIP(),
		headersLog,
		c.Request.URL.RawQuery,
		len(bodyBytes),
		bodyLog,
	)

	// Return a 200 OK status with the incoming data in response for easy debugging
	// We can try to return the body parse structure or just raw string/json
	var jsonBody interface{}
	if err := json.Unmarshal(bodyBytes, &jsonBody); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Data received and logged successfully",
			"received_data": gin.H{
				"method":       c.Request.Method,
				"uri":          c.Request.RequestURI,
				"headers":      c.Request.Header,
				"query_params": c.Request.URL.Query(),
				"body":         jsonBody,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Data received and logged successfully",
			"received_data": gin.H{
				"method":       c.Request.Method,
				"uri":          c.Request.RequestURI,
				"headers":      c.Request.Header,
				"query_params": c.Request.URL.Query(),
				"body":         string(bodyBytes),
			},
		})
	}
}
