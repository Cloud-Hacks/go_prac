package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("./articleInfo.db")

	// open a db performing sql operation
	db, err := sql.Open("sqlite3", "./articleInfo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sql query to create a table and execute the table
	sqlStmt := `
	create table articleList (id integer not null primary key, title text, author text, ratings integer, review text);
	delete from articleList;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// sql query to create a prepared stmnt in a table
	stmt, err := tx.Prepare("insert into articleList(id, title, author, ratings, review) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 5; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("kloudknox%03d", i), fmt.Sprintf("%s", "ETY"), fmt.Sprintf("%d", i), fmt.Sprintf("Good"))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	// sql query to select or fetch the data from db
	rows, err := db.Query("select id, title, author, ratings, review from articleList")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id      int
			title   string
			author  string
			ratings int32
			review  string
		)
		err = rows.Scan(&id, &title, &author, &ratings, &review)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, title, author, ratings)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// sql query to fetch the data by id
	stmt, err = db.Prepare("select title, author, ratings from articleList where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var (
		title   string
		author  string
		ratings int32
	)
	err = stmt.QueryRow("3").Scan(&title, &author, &ratings)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(title)

	// sql query to delete an item
	_, err = db.Exec("delete from articleList")
	if err != nil {
		log.Fatal(err)
	}

	// sql query to update the db
	_, err = db.Exec("insert into articleList(id, title, author, review, ratings) values(1, 'Lets GO', 'Alex', 4, 'Good'), (2, '3 Idiots', 'Chetan', 3, 'Better'), (3, 'Three mistakes in my life', 'Ety', 4, 'Excellet')")
	if err != nil {
		log.Fatal(err)
	}

	// sql query to select or fetch the data from db
	rows, err = db.Query("select id, title, ratings, review from articleList")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var author string
		var ratings int32
		err = rows.Scan(&id, &title, &author, &ratings)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
