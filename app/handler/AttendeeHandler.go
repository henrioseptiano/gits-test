package handler

import "github.com/gin-gonic/gin"

func (ah *AppHandler) ListAttendeeHandler(c *gin.Context) {
	listAttendeeResponse, err := ah.AppUsecase.ListAttendeeUsecase()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success List Data Attendee",
		"data":    listAttendeeResponse,
	})
	return
}
