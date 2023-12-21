package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Database connection string
const dsn = "root:@tcp(localhost:3306)/game_champion"

// Holds the database connection
var db *sql.DB

// Establishes a database connection
func Connect() {
	conection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successful connection")
	db = conection
	Ping()
}

// Closes the database connection
func Close() {
	db.Close()
}

// Verifies the database connection
func Ping() {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// check if a table exists or not
func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// iterate through the table
	return rows.Next()
}

// NOTE: db.Query or db.Exec can now be used without db.
// Exec polymorphism
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// Query polymorphism
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}

// InsertChampion inserts a champion into the database and returns its ID
func InsertChampion(name, title, lore string) (int64, error) {
	// The database connection must be established before calling this function
	// You can handle the connection according to your needs

	// Example SQL statement for insertion into the Champion table
	query := "INSERT INTO Champion (Name, Title, Lore) VALUES (?, ?, ?)"

	// Execute the query and get the automatically generated ID
	result, err := Exec(query, name, title, lore)
	if err != nil {
		log.Fatalf("Error inserting champion into the database: %v", err)
		return 0, err
	}

	// Get the ID of the newly inserted champion
	championID, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error getting the ID of the inserted champion: %v", err)
		return 0, err
	}

	return championID, nil
}

// GetTagID gets the ID of an existing tag or returns 0 if it doesn't exist.
func GetTagID(tag string) (int, error) {
	var tagID int
	err := db.QueryRow("SELECT Id FROM Tags WHERE Name = ?", tag).Scan(&tagID)
	if err == sql.ErrNoRows {
		return 0, nil // Tag not found
	} else if err != nil {
		return 0, err // Other error
	}
	return tagID, nil // ID of the found tag
}

// GetTagID gets the ID of an existing tag or returns 0 if it doesn't exist.
func GetChampionID(champion string) (int, error) {
	var championID int
	err := db.QueryRow("SELECT Id FROM Champion WHERE Name = ?", champion).Scan(&championID)
	if err == sql.ErrNoRows {
		return 0, nil // Tag not found
	} else if err != nil {
		return 0, err // Tag not found
	}
	return championID, nil
}

// InsertTag inserts a new tag and returns its ID.
func InsertTag(championID int, tag string) error {
	_, err := Exec("INSERT INTO Tags (Id_Champion, Name) VALUES (?, ?)", championID, tag)
	return err
}

// GetSkinID gets the ID of an existing tag or returns 0 if it doesn't exist.
func GetSkinID(Id_Num string) (int64, error) {

	// Example SQL statement for insertion into the Champion table
	query := "SELECT Id FROM Skins WHERE Id_Num = ?"

	// Execute the query and get the automatically generated ID
	result, err := Exec(query, Id_Num)
	if err != nil {
		log.Fatalf("Error inserting champion into the database: %v", err)
		return 0, err
	}

	// Get the ID of the newly inserted champion
	championID, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error getting the ID of the inserted champion: %v", err)
		return 0, err
	}

	return championID, nil

}

// InsertTag inserts a new tag and returns its ID.
func InsertSkins(Id_Num string, num, championID int, name string) error {
	_, err := Exec("INSERT INTO Skins (Id_Num, Num, Id_Champion, Name) VALUES (?,?, ?,?)", Id_Num, num, championID, name)
	return err
}
