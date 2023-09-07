package storage

import (
	"context"
	"web2022/internal/models"
)

// Методы БД по работе с пользователем.
func (s *Storage) GetUserByID(id int) (models.User, error) {
	var user models.User

	err := s.pool.QueryRow(context.Background(), `
		SELECT id, name, login, password, session_id
		FROM users WHERE id = $1`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.Password,
		&user.SessionID,
	)

	return user, err
}

func (s *Storage) GetUserPasswordByName(name string) (string, error) {
	var pwd string
	err := s.pool.QueryRow(context.Background(), `SELECT password FROM users WHERE name = $1`, name).Scan(&pwd)

	return pwd, err
}

func (s *Storage) NewUser(user models.User) (id int, err error) {
	err = s.pool.QueryRow(context.Background(), `
		INSERT INTO users(name, login, password, session_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;`,
		user.Name,
		user.Login,
		user.Password,
		user.SessionID,
	).Scan(&id)

	return id, err
}

func (s *Storage) DeleteUserByID(id int) (err error) {
	_, err = s.pool.Exec(context.Background(), `
		DELETE FROM users WHERE id = $1;
		DELETE FROM notes WHERE user_id = $1;`,
		id,
	)
	return err
}
