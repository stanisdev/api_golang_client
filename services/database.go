package services

import(
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

func DatabaseConnect() {
	params := viper.GetString("database.username") + ":" + viper.GetString("database.password") + "@/" + viper.GetString("database.dbname") + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", params)
  if err != nil {
    panic("Failed to connect database")
	}
	fmt.Println("Database connected")
  defer db.Close()
}
