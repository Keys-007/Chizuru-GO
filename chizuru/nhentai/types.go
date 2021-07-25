package nhentai

type NHentaiResponse struct {
	ID           int    `json:"id"`
	MediaID      string `json:"media_id"`
	Title        Title  `json:"title"`
	Images       Images `json:"images"`
	Scanlator    string `json:"scanlator"`
	UploadDate   int    `json:"upload_date"`
	Tags         []Tags `json:"tags"`
	NumPages     int    `json:"num_pages"`
	NumFavorites int    `json:"num_favorites"`
}
type Title struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Pretty   string `json:"pretty"`
}
type Pages struct {
	T string `json:"t"`
	W int    `json:"w"`
	H int    `json:"h"`
}
type Cover struct {
	T string `json:"t"`
	W int    `json:"w"`
	H int    `json:"h"`
}
type Thumbnail struct {
	T string `json:"t"`
	W int    `json:"w"`
	H int    `json:"h"`
}
type Images struct {
	Pages     []Pages   `json:"pages"`
	Cover     Cover     `json:"cover"`
	Thumbnail Thumbnail `json:"thumbnail"`
}
type Tags struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Count int    `json:"count"`
}
