package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var connector *sql.DB

func Connect() bool {
	var e error

	connector, e = sql.Open("postgres",
		`
host=10.14.206.27
port=5432
user=student
password=1234
dbname=blog
sslmode=disable
`)
	if e != nil {
		fmt.Println(e)
		return false
	}

	e = connector.Ping()
	if e != nil {
		fmt.Println(e)
		return false
	}

	fmt.Println("DATABASE IS CONNECTED")

	return true
}

func CreateBlog(caption, content string) {
	_, e := connector.Exec(`insert into "Blog" 
    ("Caption", "Content", "Published", "Image")
values ($1, $2, CURRENT_TIMESTAMP, '')`, caption, content)
	if e != nil {
		fmt.Println(e)
	}
}
