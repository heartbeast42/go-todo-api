package main

import (
	_ "github.com/ClarkWalker/scorboard-api/database"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Migrations ...
func Migrations() {
	// dbName := os.Getenv("DATABASE_URL")
	// if dbName == "" {
	// 	dbName = "host=localhost dbname=ghq0 sslmode=disable"
	// }
	// fmt.Println("migrations db", dbName)
	//
	// db, err := gorm.Open("postgres", dbName)
	// if err != nil {
	// 	panic("Oh noos you got a database nopers")
	// }
	// defer db.Close()
	db := Connection()
	// Migrate the schema
	db.AutoMigrate(&Product{})

	// clear db

	// migrate schema

	//

} // end Migrations