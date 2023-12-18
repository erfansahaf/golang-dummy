package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"personal/sika/internal/database/models"
)

func GetUser(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	rows, err := db.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong! (It's us, not you!)",
			})
			return
		}
	}
	if user.ID == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "User not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"user": user,
		},
	})
	return
}
