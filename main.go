package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// #region city
type City struct {
	ID          int    `json:"ID,omitempty" db:"ID"`
	Name        string `json:"name,omitempty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

var (
	db *sqlx.DB
)

// #endregion city
func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	//cityInput := os.Args[1]

	conf := mysql.Config{
		User:      os.Getenv("DB_USERNAME"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOSTNAME") + ":" + os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_DATABASE"),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conntected")
	db = _db
	
	e := echo.New()

	e.GET("/cities/:cityName", getCityInfoHandler)

	e.POST("/cities", insertCityHandler)

	e.Start(":8080")		

	/*
	// #region get
	var city City
	var populCountry int
	err = db.Get(&city, "SELECT * FROM city WHERE Name = ?", cityInput)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such city Name = '%s'\n", cityInput)
	} else if err != nil {
		log.Fatalf("DB Error: %s\n", err)
	}
	// #endregion get
	fmt.Printf(cityInput + "の人口は%d人です\n", city.Population)

	err = db.Get(&populCountry, "SELECT Population FROM country WHERE Code = ?", city.CountryCode)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such city Name = '%s'\n", cityInput)
	} else if err != nil {
		log.Fatalf("DB Error: %s\n", err)
	}
	percent := float64(city.Population)/float64(populCountry)*100
	fmt.Printf("%f%%の人口がいます\n",percent)
	*/
	/*
	var cities []City
	err = db.Select(&cities, "SELECT * FROM city WHERE CountryCode = 'JPN'") //?を使わない場合、第3引数以降は不要
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("日本の都市一覧")
	for _, city := range cities {
		fmt.Printf("都市名: %s, 人口: %d\n", city.Name, city.Population)
	}
	*/
}

func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	var city City
	if err := db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No such city Name = %s", cityName))
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	return c.JSON(http.StatusOK, city)
}

func insertCityHandler(c echo.Context) error {
	newCity := &City{}
	err := c.Bind(newCity)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}

	result ,err := db.Exec("INSERT INTO city (Name,CountryCode,District,Population) VALUES (?,?,?,?)", newCity.Name, newCity.CountryCode, newCity.District, newCity.Population)
	if err != nil {
		log.Fatalf("Faild to insert city")
	}

	id, _ := result.LastInsertId()
	newCity.ID = int(id)

	return c.JSON(http.StatusOK, newCity)
}