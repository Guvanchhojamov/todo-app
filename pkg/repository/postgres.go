package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Table name constants
const (
	usersTable      = "users"
	listsTable      = "todo_lists"
	itemsTable      = "todo_items"
	usersListsTable = "users_lists"
	listsItemsTable = "lists_items"
)

type Configs struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(config Configs) (*sqlx.DB, error) {
	dbString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)
	db, err := sqlx.Open("postgres", dbString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal("Database connection error: \n", err.Error())
		return nil, err
	}
	return db, nil
}
