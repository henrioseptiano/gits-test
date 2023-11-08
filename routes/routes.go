package routes

import (
	"gits-test/app"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handler app.AppHandlerInterface
}

func NewRouter(handler app.AppHandlerInterface) *Router {
	return &Router{
		Handler: handler,
	}
}

func (r *Router) AppRoutes(ginRouter *gin.Engine) *gin.Engine {
	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	groupRouter := ginRouter.Group("/api/v1")
	{
		// Member
		groupRouter.POST("/member", r.Handler.CreateMemberHandler)
		groupRouter.GET("/member", r.Handler.ListMemberHandler)

		// Gathering
		groupRouter.POST("/gathering", r.Handler.CreateGatheringHandler)
		groupRouter.GET("/gathering", r.Handler.ListGatheringHandler)

		// Invitation
		groupRouter.POST("/invitation", r.Handler.CreateInvitationHandler)
		groupRouter.PUT("/acceptnotaccept/:invitation_id", r.Handler.AcceptNotAcceptHandler)
		groupRouter.GET("/invitation", r.Handler.ListInvitationHandler)

		// Attendee
		groupRouter.GET("/attendee", r.Handler.ListAttendeeHandler)
	}

	return ginRouter
}
