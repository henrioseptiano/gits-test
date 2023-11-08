package handler

import (
	"gits-test/models/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ah *AppHandler) CreateInvitationHandler(c *gin.Context) {
	var request request.CreateInvitationRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ah.AppUsecase.CreateInvitationUsecase(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ah *AppHandler) AcceptNotAcceptHandler(c *gin.Context) {
	invitationID := c.Param("invitation_id")
	var request request.AcceptNotAcceptRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ah.AppUsecase.AcceptNotAcceptUsecase(request, invitationID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ah *AppHandler) ListInvitationHandler(c *gin.Context) {
	listInvitationResponse, err := ah.AppUsecase.ListInvitationUsecase()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success List Data Member",
		"data":    listInvitationResponse,
	})
	return
}
