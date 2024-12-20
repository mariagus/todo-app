package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID      int
	Title   string
	Urgency int
}

var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todos (
	 id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	 title TEXT,
	 urgency INTEGER DEFAULT 0
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllTodos() ([]Todo, error) {
	rows, err := DB.Query("SELECT id, title, urgency FROM todos ORDER BY urgency DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Urgency); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func createTodo(title, urgency string) error {
	_, err := DB.Exec("INSERT INTO todos (title, urgency) VALUES (?, ?)", title, urgency)
	return err
}

func deleteTodoByID(id string) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
