package getting

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Album struct {
	Id     int32
	Title  string
	Artist string
	Rating int32
}

func GetAlbums(dbpool *pgxpool.Pool) (albums []Album) {

	var allAlbums []Album

	const getAllAlbumsQuery = "SELECT * from public.album;"

	rows, err := dbpool.Query(context.Background(), getAllAlbumsQuery)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		values, err := rows.Values()

		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		a := Album{
			Id:     values[0].(int32),
			Title:  values[1].(string),
			Artist: values[2].(string),
			Rating: values[3].(int32),
		}

		allAlbums = append(allAlbums, a)

	}

	return allAlbums

}

func GetAlbum(dbpool *pgxpool.Pool, albumId string) (album Album) {

	var id int32
	var title string
	var artist string
	var rating int32
	var a Album

	const getAlbumByIdQuery = "SELECT * from public.album WHERE id = $1;"

	err := dbpool.QueryRow(context.Background(), getAlbumByIdQuery, albumId).Scan(&id, &title, &artist, &rating)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	a = Album{
		Id:     id,
		Title:  title,
		Artist: artist,
		Rating: rating,
	}

	return a

}
