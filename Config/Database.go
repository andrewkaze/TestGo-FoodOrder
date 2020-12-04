package Config

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

type DBConfig struct {
	file string
}

func BuildDBConfig() *DBConfig{
	return &DBConfig{file: "restodb.db"}
}
func DbUrl(dbConfig *DBConfig) string{
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s", cwd, dbConfig.file)
}