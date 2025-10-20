package database

func (db *appdbimpl) SetUserPhoto(userID int, photoBytes []byte) error {

	_, err := db.c.Exec(`
        UPDATE Users 
        SET photo = ? 
        WHERE ID = ?`, photoBytes, userID)

	if err != nil {
		return err
	}

	return nil
}
