package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-adodb"
	//_ "github.com/lib/pq"
)

type Fruit struct {
	Name      string         `db:"Name"`
	Price     int32          `db:"Price"`
	Color     string         `db:"Color"`
	Code      string         `db:"Code"`
	StoreCode sql.NullString `db:"StoreCode"`
}

func main() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("adodb", "Provider=SQLOLEDB;Data Source=qds114812583.my3w.com,1433;Initial Catalog=qds114812583_db;User ID=qds114812583;password=Eland123;")
	if err != nil {
		log.Fatalln(err)
	}

	// Query the database, storing results in a []Person (wrapped in []interface{})
	fruits := []Fruit{}
	err2 := db.Select(&fruits, "select Name,Price,Color,Code,StoreCode from dbo.Fruit")
	fmt.Printf("%v", 11111)
	log.Println("err2:", err2)
	fmt.Printf("%v", 2222)
	fmt.Printf("%+v", fruits)
	for i, v := range fruits {
		fmt.Println(i, v)
	}
	fmt.Println(fruits[3])
	fmt.Println(fruits[3].StoreCode.String)

}
