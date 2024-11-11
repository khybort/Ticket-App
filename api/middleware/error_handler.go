package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			statusCode := c.Writer.Status()

			log.Printf("Error: %d | %s | %s | %s", statusCode, c.Request.Method, c.Request.URL.Path, c.Request.UserAgent())

			var message string
			switch statusCode {
			case http.StatusNotFound:
				message = "The requested resource was not found."
			case http.StatusInternalServerError:
				message = "An internal server error occurred."
			default:
				message = "An error occurred while processing the request."
			}

			response := JSONResponse{
				Error:   http.StatusText(statusCode),
				Message: message,
			}
			c.JSON(statusCode, response)
			c.Abort()
		}
	}
}
