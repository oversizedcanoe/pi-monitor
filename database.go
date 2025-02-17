package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "./database.db")

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func closeDatabase(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func executeCommand(text string, args ...any) int {
	db := openDatabase()

	result, err := db.Exec(text, args...)

	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	closeDatabase(db)

	return int(count)
}

func queryRow(queryText string, args ...any) *sql.Row {
	db := openDatabase()

	row := db.QueryRow(queryText, args...)

	closeDatabase(db)

	return row
}

func queryRows(queryText string, args ...any) *sql.Rows {
	db := openDatabase()

	rows, err := db.Query(queryText, args...)

	if err != nil {
		panic(err)
	}

	// Ensure to call defer rows.Close() in calling method... Probably a better way to handle this

	closeDatabase(db)

	return rows
}
