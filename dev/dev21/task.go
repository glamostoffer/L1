package main

import (
	"database/sql"
	"fmt"
)

// Интерфейс для работы с базой данных
type DatabaseAdapter interface {
	Connect() error
	Disconnect() error
	GetItem(itemId int) (string, error)
}

// Адаптер для работы с MySQL базой данных
type MySQLAdapter struct {
	db *sql.DB
}

func (a *MySQLAdapter) Connect() error {
	// Подключение к MySQL базе данных
	return nil
}

func (a *MySQLAdapter) Disconnect() error {
	// Отключение от MySQL базы данных
	return nil
}

func (a *MySQLAdapter) GetItem(itemId int) (string, error) {
	// Выполнение запроса к MySQL базе данных
	return fmt.Sprintf("get item with id: %d form MySQL", itemId), nil
}

// Адаптер для работы с PostgreSQL базой данных
type PostgreSQLAdapter struct {
	// Добавьте необходимые поля для работы с PostgreSQL
}

func (a *PostgreSQLAdapter) Connect() error {
	// Подключение к PostgreSQL базе данных
	return nil
}

func (a *PostgreSQLAdapter) Disconnect() error {
	// Отключение от PostgreSQL базе данных
	return nil
}

func (a *PostgreSQLAdapter) GetItem(itemId int) (string, error) {
	// Выполнение запроса к PostgreSQL базе данных
	return fmt.Sprintf("get item with id: %d form PostgreSQL", itemId), nil
}

func main() {
	// Используем MySQL базу данных
	mysqlAdapter := &MySQLAdapter{}
	useDatabase(mysqlAdapter)

	// Переключаемся на использование PostgreSQL базы данных
	postgreSQLAdapter := &PostgreSQLAdapter{}
	useDatabase(postgreSQLAdapter)
}

func useDatabase(adapter DatabaseAdapter) {
	err := adapter.Connect()
	if err != nil {
		fmt.Println("error connecting to the database:", err)
		return
	}
	defer adapter.Disconnect()

	result, err := adapter.GetItem(4)
	if err != nil {
		fmt.Println("error executing query:", err)
		return
	}

	fmt.Println("get item result:", result)
}
