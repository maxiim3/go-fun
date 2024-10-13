package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func fire() {

}

func main() {
	// Clean the db file
	// os.Remove("./persons.db")

	// create new db
	db, err := sql.Open("sqlite3", "./persons.db")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS persons (id integer not null primary key, name text, age int, email text)`)
	defer db.Close()

	////////////////////// START SERVER
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	///////////////////// GET TEST index
	e.GET("/", func(c echo.Context) error {
		fmt.Println("TEST")
		return c.String(200, "WORKING")
	})

	////////////////////// GET ALL PERSONS
	e.GET("/all", func(c echo.Context) error {
		fmt.Println("Requesting all DATA")
		persons := QueryAllPersons(db)
		fmt.Println("Success")

		return c.JSON(200, persons)
	})

	///////////////////// CREATE NEW PERSON
	e.POST("/person", func(c echo.Context) error {
		fmt.Println("Requesting Add a new Person" )
		r := c.Request()

		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("error in POST", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Body Request"})
		}

		var newP Person
		err = newP.Scan(data)
		if err != nil {
			log.Fatal("error in POST", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Data"})
		}

		InsertPerson(newP.Name, uint(newP.Age), newP.Email, db)

		return c.JSON(200, map[string]string{
			"msg": "Successfly created a new person",
			"created": newP.Name,
		})
	})

	/////////////////////// START
	e.Logger.Fatal(e.Start(":8080"))
}
