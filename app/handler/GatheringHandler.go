package handler

import (
	"gits-test/models/request"

	"github.com/gin-gonic/gin"
)

func (ah *AppHandler) CreateGatheringHandler(c *gin.Context) {
	var createGatheringRequest request.CreateGatheringRequest
	if err := c.BindJSON(&createGatheringRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if err := ah.AppUsecase.CreateGatheringUsecase(createGatheringRequest); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

func (ah *AppHandler) ListGatheringHandler(c *gin.Context) {
	listGathering, err := ah.AppUsecase.ListGatheringUsecase()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success List Gathering",
		"data":    listGathering,
	})
}
