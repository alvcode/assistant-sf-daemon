package repository

import (
	"database/sql"
)

type ConfigRepository interface {
	CreateTableIfNotExists() error
	Get(name string) (string, error)
	Delete(name string) error
	Upsert(name string, value string) error
}

type configRepository struct {
	db *sql.DB
}

func newConfigRepository(db *sql.DB) ConfigRepository {
	return &configRepository{db: db}
}

func (r *configRepository) CreateTableIfNotExists() error {
	query := `CREATE TABLE IF NOT EXISTS config (
			name TEXT NOT NULL,
			value TEXT NOT NULL
		)`

	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *configRepository) Get(name string) (string, error) {
	var value string
	err := r.db.QueryRow("SELECT value FROM config WHERE name = ?", name).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, err
}

func (r *configRepository) Delete(name string) error {
	_, err := r.db.Exec("DELETE FROM config WHERE name = ?", name)
	if err != nil {
		return err
	}
	return nil
}

func (r *configRepository) Upsert(name string, value string) error {
	err := r.Delete(name)
	if err != nil {
		return err
	}

	query := `INSERT INTO config (name, value) VALUES (?, ?)`
	_, err = r.db.Exec(query, name, value)
	if err != nil {
		return err
	}

	return nil
}
