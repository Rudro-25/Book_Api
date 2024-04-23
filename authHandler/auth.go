package authHandler

import (
	"github.com/Rudro-25/book_api_server/dataHandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func authenticate(username, password string) bool {
	for _, user := range dataHandler.Users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func LoginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if authenticate(loginData.Username, loginData.Password) {
		// In a real-world scenario, you would generate a JWT token here
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}
