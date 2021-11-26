package getting

import (
	"context"
	"fmt"
	"log"
	"os"
	a "test/pgx/pkg/album"
	"test/pgx/pkg/format"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetAlbums(dbpool *pgxpool.Pool) (albums []string) {
	var albumsToFormat []a.Album
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

		a := a.Album{
			Id:     values[0].(int32),
			Title:  values[1].(string),
			Artist: values[2].(string),
			Rating: values[3].(int32),
		}

		albumsToFormat = append(albumsToFormat, a)
	}
	formattedAlbums := format.FormatAlbums(albumsToFormat)
	return formattedAlbums
}

func GetAlbum(dbpool *pgxpool.Pool, albumId string) (album []string) {
	var id int32
	var title string
	var artist string
	var rating int32
	var albumToFormat []a.Album

	const getAlbumByIdQuery = "SELECT * from public.album WHERE id = $1;"

	err := dbpool.QueryRow(context.Background(), getAlbumByIdQuery, albumId).Scan(&id, &title, &artist, &rating)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	a := a.Album{
		Id:     id,
		Title:  title,
		Artist: artist,
		Rating: rating,
	}

	albumToFormat = append(albumToFormat, a)
	formattedAlbum := format.FormatAlbums(albumToFormat)
	return formattedAlbum
}
