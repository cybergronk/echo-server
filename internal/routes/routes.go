package routes

import (
	"net/http"

	"github.com/cybergronk/echo-server/internal/domain"
	"github.com/gin-gonic/gin"
)

// RedirectToHealth redirects home requests to health.
func RedirectToHealth(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/health")
}

// ReturnAliveStatus returns a good message if the server is up.
func ReturnAliveStatus(c *gin.Context) {
	c.String(http.StatusOK, "I live!")
}

// EchoBody responds to requests with desired response.
func EchoBody(c *gin.Context) {
	var (
		what       string
		statusCode int
	)
	{
		var echo *domain.Echo
		if err := c.ShouldBindBodyWithJSON(&echo); err != nil {
			what = "Oops! Incoming request had invalid JSON payload"
			statusCode = http.StatusBadRequest
		}
	}

	c.String(statusCode, what)
}
