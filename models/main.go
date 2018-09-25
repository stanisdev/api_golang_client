package models

import(
	"app/services"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	validator "github.com/asaskevich/govalidator"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"time"
)

var instance *gorm.DB

type DbMethods struct {
	db *gorm.DB
}

func DatabaseConnect() {
	params := viper.GetString("database.username") + ":" + viper.GetString("database.password") + "@/" + viper.GetString("database.dbname") + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", params)
  if err != nil {
    panic("Failed to connect database")
	}
	db.LogMode(true)
	fmt.Println("Database connected")
	instance = db
}

func DatabaseMigrate() {
	instance.AutoMigrate(&User{}) // Create User table

	salt := services.GenerateRandomString(10)
	hash, err := services.GetPasswordHash("pRCek5iFYm" + salt)
	if (err != nil) {
		panic(err)
	}
	user := User{}
	instance.Where(&User{Username: "mr.admin"}).Find(&user) // Create Admin record
	if (user.ID < 1) {
		if err := instance.Create(&User{Username: "mr.admin", Password: hash, Salt: salt, LastLogin: time.Now()}); err != nil {
			fmt.Println("An error occurred while creating the \"Admin\" entry")
			fmt.Println(err)
		}
	}
	instance.AutoMigrate(&Notification{}) // Create Notification table
	instance.AutoMigrate(&Company{}) // Create Company table
	instance.Model(&Notification{}).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")
}

func GetConnection() *gorm.DB {
	return instance
}

func ValidateModel(modelInstance interface{}) bool {
	_, err := validator.ValidateStruct(modelInstance)
	return err == nil
}