package validator

import (
	a "test/pgx/pkg/album"

	v "github.com/go-ozzo/ozzo-validation"
)

func ValidateAlbum(album a.AlbumInput) error {
	return v.ValidateStruct(&album,
		v.Field(&album.Title, v.Required, v.Length(1, 50).Error("The title must be between 1 and 50 characters")),
		v.Field(&album.Artist, v.Required, v.Length(1, 50).Error("The artist must be between 1 and 50 characters")),
		v.Field(&album.Rating, v.Required, v.Min(1), v.Max(5).Error("The rating must be a number between 1 and 5")),
	)

}
