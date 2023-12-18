package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"personal/sika/internal/database/models"
	"strconv"
)

const limit int64 = 5

func GetUsers(c *gin.Context, db *sql.DB) {
	var page, offset int64
	page = getPage(c)
	offset = (page - 1) * limit
	rows, err := db.Query("SELECT id, first_name, last_name FROM users ORDER BY id DESC LIMIT ? OFFSET ? ;", limit, offset)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
			log.Fatal(err)
			ServerError(c)
			return
		}
		users = append(users, user)
	}

	cnt := len(users)
	pagination := gin.H{
		"current":   page,
		"count":     cnt,
		"next_page": nil,
	}
	if int64(cnt) == limit {
		pagination["next_page"] = page + 1
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"data":       users,
		"pagination": pagination,
	})
	return
}

func getPage(c *gin.Context) int64 {
	var page int64
	query := c.Request.URL.Query()
	val, ok := query["page"]
	if ok {
		i, err := strconv.ParseInt(val[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		} else {
			page = i
		}
	}
	if page < 1 {
		page = 1
	}

	return page
}
