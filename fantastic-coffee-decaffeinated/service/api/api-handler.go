package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// get the logged user stream
	rt.router.GET("/my-stream", rt.getMyStream)
	// loggin
	rt.router.POST("/session", rt.doLogin)
	// Get searched usernames list
	rt.router.GET("/profiles", rt.searchUsers)
	// get a user profile
	rt.router.GET("/profiles/:username", rt.getProfile)
	// get a post
	rt.router.GET("/posts/:idPhoto", rt.getPost)
	// set a new username
	rt.router.PUT("/profiles/:username", rt.setMyUsername)
	// get a photo
	rt.router.GET("/photos/:idPhoto", rt.getPhoto)
	// Upload a photo
	rt.router.POST("/photos", rt.uploadPhoto)
	// Delete a photo
	rt.router.DELETE("/photos/:idPhoto", rt.deletePhoto)
	// Follow one user
	rt.router.POST("/follows", rt.followUser)
	// unFollow one user
	rt.router.DELETE("/follows/:username", rt.unfollowUser)
	// Like one user's photo
	rt.router.POST("/photos/:idPhoto/likes", rt.likePhoto)
	// Remove a like
	rt.router.DELETE("/photos/:idPhoto/likes/:username", rt.removeLike)
	// Get the comments from a photoID
	rt.router.GET("/photos/:idPhoto/comments", rt.getComments)
	// Comment a photo
	rt.router.POST("/photos/:idPhoto/comments", rt.commentPhoto)
	// Remove a comment
	rt.router.DELETE("/photos/:idPhoto/comments/:idComment", rt.removeComment)
	// ban a user
	rt.router.POST("/bans", rt.banUser)
	// unban a user
	rt.router.DELETE("/bans/:username", rt.unbanUser)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
