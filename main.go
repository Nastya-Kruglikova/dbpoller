package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func readConfig() {
	viper.SetConfigName("test")
	viper.AddConfigPath("/Users/nastya/Projects/golang/dbpoller")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s\n", err)
	}
}

func main() {
	readConfig()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.Get("db_user"), viper.Get("db_password"), viper.Get("db_name"))
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	for {
		var result int
		err = db.QueryRow("SELECT 1").Scan(&result)
		checkErr(err)
		fmt.Printf("Cleaned up %d sessions\n", result)
		time.Sleep(100 * time.Millisecond)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
