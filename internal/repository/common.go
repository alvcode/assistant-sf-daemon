package repository

import (
	"assistant-sf-daemon/internal/dict"
	"assistant-sf-daemon/internal/service"
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var connection *sql.DB

type Creator interface {
	Config() ConfigRepository
}

type creator struct{}

func GetCreator() Creator {
	return &creator{}
}

func getConnection() *sql.DB {
	if connection != nil {
		return connection
	}
	appPath, err := service.GetAppPath()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("sqlite3", filepath.Join(appPath, dict.DBName)+"?_journal_mode=WAL")
	if err != nil {
		panic(err)
	}
	connection = db
	return connection
}

func (c *creator) Config() ConfigRepository {
	return newConfigRepository(getConnection())
}
