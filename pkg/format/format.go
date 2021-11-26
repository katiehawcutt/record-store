package format

import (
	"fmt"
	a "test/pgx/pkg/album"
)

func FormatAlbums(albums []a.Album) (formattedData []string) {
	var d []string
	for _, v := range albums {
		album := fmt.Sprintf("Id: %d, Title: %v, Artist: %v, Rating: %d\n",
			v.Id, v.Title, v.Artist, v.Rating)

		d = append(d, album)
	}
	return d
}
