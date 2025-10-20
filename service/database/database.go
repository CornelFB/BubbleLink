package database

import (
	"database/sql"
	"errors"
	"fmt"
	//"image"
	//"image/jpeg"
	// "time"
)

type AppDatabase interface {
	Ping() error

	CheckIfUserExists(username string) (bool, error)

	AddNewUser(username string, country string, city string, securityKey string) (int, error)

	GetUserName(userID int) (string, error)
	SetUserName(userID int, username string) error

	GetUserPhoto(userID int) ([]byte, error)
	SetUserPhoto(userID int, photoBytes []byte) error

	GetUserKey(userID int) (string, error)
	GetUserID(username string) (int, error)
	//	GetUserIDbyKey(security_key string) (int, error)

}

type appdbimpl struct {
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building an AppDatabase")
	}

	usersTableStmt := `CREATE TABLE IF NOT EXISTS Users (
				id INTEGER NOT NULL PRIMARY KEY,
				username TEXT NOT NULL,
				country TEXT NOT NULL,
				city TEXT NOT NULL,
				security_key TEXT NOT NULL,
				jpeg_photo BLOB
				);`
	if _, err := db.Exec(usersTableStmt); err != nil {
		return nil, fmt.Errorf("error creating Users table: %w", err)
	}
	bubblesTableStmt := `CREATE TABLE IF NOT EXISTS Bubbles (
		id INTEGER NOT NULL PRIMARY KEY,
		qr_code TEXT NOT NULL UNIQUE,
		description TEXT,
		jpeg_photo BLOB
	);`
	if _, err := db.Exec(bubblesTableStmt); err != nil {
		return nil, fmt.Errorf("error creating Bubbles table: %w", err)
	}

	// --- Places table ---
	// A place has coordinates and links to one bubble.
	placesTableStmt := `CREATE TABLE IF NOT EXISTS Places (
		id INTEGER NOT NULL PRIMARY KEY,
		bubble_id INTEGER,
		x INTEGER NOT NULL,
		y INTEGER NOT NULL,
		FOREIGN KEY (bubble_id) REFERENCES Bubbles(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(placesTableStmt); err != nil {
		return nil, fmt.Errorf("error creating Places table: %w", err)
	}

	// --- Posts table ---
	// Each post belongs to a bubble and is written by a user.
	postsTableStmt := `CREATE TABLE IF NOT EXISTS Posts (
		id INTEGER NOT NULL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		bubble_id INTEGER NOT NULL,
		text TEXT,
		jpeg_photo BLOB,
		FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
		FOREIGN KEY (bubble_id) REFERENCES Bubbles(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(postsTableStmt); err != nil {
		return nil, fmt.Errorf("error creating Posts table: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
