package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	tableNameUser = "users"
)

func main() {
	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	// 環境変数からデータベース接続情報を取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// データベース接続URLを生成
	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	// データベースに接続
	Db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// データベースに接続できるか確認
	err = Db.Ping()
	if err != nil {
		log.Panicln("Database ping failed:", err)
	}
	log.Println("Successfully connected to the database")

	// テーブルを作成
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		age INTEGER,
		name VARCHAR(500),
		role CHAR(15))`, tableNameUser)
	_, err = Db.Exec(cmdT)
	if err != nil {
		log.Panicln("Error creating table:", err)
	}

	log.Println("Table 'users' created successfully")
}
