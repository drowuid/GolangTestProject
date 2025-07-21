package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./jokes.db")
	return err
}

type Joke struct {
	ID int
	Setup string
	Punchline string
	Category string
	Likes int
	Dislikes int
}

func GetAllJokes() ([]Joke, error) {
	rows, err := DB.Query("SELECT id, setup, punchline, category, likes, dislikes FROM jokes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jokes []Joke
	for rows.Next() {
		var j Joke
		rows.Scan(&j.ID, &j.Setup, &j.Punchline, &j.Category, &j.Likes, &j.Dislikes)
		jokes = append(jokes, j)
		}
	return jokes, nil
}

func InsertJoke(setup, punchline, category string) error {
	_, err := DB.Exec("INSERT INTO jokes (setup, punchline, category) VALUES (?, ?, ?)", setup, punchline, category)
	return err
}

func GetJokeByID(id int) (Joke, error) {
	var j Joke
	row := DB.QueryRow("SELECT id, setup, punchline, category, likes, dislikes FROM jokes WHERE id = ?", id)
	err := row.Scan(&j.ID, &j.Setup, &j.Punchline, &j.Category, &j.Likes, &j.Dislikes)
	return j, err
}

func UpdateJoke(id int, setup, punchline, category string) error {
	_, err := DB.Exec("UPDATE jokes SET setup = ?, punchline = ?, category = ?  WHERE id = ?", setup, punchline, category, id)
	return err
}

func DeleteJoke(id int) error {
	_, err := DB.Exec("Delete from jokes WHERE id = ?", id)
	return err
}
