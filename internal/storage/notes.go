package storage

import (
	"context"
	"web2022/internal/models"
)

func (s *Storage) GetNotesByUserID(userID int) ([]models.Note, error) {
	rows, err := s.pool.Query(context.Background(), `
	SELECT id, name, description, status
	FROM notes
	WHERE user_id = $1; 
	`, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data []models.Note
	for rows.Next() {
		var item models.Note
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Status,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}
	return data, nil
}

func (s *Storage) NewNote(item models.Note) (id int, err error) {
	err = s.pool.QueryRow(context.Background(), `
		INSERT INTO notes(name, description, status, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING Id;
	`,
		item.Name,
		item.Description,
		item.Status,
		item.UserID,
	).Scan(&id)

	return id, err
}

func (s *Storage) UpdateNote(item models.Note) (err error) {
	_, err = s.pool.Query(context.Background(), `
	UPDATE notes SET
		name = $2,
		description = $3,
		status = $4
	WHERE id = $1; 
	`,
		item.ID,
		item.Name,
		item.Description,
		item.Status,
	)

	return err
}

func (s *Storage) DeleteNoteByID(id int) (err error) {
	_, err = s.pool.Query(context.Background(), "DELETE FROM notes WHERE id = $1;", id)

	return err
}
