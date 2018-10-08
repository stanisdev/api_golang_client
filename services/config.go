package services

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
	"path/filepath"
	"path"
)

var dynamicConfig map[string]string

func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./src/app")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}	
}

func SetDynamicConfig() {
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	if (rootDir[len(rootDir)-4:] != "/bin") {
		panic("Wrong env structure")
	}
	rootDir = rootDir[0:len(rootDir)-4]
	appDir := path.Join(rootDir, "src", "app")

	dynamicConfig = map[string]string{
		"RootDir": rootDir,
		"AppDir": appDir,
		"UploadsDir": path.Join(appDir, "uploads"),
	}
}

func GetDynamicConfig() map[string]string {
	return dynamicConfig
}