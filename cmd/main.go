package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"test/pgx/pkg/adding"
	c "test/pgx/pkg/convert"
	"test/pgx/pkg/deleting"
	"test/pgx/pkg/format"
	"test/pgx/pkg/getting"
	"test/pgx/pkg/updating"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Welcome to Katie's Records!")
		os.Exit(1)
	}
	var operation = os.Args[1]
	var input = os.Args[2:]

	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	switch operation {
	case "getAlbums":
		albums := getting.GetAlbums(dbpool)
		fmt.Print(albums, "\n")

	case "getAlbum":
		albumId := c.ConvertStringToInt32(input[0])
		album := getting.GetAlbum(dbpool, albumId)
		fmt.Print(album, "\n")

	case "addAlbum":
		albumToAdd := format.FormatInputData(input)
		message := adding.AddAlbum(dbpool, albumToAdd)
		fmt.Print(message)

	case "updateAlbum":
		albumId := c.ConvertStringToInt32(input[0])
		albumDetails := format.FormatInputData(input[1:])
		message := updating.UpdateAlbum(dbpool, albumId, albumDetails)
		fmt.Print(message)

	case "deleteAlbum":
		albumId := c.ConvertStringToInt32(input[0])
		message := deleting.DeleteAlbum(dbpool, albumId)
		fmt.Print(message)

	default:
		fmt.Println("Please enter a valid input")
	}

}
