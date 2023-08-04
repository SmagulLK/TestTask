package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, message string, status int) {
	log.Fatalf("Error of connecting to DB: %s \n", message)
	c.AbortWithStatusJSON(status, error{message})
}
