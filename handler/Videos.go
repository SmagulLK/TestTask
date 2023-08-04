package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVideos(c *gin.Context) {
	arrayOfBooks, err := h.Service.GetVideos()
	fmt.Println(arrayOfBooks)
	if err != nil {
		newErrorResponse(c, err.Error(), 500)
		return
	}
	c.JSON(200, gin.H{
		"message": "",
		"data":    arrayOfBooks,
	})
}
