package datastore

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS users(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE
    );
    CREATE TABLE IF NOT EXISTS tweetstore(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        tweet TEXT NOT NULL UNIQUE,
        FOREIGN KEY(id) REFERENCES users(id)
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) CreateUser(userName string) error {
	query := `INSERT INTO users(name) VALUES (?)`

	_, err := r.db.Exec(query, userName)
	return err
}

func (r *SQLiteRepository) GetTweetsFromUser(userName string) ([]string, error) {
	query := `SELECT tweet
		FROM tweetstore 
		    INNER JOIN users ON users.id = tweetstore.user_id
		WHERE users.name = ?`

	rows, err := r.db.Query(query, userName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tweets := []string{}

	for rows.Next() {
		var tweet string
		rows.Scan(&tweet)
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}

func (r *SQLiteRepository) CreateTweet(userName string, tweet string) error {
	query := `SELECT id FROM users WHERE name = ?`

	row := r.db.QueryRow(query, userName)
	var id int
	row.Scan(&id)

	query = `INSERT INTO tweetstore(user_id, tweet) VALUES (?, ?)`
	_, err := r.db.Exec(query, id, tweet)
	return err
}

func CreateSqliteDB(fileName string) *sql.DB {
	os.Remove(fileName)
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Setup(fileName string) *SQLiteRepository {
	db := CreateSqliteDB(fileName)
	repo := NewSQLiteRepository(db)
	if err := repo.Migrate(); err != nil {
		log.Fatal(err)
	}

	return repo
}
