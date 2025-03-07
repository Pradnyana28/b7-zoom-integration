package controllers

import (
	"b7-zoom-integration/apps/b7-zoom-integration-apis/internal/domains"

	"github.com/gin-gonic/gin"
)

// ListMeetings list the created meetings
func ListMeetings(c *gin.Context) {
	results := domains.ListMeetings()
	c.JSON(200, results)
}

// CreateMeeting creates a new meeting
func CreateMeeting(c *gin.Context) {
	var input *domains.CreateMeetingInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := domains.CreateMeeting(input)
	if err != nil {
		c.JSON(412, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Meeting created successfully"})
}

// AddMeetingsRoutes adds the meetings routes to the router
func AddMeetingsRoutes(router *gin.Engine) {
	meetings := router.Group("/v1/meetings")
	{
		meetings.GET("/", ListMeetings)
		meetings.POST("/", CreateMeeting)
	}
}
