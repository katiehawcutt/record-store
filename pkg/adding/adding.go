package adding

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func AddAlbum(dbpool *pgxpool.Pool, album []string) (message string) {

	title := album[0]
	artist := album[1]
	rating := album[2]

	const createAlbumQuery = `
	INSERT INTO public.album(title, artist, rating)
	VALUES ($1, $2, $3);
	`

	_, err := dbpool.Exec(context.Background(), createAlbumQuery, title, artist, rating)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	message = fmt.Sprintf("Album with the title of %v has been added to the database\n", title)
	return message

}
