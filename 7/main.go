package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
}

type LogEntry struct {
	User User   `json:"user"`
	Dist string `json:"dist"`
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Src   string `json:"src"`
	Time  string `json:"time"`
}

const (
	tableNameUser = "users"
)

func main() {
	// .envファイルから環境変数を読み込む
	err := loadEnv()
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	// データベースに接続
	Db, err := connectDB()
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// テーブルを作成
	err = createTable(Db)
	if err != nil {
		log.Panicln("Error creating table:", err)
	}

	// sample.logファイルを開く
	file, err := os.Open("sample.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// トランザクションの開始
	tx, err := Db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// ファイルをスキャンするためのbufio.Scannerを作成
	scanner := bufio.NewScanner(file)
	id := 0
	// 各行を処理
	for scanner.Scan() {
		// JSONデータを取得
		jsonStr := scanner.Text()

		// JSONを構造体にデコード
		var entry LogEntry
		if err := json.Unmarshal([]byte(jsonStr), &entry); err != nil {
			log.Println("error decoding JSON:", err)
			continue
		}

		// User型に変換
		var user User
		user.Age = entry.User.Age
		user.Name = entry.User.Name
		user.Role = entry.User.Role
		user.ID = id
		id += 1

		// ユーザーをデータベースに挿入
		_, err := tx.Exec("INSERT INTO users (id, name, age, role) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Age, user.Role)
		if err != nil {
			tx.Rollback() // エラーがあればロールバック
			log.Fatal(err)
		}
	}

	// トランザクションのコミット
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	log.Println("Users inserted successfully")

	// スキャン中にエラーが発生した場合のエラーチェック
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//環境変数を読み込む
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

// データベースに接続
func connectDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

// テーブルを作成
func createTable(db *sql.DB) error {
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		age INTEGER,
		name VARCHAR(500),
		role CHAR(15))`, tableNameUser)
	_, err := db.Exec(cmdT)
	if err != nil {
		return err
	}
	log.Println("Table 'users' created successfully")
	return nil
}