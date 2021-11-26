package adding

import (
	"context"
	"fmt"
	"os"

	a "test/pgx/pkg/album"

	"github.com/jackc/pgx/v4/pgxpool"
)

func AddAlbum(dbpool *pgxpool.Pool, album a.AlbumInput) (message string) {

	const createAlbumQuery = `
	INSERT INTO public.album(title, artist, rating)
	VALUES ($1, $2, $3);
	`
	_, err := dbpool.Exec(context.Background(),
		createAlbumQuery, album.Title, album.Artist, album.Rating)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	message = fmt.Sprintf("Album with the title of %v has been added to the database\n", album.Title)
	return message

}
