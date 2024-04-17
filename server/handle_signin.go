package server

import (
	"github.com/eXtern-OS/core9-common/db"
	"github.com/eXtern-OS/core9-common/models/user"
	"github.com/eXtern-OS/core9-common/utils"
	"github.com/eXtern-OS/core9-id/tokens"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type SignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func HandleSignIn(c *gin.Context) {
	var signin SignIn
	if err := c.BindJSON(&signin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash := utils.SHA256(signin.Password)
	var u user.User
	err := db.DefaultClient.FindOne(bson.M{"login": signin.Login}, &u, "users", "users")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if u.PasswordHashed != hash {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
		return
	}
	access, refresh, err := tokens.NewToken(u.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": access.T, "refresh_token": refresh.T})
}
