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
	GetName() (string, error)
	SetName(name string) error

	//Get an username if present in the DB, if not the username provided will
	//be inserted in the DB. Return the
	GetOrInsertUser(name string) (string, error)

	InsertUser(name string) (string, error)

	GetIdByName(name string) (string, error)

	//get the username by the id provided
	GetNameByID(userId string) (string, error)

	ModifyUsername(oldName string, newName string) error

	InsertPhoto(name string, idphoto string) (string, error)

	DeletePhoto(idphoto string) (string, error, int)

	UsernameInDB(name string) bool

	InsertFollower(follower string, followed string) (string, error)

	// Delete a followed user.
	DeleteFollowed(follower string, followed string) (string, error)

	// Give a Like
	LikePhoto(username string, idphoto string) (string, error)

	// Remove a like
	RemoveLike(username string, idphoto string) (string, error)

	// Insert the comment on the photoID provided in the DB.
	// Return a feedback message and nil if successfull.
	// Return a feedback message and an error excetution query otherwise.
	CommentPhoto(username string, idphoto string, comment string) (string, error)

	// Get comments
	GetComments(loggedUser string, photoID string) ([]utilities.Comment, error)

	// Delete a comment
	RemoveComment(username string, idphoto string, idcomment string) (string, error)

	// Check the username profile ownership
	CheckOwnership(userId string, username string) bool

	// get the username from a photoId
	GetNameFromPhotoId(photoId string) (string, error)

	// Ban the provided user
	BanUser(banner string, banned string) (string, error)

	// unban user
	UnbanUser(banner string, banned string) (string, error)

	// Check bans
	CheckBan(loggedUser string, targetUser string) bool

	// Get user thumbnails objects
	GetThumbnails(username string) ([]utilities.Thumbnail, error)

	// get followed thumbnails objects
	GetFollowedThumbnails(loggedUser string) ([]utilities.Thumbnail, error)

	GetUsernames(targetUser string) ([]utilities.Username, error)

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
		// Getting absolute path of db_schema.txt
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
