package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "sadat"
	DB_PASSWORD = "11235813"
	DB_NAME     = "school_db"
)

func main() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=localhost port=5433 sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the school_db database!")

	// creating tbale
	// createTable(db)

	// inserting a student
	insertStudent(db, "Aa", 24, "10")

	// update students name
	// updateStudent(db, 2, "Preity")

	// Query the students
	queryStudents(db)

	// deleet a student
	//deleteStudent(db, 5)

	/*
		// Create a table
		createTableQuery := `
		CREATE TABLE IF NOT EXISTS students (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			age INT NOT NULL,
			class TEXT NOT NULL
		);`

		_, err = db.Exec(createTableQuery)
		if err != nil {
			log.Fatal("Failed to create table:", err)
		}
		fmt.Println("Table 'students' created or already exists.")

		// now inserting vlues in the table

		// var id int
		// var name string
		// err = db.QueryRow("INSERT INTO students(name,age,class) VALUES($1,$2,$3) returning id,name;", "Sadat", 24, "10").Scan(&id, &name)
		// if err != nil {
		// 	log.Fatal("Failed to insert row:", err)
		// }
		// fmt.Println("last inserted id =", id, "name", name)

		// now updating

		fmt.Println("# Updating")
		stmt, err := db.Prepare("update students set name=$1 where id=$2")
		if err != nil {
			log.Fatal("Failed to update: Prepare error", err)
		}

		id := 2
		res, err := stmt.Exec("Preeety", id)
		if err != nil {
			log.Fatal("Failed to update: excution error", err)
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal("Affected Row Error", err)
		}

		fmt.Println(affect, "rows changed")

		// now querying
		fmt.Println("# Querying")
		rows, err := db.Query("SELECT * FROM students")
		if err != nil {
			log.Fatal("Qeury error", err)
		}
		// fmt.Println("printng select result", rows)

		fmt.Println("id  |   name   |   age  | class")
		for rows.Next() {
			var id int
			var name string
			var age int
			var class string
			err = rows.Scan(&id, &name, &age, &class)

			if err != nil {
				log.Fatal("Scan error", err)
			}

			fmt.Printf("%3v | %8v | %6v | %2v\n", id, name, age, class)
		}

		// now deleting from databse

		fmt.Println("# Deleting")
		stmt, err = db.Prepare("delete from students where id=$1")
		if err != nil {
			log.Fatal("Preapre statement error", err)
		}

		id = 5
		res, err = stmt.Exec(id)
		if err != nil {
			log.Fatal("Delete execution error", err)
		}
		affect, err = res.RowsAffected()
		if err != nil {
			log.Fatal("Rows delete affected error", err)
		}
		fmt.Println(affect, "rows changed")

	*/

}

// Function to create the students table
func createTable(db *sql.DB) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL,
		class TEXT NOT NULL
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	fmt.Println("Table 'students' created or already exists.")
}

// Function to insert a student
func insertStudent(db *sql.DB, name string, age int, class string) {
	var id int
	err := db.QueryRow("INSERT INTO students(name, age, class) VALUES($1, $2, $3) RETURNING id;", name, age, class).Scan(&id)
	if err != nil {
		log.Fatal("Failed to insert row:", err)
	}
	fmt.Println("Last inserted id =", id)
}

// Function to update a student's name
func updateStudent(db *sql.DB, id int, newName string) {
	fmt.Println("# Updating")
	stmt, err := db.Prepare("UPDATE students SET name=$1 WHERE id=$2")
	if err != nil {
		log.Fatal("Failed to update: Prepare error", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(newName, id)
	if err != nil {
		log.Fatal("Failed to update: Execution error", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Affected Rows Error", err)
	}

	fmt.Println(affected, "rows changed")
}

// Function to query and print students
func queryStudents(db *sql.DB) {
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		log.Fatal("Query error", err)
	}

	defer rows.Close()

	fmt.Println("id  |   name   |   age  | class")
	for rows.Next() {
		var id int
		var name string
		var age int
		var class string
		err = rows.Scan(&id, &name, &age, &class)
		if err != nil {
			log.Fatal("Scan error", err)
		}

		fmt.Printf("%3v | %8v | %6v | %2v\n", id, name, age, class)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Row iteration error", err)
	}
}

// Function to delete a student by id
func deleteStudent(db *sql.DB, id int) {
	fmt.Println("# Deleting")
	stmt, err := db.Prepare("DELETE FROM students WHERE id=$1")
	if err != nil {
		log.Fatal("Prepare statement error", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal("Delete execution error", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Rows delete affected error", err)
	}
	fmt.Println(affected, "rows changed")
}
