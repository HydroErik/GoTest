package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	connStr := os.Getenv("POSGRSTRING")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	query := "SELECT * FROM sqlth_te LIMIT 3;"
	rows, err := db.Query(query)

	if err != nil{
		log.Fatal(err)
	}

	type db_row struct {
		id int
		path string
		scid int
		datatype int
		querymode int
		created int
		retired int
	}

	
	//var res string
	//defer rows.Close()
	colmnNames, err := rows.Columns()
	fmt.Println(colmnNames)

	for rows.Next() {
		var cur_row db_row
		rows.Scan(&cur_row.id, &cur_row.path, &cur_row.scid, &cur_row.datatype, &cur_row.querymode, &cur_row.created, &cur_row.retired)
		fmt.Printf("Tag Id: %d tag path: %s \n",cur_row.id,cur_row.path)
		
	}

}
	
	
