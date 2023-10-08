package main

import (
	"github.com/gin-gonic/gin"
	"go-jwt/controllers/auth"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", auth.Register)
	public.POST("/login", auth.Login)

	err := r.Run(":8081")

	if err != nil {
		return
	}

}
