package controllers

import (
	"github.com/imsujan276/untold-story/api/middlewares"
)

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		// Login Route
		v1.POST("/login", s.Login)

		// Reset password:
		v1.POST("/password/forgot", s.ForgotPassword)
		v1.POST("/password/reset", s.ResetPassword)

		//Users routes
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetUsers)
		v1.GET("/users/:id", s.GetUser)
		v1.PUT("/users/:id", middlewares.AuthTokenMiddleware(), s.UpdateUser)
		v1.PUT("/avatar/users/:id", middlewares.AuthTokenMiddleware(), s.UpdateAvatar)
		v1.DELETE("/users/:id", middlewares.AuthTokenMiddleware(), s.DeleteUser)

		//Posts routes
		v1.POST("/posts", middlewares.AuthTokenMiddleware(), s.CreatePost)
		v1.GET("/posts", s.GetPosts)
		v1.GET("/posts/:id", s.GetPost)
		v1.PUT("/posts/:id", middlewares.AuthTokenMiddleware(), s.UpdatePost)
		v1.DELETE("/posts/:id", middlewares.AuthTokenMiddleware(), s.DeletePost)
		v1.GET("/user_posts/:id", s.GetUserPosts)

		//Like route
		v1.GET("/likes/:id", s.GetLikes)
		v1.POST("/likes/:id", middlewares.AuthTokenMiddleware(), s.LikePost)
		v1.DELETE("/likes/:id", middlewares.AuthTokenMiddleware(), s.UnLikePost)

		//Comment routes
		v1.POST("/comments/:id", middlewares.AuthTokenMiddleware(), s.CreateComment)
		v1.GET("/comments/:id", s.GetComments)
		v1.PUT("/comments/:id", middlewares.AuthTokenMiddleware(), s.UpdateComment)
		v1.DELETE("/comments/:id", middlewares.AuthTokenMiddleware(), s.DeleteComment)
	}
}
