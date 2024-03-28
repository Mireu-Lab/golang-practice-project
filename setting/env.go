package setting

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Envinfo() map[string]string {
	DB_Info := map[string]string{
		"Host":     "",
		"Port":     "",
		"User":     "",
		"Password": "",
	}

	err := godotenv.Load("./.env")

	if err != nil {
		log.Print(".env 파일을 찾을 수 없습니다.")
		return nil
	}

	DB_Info["Host"] = os.Getenv("RDB_Host")
	DB_Info["Port"] = os.Getenv("RDB_Port")
	DB_Info["User"] = os.Getenv("RDB_User")
	DB_Info["Password"] = os.Getenv("RDB_Passwd")

	fmt.Println(DB_Info["Host"], DB_Info["Port"], DB_Info["User"], DB_Info["Password"])

	if DB_Info["Host"] == "" || DB_Info["Port"] == "" || DB_Info["User"] == "" || DB_Info["Password"] == "" {
		log.Print("데이터가 없음")
		return nil
	}

	return DB_Info
}
