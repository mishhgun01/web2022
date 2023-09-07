package models

type User struct {
	ID        int
	Name      string
	Login     string
	Password  string
	SessionID string
}

type Note struct {
	ID          int
	UserID      int
	Name        string
	Description string
	Status      uint8
}
