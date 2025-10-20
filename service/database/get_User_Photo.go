package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetUserPhoto(userID int) ([]byte, error) {

	var photoBytes []byte
	err := db.c.QueryRow(`
		SELECT jpeg_photo FROM Users WHERE ID = ?`, userID).Scan(&photoBytes)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return nil, fmt.Errorf("user %d not found or photo missing: %w", userID, err)
		}

		return nil, fmt.Errorf("error querying user photo: %w", err)
	}

	if len(photoBytes) == 0 {
		return nil, sql.ErrNoRows
	}

	return photoBytes, nil
}
