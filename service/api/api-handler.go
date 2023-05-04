package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	// loggin
	rt.router.POST("/session", rt.doLogin)
	// Get searched usernames list
	rt.router.GET("/profiles", rt.searchUsers)
	// get a user profile
	rt.router.GET("/profiles/:username", rt.getProfile)
	// set a new username
	rt.router.PUT("/profiles/:username", rt.setMyUsername)
	// get the logged user stream
	rt.router.GET("/my-stream", rt.getMyStream)
	// get a post
	rt.router.GET("/posts/:idPhoto", rt.getPost)
	// get a photo
	rt.router.GET("/photos/:idPhoto", rt.getPhoto)
	// Delete a photo
	rt.router.DELETE("/photos/:idPhoto", rt.deletePhoto)
	// Upload a photo
	rt.router.POST("/photos", rt.uploadPhoto)
	// Follow one user
	rt.router.POST("/follows", rt.followUser)
	// unFollow one user
	rt.router.DELETE("/follows/:username", rt.unfollowUser)
	// Like one user's photo
	rt.router.POST("/photos/:idPhoto/likes", rt.likePhoto)
	// Remove a like
	rt.router.DELETE("/photos/:idPhoto/likes/:username", rt.removeLike)
	// Comment a photo
	rt.router.POST("/photos/:idPhoto/comments", rt.commentPhoto)
	// Get the comments from a photoID
	rt.router.GET("/photos/:idPhoto/comments", rt.getComments)
	// Remove a comment
	rt.router.DELETE("/photos/:idPhoto/comments/:idComment", rt.removeComment)
	// ban a user
	rt.router.POST("/bans", rt.banUser)
	// unban a user
	rt.router.DELETE("/bans/:username", rt.unbanUser)

	return rt.router
}
