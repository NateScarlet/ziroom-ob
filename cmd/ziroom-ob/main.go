package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/NateScarlet/ziroom-ob/pkg/db"
	"github.com/NateScarlet/ziroom-ob/pkg/email"
	"github.com/NateScarlet/ziroom-ob/pkg/watch"
)

func main() {
	godotenv.Load()
	dbPath := os.Getenv("DATABASE_DIR")
	if dbPath == "" {
		dbPath = "/tmp/ziroom-ob/db"
	}

	err := email.Send("服务启动", time.Now().Format(time.RFC3339))
	if err != nil {
		panic(err)
	}
	log.Print("已发送测试邮件, 请检查收件箱确保能收到通知。")

	os.MkdirAll(dbPath, 0644)
	db := db.Connect(dbPath)
	defer db.Close()
	watch.StartAll(db)
	select {}
}
