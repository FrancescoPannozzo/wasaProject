package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.POST("/session", rt.doLogin)

	rt.router.POST("/profiles/:username", rt.setMyUsername)

	//rt.router.POST("/profiles/:username/photos/:idPhoto", rt.uploadPhoto)
	rt.router.POST("/profiles/:username/photos", rt.uploadPhoto)
	rt.router.DELETE("/profiles/:username/photos/:idPhoto", rt.deletePhoto)
	// Follow one user
	rt.router.POST("/profiles/:username/follows", rt.followUser)
	// unFollow one user
	rt.router.DELETE("/profiles/:username/follows", rt.unfollowUser)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
