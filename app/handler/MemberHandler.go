package handler

import (
	"gits-test/models/request"

	"github.com/gin-gonic/gin"
)

func (ah *AppHandler) CreateMemberHandler(c *gin.Context) {
	var createMemberRequest request.CreateMemberRequest
	if err := c.BindJSON(&createMemberRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if err := ah.AppUsecase.CreateMemberUsecase(createMemberRequest); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Create Member Success!",
	})
	return
}

func (ah *AppHandler) ListMemberHandler(c *gin.Context) {
	listMemberResponse, err := ah.AppUsecase.ListMemberUsecase()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success List Data Member",
		"data":    listMemberResponse,
	})
	return
}
