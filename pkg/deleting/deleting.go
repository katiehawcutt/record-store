package deleting

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func DeleteAlbum(dbpool *pgxpool.Pool, albumId string) (message string) {

	const deleteAlbumByIdQuery = "DELETE FROM public.album WHERE id = $1;"

	commandTag, err := dbpool.Exec(context.Background(), deleteAlbumByIdQuery, albumId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	if commandTag.RowsAffected() != 1 {
		fmt.Println("No row found to delete")
		os.Exit(1)
	}

	message = fmt.Sprintf("Album with the id of %v has been deleted\n", albumId)
	return message

}
