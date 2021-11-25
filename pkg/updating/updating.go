package updating

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func UpdateAlbum(dbpool *pgxpool.Pool, album []string) (message string) {

	id := album[0]
	title := album[1]
	artist := album[2]
	rating := album[3]

	const updateAlbumQuery = `
	UPDATE public.album
	SET title = $1, artist = $2, rating = $3
	WHERE id = $4;
	`

	_, err := dbpool.Exec(context.Background(), updateAlbumQuery, title, artist, rating, id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	message = fmt.Sprintf("Album with the id of %v has been updated\n", id)
	return message
}
