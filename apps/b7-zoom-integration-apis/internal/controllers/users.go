package controllers

import (
	"b7-zoom-integration/apps/b7-zoom-integration-apis/internal/domains"

	"github.com/gin-gonic/gin"
)

// CurrentUser returns the current user
func CurrentUser(c *gin.Context) {
	results, err := domains.Me()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, results)
}

// AddUsersRoutes adds the users routes to the router
func AddUsersRoutes(router *gin.Engine) {
	users := router.Group("/v1/users")
	{
		users.GET("/me", CurrentUser)
	}
}
