package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	_urlShortHttpDelivery "github.com/budhip/url-shortener/delivery/http"
	_urlShortRepo "github.com/budhip/url-shortener/repository/mysql"
	_urlShortUCase "github.com/budhip/url-shortener/usecase"

	_ "github.com/go-sql-driver/mysql"
)


func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("connected to database")

	r := mux.NewRouter()

	urlShortRepo := _urlShortRepo.NewMysqlUrlShortenerRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	urlShortU := _urlShortUCase.NewUrlShortenerUseCase(urlShortRepo, timeoutContext)

	_urlShortHttpDelivery.NewUrlShortenerHandler(r, urlShortU)

	// Start HTTP server.
	log.Println("Server started. Listening on port: ", viper.GetString("server.address"))
	errServe := http.ListenAndServe(viper.GetString("server.address"), r)
	if errServe != nil {
		log.Println("server stopped!")
		log.Fatal(errServe)
	}
}