package database

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_Username string
var db_Password string
var db_Name string
var db_Host string
var db_Port string
var Db *gorm.DB

func getenv(ev, defval string) string {
    if  value, ok := os.LookupEnv(ev); ok {
       return value
    }
    return defval
}


func InitDb() *gorm.DB {
        db_Username = getenv("DB_USERNAME", "root")
        db_Password = getenv("DB_PW", "admin")
        db_Name = getenv("DB_NAME", "my_db")
        db_Host = getenv("DB_HOST", "127.0.0.1")
        db_Port = getenv("DB_PORT", "3306")
	Db = connectDB()
	return Db
}

func connectDB() (*gorm.DB) {
	var err error
	dsn := db_Username +":"+ db_Password +"@tcp"+ "(" + db_Host + ":" + db_Port +")/" + db_Name + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
