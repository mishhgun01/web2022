package models

// User - модель данных пользователя.
type User struct {
	ID        int
	Name      string
	Login     string
	Password  string
	SessionID string
}

// Note - модель данных заметки.
type Note struct {
	ID          int
	UserID      int
	Name        string
	Description string
	Status      uint8
}
