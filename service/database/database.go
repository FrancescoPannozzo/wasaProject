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
	// Get a username if present in the DB. If not present, the username provided will
	// be inserted in the DB. Returns the username and nil if successful
	GetOrInsertUser(name string) (string, error)

	// Insert a user into the DB. Returns a feedback string and nil if succesful
	InsertUser(name string) (string, error)

	// Get the userId by the provided username
	GetIdByName(name string) (string, error)

	// Get the username by the provided user id
	GetNameByID(userId string) (string, error)

	// Set a new username, return nil if succesfull
	ModifyUsername(oldName string, newName string) error

	// Insert photo data into the DB
	InsertPhoto(name string, idphoto string, extension string) (string, error)

	// Delete a photo, return a feedback string an nil if successful
	DeletePhoto(idphoto string) (string, error)

	// Check if the provided username is in the DB
	UsernameInDB(name string) bool

	// Insert the user as follower into the DB. Return a
	InsertFollower(follower string, followed string) (string, error)

	// Delete a followed user. Return a feedback string an nil if successful
	DeleteFollowed(follower string, followed string) (string, error)

	// Insert the Like data into the DB. Return a feedback string an nil if successful
	LikePhoto(username string, idphoto string) (string, error)

	// Remove a like data drom the DB. Return a feedback string an nil if successful
	RemoveLike(username string, idphoto string) (string, error)

	// Insert the comment to the photo(idphoto) in the DB.
	// Return a feedback message and nil if successful.
	CommentPhoto(username string, idphoto string, comment string) (string, error)

	// Get the comments list of the provided photoId. Return nil if successful
	GetComments(loggedUser string, photoID string) ([]utilities.Comment, error)

	// Delete a comment. Return a feedback string an nil if successful
	RemoveComment(idcomment string) (string, error)

	// Check the username profile ownership
	CheckOwnership(userId string, username string) bool

	// Get the username from a photoId. Return the username and nil if successful
	GetNameFromPhotoId(photoId string) (string, error)

	// Get the username from a commentId. Return the username and nil if successful
	GetNameFromCommentId(commentId string) (string, error)

	// Ban the provided user. Return a feedback string an nil if successful
	BanUser(banner string, banned string) (string, error)

	// Unban the provided user. Return a feedback string an nil if successful
	UnbanUser(banner string, banned string) (string, error)

	// Check if the target user is banned from the logged user.
	CheckBan(loggedUser string, targetUser string) bool

	// Get user thumbnails objects. Returns the thumbnailsObject list and nil if successful
	GetThumbnails(username string) ([]utilities.Thumbnail, error)

	// Get a post. Returns the postObject and nil if successful
	GetPost(loggedUser string, photoId string) (utilities.Post, error)

	// Get followed thumbnails objects (the own stream). Returns the thumbnailsObject list and nil if successful
	GetFollowedThumbnails(loggedUser string) ([]utilities.Thumbnail, error)

	// Get a username list by searching with the provided tergetUSer. Returns the usernames and nil if successful
	GetUsernames(targetUser string) ([]string, error)
	// Get the followed users by the provided loggedUser
	GetFollowed(loggedUser string) ([]string, error)
	// Return the logged user followers
	GetFollowers(LoggedUser string) ([]string, error)

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
		abs, errAbs := filepath.Abs("./service/database/db_schema.sql")
		if errAbs != nil {
			return nil, fmt.Errorf("error creating filepath: %w", err)
		}

		dat, errFile := ioutil.ReadFile(abs)
		if errFile != nil {
			return nil, fmt.Errorf("error reading database structure from file: %w", err)
		}

		sqlStmt := string(dat)
		_, errEx := db.Exec(sqlStmt)
		if errEx != nil {
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
