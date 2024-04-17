package server

import "github.com/gin-gonic/gin"

func SetEndpoints(r *gin.Engine) {
	r.POST("/signup")
	r.POST("/signin", HandleSignIn)
}
