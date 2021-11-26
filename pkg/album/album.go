package album

type Album struct {
	Id     int32
	Title  string
	Artist string
	Rating int32
}

type AlbumInput struct {
	Title  string
	Artist string
	Rating int32
}
