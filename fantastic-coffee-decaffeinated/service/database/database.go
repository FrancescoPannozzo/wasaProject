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

	//CheckUser(name string) (string, error)
	GetOrInsertUser(name string) (string, error, int)
	InsertUser(name string) (string, error, int)
	GetIdByName(name string) (string, error)
	GetNameByID(userId string) (string, error)
	ModifyUsername(oldName string, newName string) error
	InsertPhoto(name string, idphoto string) (string, error, int)
	DeletePhoto(idphoto string) (string, error, int)
	UsernameInDB(name string) bool
	InsertFollower(follower string, followed string) (string, error, int)
	//GetPhoto(name string, idphoto)
	DeleteFollowed(follower string, followed string) (string, error, int)
	// Give a Like
	LikePhoto(username string, idphoto string) (string, error, int)
	// Remove a like
	RemoveLike(username string, idphoto string) (string, error, int)
	// Comment a photo, usernaname is the writing comment user
	//comment is the text of the comment
	CommentPhoto(username string, idphoto string, comment string) (string, error, int)
	// Delete a comment
	RemoveComment(username string, idphoto string, idcomment string) (string, error, int)
	// Check the username profile ownership
	CheckOwnership(userId string, username string) bool
	// get the username from a photoId
	GetNameFromPhotoId(photoId string) (string, error)
	// ban user
	BanUser(banner string, banned string) (string, error, int)
	// unban user
	UnbanUser(banner string, banned string) (string, error, int)
	//Check bans
	CheckBan(loggedUser string, targetUser string) bool
	// get user thumbnails objects
	GetThumbnails(username string) ([]utilities.Thumbnail, error, int)
	// get followed thumbnails objects
	GetFollowedThumbnails(loggedUser string) ([]utilities.Thumbnail, error, int)

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
