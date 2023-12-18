package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"message": "Something went wrong! (It's us, not you!)",
	})
}
