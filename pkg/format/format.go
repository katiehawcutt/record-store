package format

import (
	"fmt"

	a "test/pgx/pkg/album"
	c "test/pgx/pkg/convert"
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

func FormatInputData(input []string) (album a.AlbumInput) {
	title := input[0]
	artist := input[1]
	rating := input[2]

	// CLI arguments in Golang (os.Args) are always passing as a string so
	// we need to convert the string to the corresponding int32 value
	ratingAsInt32 := c.ConvertStringToInt32(rating)

	a := a.AlbumInput{
		Title:  title,
		Artist: artist,
		Rating: ratingAsInt32,
	}

	return a
}
