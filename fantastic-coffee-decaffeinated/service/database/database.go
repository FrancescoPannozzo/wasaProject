/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var DBcon AppDatabase

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	//Base example
	GetName() (string, error)
	//Base example
	SetName(name string) error

	// Get a username if present in the DB. If not present, the username provided will
	// be inserted in the DB. Returns the username and nil if successfull
	GetOrInsertUser(name string) (string, error)

	// Insert a user into the DB. Returns a feedback string and nil if succesfull
	InsertUser(name string) (string, error)

	// Get the userId by the provided username
	GetIdByName(name string) (string, error)

	// Get the username by the provided user id
	GetNameByID(userId string) (string, error)

	// Set a new username, return nil if succesfull
	ModifyUsername(oldName string, newName string) error

	// Insert photo data into the DB
	InsertPhoto(name string, idphoto string) (string, error)

	// Delete a photo, return a feedback string an nil if successfull
	DeletePhoto(idphoto string) (string, error)

	// Check if the provided username is in the DB
	UsernameInDB(name string) bool

	// Insert the user as follower into the DB. Return a
	InsertFollower(follower string, followed string) (string, error)

	// Delete a followed user. Return a feedback string an nil if successfull
	DeleteFollowed(follower string, followed string) (string, error)

	// Insert the Like data into the DB. Return a feedback string an nil if successfull
	LikePhoto(username string, idphoto string) (string, error)

	// Remove a like data drom the DB. Return a feedback string an nil if successfull
	RemoveLike(username string, idphoto string) (string, error)

	// Insert the comment to the photo(idphoto) in the DB.
	// Return a feedback message and nil if successfull.
	CommentPhoto(username string, idphoto string, comment string) (string, error)

	// Get the comments list of the provided photoId. Return nil if successfull
	GetComments(loggedUser string, photoID string) ([]utilities.Comment, error)

	// Delete a comment. Return a feedback string an nil if successfull
	RemoveComment(idcomment string) (string, error)

	// Check the username profile ownership
	CheckOwnership(userId string, username string) bool

	// Get the username from a photoId. Return the username and nil if successfull
	GetNameFromPhotoId(photoId string) (string, error)

	// Get the username from a commentId. Return the username and nil if successfull
	GetNameFromCommentId(commentId string) (string, error)

	// Ban the provided user. Return a feedback string an nil if successfull
	BanUser(banner string, banned string) (string, error)

	// Unban the provided user. Return a feedback string an nil if successfull
	UnbanUser(banner string, banned string) (string, error)

	// Check if the target user is banned from the logged user.
	CheckBan(loggedUser string, targetUser string) bool

	// Get user thumbnails objects. Returns the thumbnailsObject list and nil if successfull
	GetThumbnails(username string) ([]utilities.Thumbnail, error)

	// Get a post. Returns the postObject and nil if successfull
	GetPost(loggedUser string, photoId string) (utilities.Post, error)

	// Get followed thumbnails objects (the own stream). Returns the thumbnailsObject list and nil if successfull
	GetFollowedThumbnails(loggedUser string) ([]utilities.Thumbnail, error)

	// Get a username list by searching with the provided tergetUSer. Returns the usernames and nil if successfull
	GetUsernames(targetUser string) ([]utilities.Username, error)
	// Get the followed users by the provided loggedUser
	GetFollowed(loggedUser string) ([]utilities.Username, error)
	// Return the logged user followers
	GetFollowers(LoggedUser string) ([]utilities.Username, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure

	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='User';`).Scan(&tableName)

	if errors.Is(err, sql.ErrNoRows) {
		// Getting absolute path of db_schema.sql
		abs, err := filepath.Abs("./service/database/db_schema.sql")

		dat, errFile := ioutil.ReadFile(abs)
		if errFile != nil {
			return nil, fmt.Errorf("error reading database structure from file: %v", err)
		}

		sqlStmt := string(dat)
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
