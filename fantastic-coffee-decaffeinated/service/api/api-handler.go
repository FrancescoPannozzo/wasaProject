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
	// Like one user's photo
	rt.router.POST("/photos/:idPhoto/likes", rt.likePhoto)
	// Remove a like
	rt.router.DELETE("/photos/:idPhoto/likes", rt.removeLike)
	// Comment a photo
	rt.router.POST("/photos/:idPhoto/comments", rt.commentPhoto)
	// Remove a comment
	rt.router.DELETE("/photos/:idPhoto/comments/:idComment", rt.removeComment)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
