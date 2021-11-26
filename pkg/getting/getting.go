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

	const getAllAlbumsQuery = "SELECT * from public.album ORDER BY id ASC;"
	rows, err := dbpool.Query(context.Background(), getAllAlbumsQuery)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("Error while iterating dataset")
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

func GetAlbum(dbpool *pgxpool.Pool, albumId int32) (album []string) {

	var albumToFormat []a.Album
	var a a.Album

	const getAlbumByIdQuery = "SELECT * from public.album WHERE id = $1;"
	err := dbpool.QueryRow(context.Background(), getAlbumByIdQuery, albumId).Scan(&a.Id, &a.Title, &a.Artist, &a.Rating)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	albumToFormat = append(albumToFormat, a)
	formattedAlbum := format.FormatAlbums(albumToFormat)
	return formattedAlbum
}
