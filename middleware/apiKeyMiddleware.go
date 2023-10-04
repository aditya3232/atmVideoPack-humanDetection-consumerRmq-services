package middleware

import (
	"net/http"

	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/log"
	"github.com/gin-gonic/gin"
)

// get api key from config
// compare api key from config with api key from request header
// if not match, return unauthorized
// if match, continue

func ApiKeyMiddleware(validAPIKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header from request
		authKeyHeader := c.GetHeader("x-api-key")

		apiKey := authKeyHeader
		if apiKey != validAPIKey {
			errors := helper.FormatErrorWithCustomMessage("Invalid API Key")
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			log.Info("AuthAttemptFail", " from ip address: ", c.ClientIP())
			return
		}

		// log percobaan auth user berserta informasi ip address and timestamp
		// log.Info("AuthAttempt", " from ip address: ", c.ClientIP(), " at ", time.Now())

		// Call next middleware/handler function in the chain
		c.Next()
	}
}
