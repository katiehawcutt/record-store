package updating

import (
	"context"
	"fmt"
	"os"

	a "test/pgx/pkg/album"
	v "test/pgx/pkg/validator"

	"github.com/jackc/pgx/v4/pgxpool"
)

func UpdateAlbum(dbpool *pgxpool.Pool, albumId int32, album a.AlbumInput) (message string) {

	validationErr := v.ValidateAlbum(album)

	if validationErr != nil {
		fmt.Println(validationErr)
		os.Exit(1)
	}

	const updateAlbumQuery = `
	UPDATE public.album
	SET title = $1, artist = $2, rating = $3
	WHERE id = $4;
	`

	_, err := dbpool.Exec(context.Background(),
		updateAlbumQuery, album.Title, album.Artist, album.Rating, albumId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	message = fmt.Sprintf("Album with the id of %v has been updated\n", albumId)
	return message
}
