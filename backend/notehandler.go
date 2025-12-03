package main

import "log"

func AddNote(n NewNote) error {
	query := `
		INSERT INTO notes (title, content, user_id)
		VALUES ($1, $2, $3)
	`
	if _, err := DB.Exec(
		query,
		n.Title,
		n.Content,
		n.UserID,
	); err != nil {
		return err
	}
	log.Print("Note Added Succesfully")
	return nil
}
func GetNotes(uid int) []NoteResponse {
	var notes []NoteResponse
	query := `
		SELECT id, title, content, folder_id, pinned, featured, created_at
		FROM notes
		WHERE user_id = $1
	`
	rows, err := DB.Query(
		query,
		uid,
	)
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var nr NoteResponse
		if err := rows.Scan(
			&nr.ID,
			&nr.Title,
			&nr.Content,
			&nr.FolderID,
			&nr.Pinned,
			&nr.Featured,
			&nr.CreatedAt,
		); err != nil {
			log.Printf("GetNotes Error %v", err)
			return nil
		}
		notes = append(notes, nr)
	}
	return notes
}
func UpdateNote(nu NoteUpdateRequest) error {
	query := `
		UPDATE notes
		SET title = COALESCE(NULLIF($1, ''), title),
		    content = COALESCE(NULLIF($2, ''), content)
		WHERE id = $3
	`
	if _, err := DB.Exec(
		query,
		nu.Title,
		nu.Content,
		nu.ID,
	); err != nil {
		return err
	}
	return nil
}
func DeleteNote(nid int) error {
	query := `
		DELETE FROM notes
		WHERE id = $1
	`
	_, err := DB.Exec(query, nid)
	return err
}
func PinNote(nid int) error {
	query := `
		UPDATE notes
		SET pinned = TRUE
		WHERE id = $1
	`
	_, err := DB.Exec(query, nid)
	return err
}
func UnpinNote(nid int) error {
	query := `
		UPDATE notes
		SET pinned = FALSE
		WHERE id = $1
	`
	_, err := DB.Exec(query, nid)
	return err
}
func FeatureNote(nid int) error {
	query := `
		UPDATE notes
		SET featured = TRUE
		WHERE id = $1
	`
	_, err := DB.Exec(query, nid)
	return err
}
func UnfeatureNote(nid int) error {
	query := `
		UPDATE notes
		SET featured = FALSE
		WHERE id = $1
	`
	_, err := DB.Exec(query, nid)
	return err
}
