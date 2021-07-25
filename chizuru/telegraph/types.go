package telegraph

type CreateAccount struct {
	Ok     bool      `json:"ok"`
	Result AccResult `json:"result"`
}
type AccResult struct {
	ShortName   string `json:"short_name"`
	AuthorName  string `json:"author_name"`
	AuthorURL   string `json:"author_url"`
	AccessToken string `json:"access_token"`
	AuthURL     string `json:"auth_url"`
}

type CreatePage struct {
	Ok     bool             `json:"ok"`
	Result CreatePageResult `json:"result"`
}
type Content struct {
	Tag      string   `json:"tag"`
	Children []string `json:"children"`
}
type CreatePageResult struct {
	Path        string    `json:"path"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AuthorName  string    `json:"author_name"`
	AuthorURL   string    `json:"author_url"`
	Content     []Content `json:"content"`
	Views       int       `json:"views"`
	CanEdit     bool      `json:"can_edit"`
}
