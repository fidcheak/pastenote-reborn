package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// User

func SignUp(w http.ResponseWriter, r *http.Request) {
	var nu NewUser
	var u UserResponse
	if err := json.NewDecoder(r.Body).Decode(&nu); err != nil {
		http.Error(w, "Wrong format JSON (SignUp)", http.StatusBadRequest)
		return
	}
	log.Print("Req success")
	if err := AddUser(nu); err != nil {
		log.Print("AddUser error")
		log.Print(err)
		return
	}
	log.Print("AddUser success")

	u = GetUserByEmailAndPassword(nu.Email, nu.Password)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(u); err != nil {
		log.Printf("Encoding error JSON-response (SignUp): %v", err)
		return
	}
}
func SignIn(w http.ResponseWriter, r *http.Request) {
	var usi UserSignIn
	var u UserResponse
	if err := json.NewDecoder(r.Body).Decode(&usi); err != nil {
		http.Error(w, "Wrong format JSON (SignIn)", http.StatusBadRequest)
		return
	}
	u = GetUserByEmailAndPassword(usi.Email, usi.Password)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(u); err != nil {
		log.Printf("Encoding error JSON-response (SignIn): %v", err)
		return
	}
	log.Print("User response sent")
}
func UserDelete(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	if err := DeleteUser(userID); err != nil {
		log.Println("User Delete error")
		return
	}
	log.Print("User deleted")
}

// Note
func NoteCreate(w http.ResponseWriter, r *http.Request) {
	var nn NewNote
	if err := json.NewDecoder(r.Body).Decode(&nn); err != nil {
		http.Error(w, "Wrong format JSON (NoteCreate)", http.StatusBadRequest)
		return
	}
	if err := AddNote(nn); err != nil {
		log.Println("Note Create error")
		return
	}
	log.Println("Note created")
}
func NotesGet(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	notes := GetNotes(userID)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		log.Printf("Encoding error JSON-response (NotesGet): %v", err)
		return
	}
	log.Print("Notes response sent")
}
func NoteUpdate(w http.ResponseWriter, r *http.Request) {
	var nur NoteUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&nur); err != nil {
		http.Error(w, "Wrong format JSON (NoteDelete)", http.StatusBadRequest)
		return
	}
	if err := UpdateNote(nur); err != nil {
		log.Printf("NoteUpdate error: %v", err)
		return
	}
	log.Println("Note updated")
}
func NoteDelete(w http.ResponseWriter, r *http.Request) {
	noteIDStr := r.URL.Query().Get("note_id")
	if noteIDStr == "" {
		http.Error(w, "missing note_id", http.StatusBadRequest)
		return
	}
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		http.Error(w, "invalid note_id", http.StatusBadRequest)
		return
	}
	if err := DeleteNote(noteID); err != nil {
		log.Println("Note Delete error")
		return
	}
	log.Print("Note deleted")
}

func NotePin(w http.ResponseWriter, r *http.Request) {
	noteIDStr := r.URL.Query().Get("note_id")
	if noteIDStr == "" {
		http.Error(w, "missing note_id", http.StatusBadRequest)
		return
	}
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		http.Error(w, "invalid note_id", http.StatusBadRequest)
		return
	}
	if err := PinNote(noteID); err != nil {
		log.Println("Note Pin Error")
		return
	}
	log.Print("Note Pinned")
}

func NoteUnpin(w http.ResponseWriter, r *http.Request) {
	noteIDStr := r.URL.Query().Get("note_id")
	if noteIDStr == "" {
		http.Error(w, "missing note_id", http.StatusBadRequest)
		return
	}
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		http.Error(w, "invalid note_id", http.StatusBadRequest)
		return
	}
	if err := UnpinNote(noteID); err != nil {
		log.Println("Note Unpin Error")
		return
	}
	log.Print("Note Unpinned")
}
