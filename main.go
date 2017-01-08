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

type FruitParam struct {
	Name      string         `db:"Name"`
	Price     int32          `db:"Price"`
	Color     string         `db:"Color"`
	Code      string         `db:"Code"`
	StoreCode sql.NullString `db:"StoreCode"`
}

var msdb *sqlx.DB

func main() {
	var err error
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	msdb, err = sqlx.Connect("adodb", "Provider=SQLOLEDB;Data Source=qds114812583.my3w.com,1433;Initial Catalog=qds114812583_db;User ID=qds114812583;password=Eland123;")
	if err != nil {
		log.Fatalln(msdb, err)
	}

	//TestSearch()//1.search
	//TestSp()//2.sp
	//TestWithParam()    //3.query with param
	//TestSpWithParam2() //4.sp with param
	TestSpWithParam3()

}

func TestSp() {
	rows, _ := msdb.Query("[up_Fruit_R1]")
	f := new(Fruit)
	for rows.Next() {
		if e := rows.Scan(&f.Name, &f.Price); e != nil {
			fmt.Println("==============", e)
		} else {
			fmt.Println("f1:", f)
		}
	}

	fmt.Printf("%+v", f)
}
func TestWithParam() {
	fruitParam := new(FruitParam)
	fruitParam.Code = "A2"
	f := Fruit{}
	err := msdb.Get(&f, "select top 1 Name,Price,Color,Code,StoreCode from dbo.Fruit WHERE Code =?", fruitParam.Code)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Printf("%+v", f)
}
func TestSpWithParam2() {

	fruitParam := new(FruitParam)
	fruitParam.Code = "A2"
	f := []Fruit{}
	err := msdb.Select(&f, "up_Fruit_R2 @Code = ?", "A2")
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Printf("%+v", f)

}

func TestSpWithParam3() {

	fruitParam := new(FruitParam)
	fruitParam.Code = "A2"
	f := []Fruit{}
	spText := fmt.Sprintf("up_Fruit_R2 @Code = %v", "A2")
	err := msdb.Select(&f, spText)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Printf("%+v", f)

}

func TestSearch() {
	// Query the database, storing results in a []Person (wrapped in []interface{})
	fruits := []Fruit{}
	err2 := msdb.Select(&fruits, "select Name,Price,Color,Code,StoreCode from dbo.Fruit")
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

type QueryLogger struct {
	queryer sqlx.Queryer
	logger  *log.Logger
}

func (p *QueryLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println(query, args)
	return p.queryer.Query(query, args...)
}

func (p *QueryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	fmt.Println(query, args)
	return p.queryer.Queryx(query, args...)
}

func (p *QueryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	fmt.Println(query, args)
	return p.queryer.QueryRowx(query, args...)
}
