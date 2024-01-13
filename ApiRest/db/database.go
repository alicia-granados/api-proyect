package db

import (
	"ApiRest/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type RealDBRepo struct {
	DB *sql.DB
}

// Establishes a database connection
func (r *RealDBRepo) Connect() {
	dsn, err := config.DSNDatabase()
	if err != nil {
		log.Fatal("Error configuring server:", err)
	}
	conection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successful connection")
	r.DB = conection
	r.DB.Ping()
}

// Closes the database connection
func (r *RealDBRepo) Close() {
	r.DB.Close()
}

// Verifies the database connection
func (r *RealDBRepo) Ping() {
	if err := r.DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

// check if a table exists or not
func (r *RealDBRepo) ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)
	rows, err := r.DB.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// iterate through the table
	return rows.Next()
}

// ExistsID checks if a specific ID exists in the table
func (r *RealDBRepo) ExistsID(table string, id int) bool {

	// Execute the SQL query
	var count int

	// Execute the SQL query and get the count
	err := r.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*)  FROM %s WHERE Id = ?", table), id).Scan(&count)
	// Execute the SQL query
	if err != nil {
		return false
	}

	return count != 0
}

// GetTagID gets the ID of an existing tag or returns 0 if it doesn't exist.
func (r *RealDBRepo) GetTagID(tag string) (int, error) {
	var tagID int
	err := r.DB.QueryRow("SELECT Id FROM Tags WHERE Name = ?", tag).Scan(&tagID)
	if err == sql.ErrNoRows {
		return 0, nil // Tag not found
	} else if err != nil {
		return 0, err // Other error
	}
	return tagID, nil // ID of the found tag
}

// InsertTag inserts a new tag and returns its ID.
func (r *RealDBRepo) InsertTag(championID int, tag string) error {
	_, err := r.DB.Exec("INSERT INTO Tags (Id_Champion, Name) VALUES (?, ?)", championID, tag)
	return err
}
