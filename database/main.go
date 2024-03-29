package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.Config{
		User:      "user",
		Passwd:    "password",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "recordings",
		Collation: "utf8mb4_general_ci",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	makeSelect()

	var newAlbum Album
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Title: ")
	scanner.Scan()
	newAlbum.Title = scanner.Text()
	fmt.Print("Artist: ")
	scanner.Scan()
	newAlbum.Artist = scanner.Text()
	fmt.Print("Price: ")
	scanner.Scan()
	priceStr := scanner.Text()
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		log.Fatal(err)
	}
	newAlbum.Price = float32(price)

	rowNum, err := insertNewAlbum(newAlbum)
	if err != nil {
		log.Fatal(err)
	}

	makeSelect()

	fmt.Print(rowNum)

}

func makeSelect() {
	albums, err := albumByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	for i := range albums {
		fmt.Printf("Album found: %v\n", albums[i])
	}
	fmt.Printf("Albums found: %v\n", albums)
}

func albumByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("select * from album where artist = ?", name)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil

}

func insertNewAlbum(album Album) (int64, error) {

	ctx := context.Background()

	insertQuery := "INSERT INTO album (title, artist, price) values (?, ?, ?)"

	stmt, err := db.PrepareContext(ctx, insertQuery)
	if err != nil {
		return 0, fmt.Errorf("insertError: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("insertError: %v", err)
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insertError: %v", err)
	}

	return lastInsertedId, nil

}
