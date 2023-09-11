package storage

import (
	"fmt"
	"testing"
	"web2022/internal/config"
	"web2022/internal/models"
)

var cfg config.Config

func TestStorage_NoteCRUD(t *testing.T) {
	cfg.Fill()
	s, err := New(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPwd,
		cfg.DbName,
	))
	if err != nil {
		t.Fatal("Не удалось подключиться к БД")
	}

	u := models.User{
		Name:     "test_user",
		Login:    "test_login",
		Password: "test",
	}

	userID, err := s.NewUser(u)
	if err != nil || userID == 0 {
		t.Fatal("Не удалось создать пользователя")
	}

	var n = models.Note{
		Name:        "Test",
		Description: "Description_test",
		UserID:      userID,
	}

	id, err := s.NewNote(n)
	if err != nil || id == 0 {
		t.Fatal("Не удалось создать запись")
	}

	notes, err := s.GetNotesByUserID(userID)
	if err != nil || len(notes) == 0 {
		t.Fatal("Ошибка в получении записи")
	}

	for _, note := range notes {
		if note.ID == id {
			n = note
		}
	}

	n.ID = id
	n.Name = "TEST_2"
	err = s.UpdateNote(n)
	if err != nil {
		t.Fatal("Ошибка в обновлении записи.")
	}

	err = s.DeleteNoteByID(n.ID)
	if err != nil {
		t.Fatal("Ошибка в удалении записи.")
	}
}

func TestStorage_UserCRUD(t *testing.T) {
	cfg.Fill()
	s, err := New(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPwd,
		cfg.DbName,
	))
	if err != nil {
		t.Fatal("Не удалось подключиться к БД")
	}

	u := models.User{
		Name:     "test_user",
		Login:    "test_login",
		Password: "test",
	}

	id, err := s.NewUser(u)
	if err != nil || id == 0 {
		t.Fatal("Не удалось создать пользователя")
	}

	u, err = s.GetUserByID(id)
	if err != nil || id != u.ID {
		t.Fatal("Ошибка при получении пользователя")
	}

	err = s.DeleteUserByID(u.ID)
	if err != nil {
		t.Fatal("Ошибка при удалении пользователя")
	}
}
