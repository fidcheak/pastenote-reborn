package main

import "database/sql"

type Note struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	FolderID int    `json:"folder_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
type NewNote struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type NoteResponse struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	FolderID  sql.NullInt64 `json:"folder_id"`
	CreatedAt string        `json:"created_at"`
}
type NoteUpdateRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
