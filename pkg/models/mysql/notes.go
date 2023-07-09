package mysql

import (
	"awesomeProject/pkg/models"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

// NoteModel - Определяем тип который обертывает пул подключения sql.DB
type NoteModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *NoteModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO notes (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *NoteModel) Get(id int) (*models.Note, error) {

	stmt := `SELECT id, title, content, created, expires FROM notes
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)
	s := &models.Note{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (m *NoteModel) Latest() ([]*models.Note, error) {

	stmt := `SELECT id, title, content, created, expires FROM notes
    WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var notes []*models.Note

	for rows.Next() {
		s := &models.Note{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		notes = append(notes, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
