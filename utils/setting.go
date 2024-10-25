package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	GINMODE string
	GINPORT string

	DBDRIVER   string
	DBHOST     string
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	LoadGinServer(file)
	LoadDatabase(file)
}
func LoadGinServer(file *ini.File) {
	GINMODE = file.Section("GinServer").Key("GINMODE").MustString("debug")
	GINPORT = file.Section("GinServer").Key("GINPORT").MustString(":3000")
}
func LoadDatabase(file *ini.File) {
	DBDRIVER = file.Section("Database").Key("DBDRIVER").MustString("mysql")
	DBHOST = file.Section("Database").Key("DBHOST").MustString("localhost")
	DBPORT = file.Section("Database").Key("DBPORT").MustString("3306")
	DBUSER = file.Section("Database").Key("DBUSER").MustString("user")
	DBPASSWORD = file.Section("Database").Key("DBPASSWORD").MustString("password")
	DBNAME = file.Section("Database").Key("DBNAME").MustString("dbname")

}
