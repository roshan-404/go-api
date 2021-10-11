//main.go
package main

import (
	Config "api/config"
	Models "api/models"
	Routes "api/routes"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)
var err error
func main() {
 Config.DB, err = gorm.Open("postgres", Config.DbURL(Config.BuildDBConfig()))
if err != nil {
  fmt.Println("Status:", err)
 }
defer Config.DB.Close()
 Config.DB.AutoMigrate(&Models.User{})
r := Routes.SetupRouter()
 //running
 r.Run()
}