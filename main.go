package main

import (
	"fmt"
	"log"

	"github.com/VictorCrespo/SISS-v2/database"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading .env file")
	}

	_, err = database.GetConnectionPool()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("Todo bien!")

}
