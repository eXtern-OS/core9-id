package server

import (
	"errors"
	"github.com/eXtern-OS/core9-common/db"
	"github.com/eXtern-OS/core9-common/models/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func HandleSignUp(c *gin.Context) {
	var u user.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var u2 user.User
	err := db.DefaultClient.FindOne(bson.M{"login": u.Login}, &u2, "users", "users")
	if !errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login already exists"})
		return
	}
	err = db.DefaultClient.FindOne(bson.M{"email": u.Email}, &u2, "users", "users")
	if !errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already in user"})
		return
	}

}
