package entries

type Article struct {
	Id          int      `json:"id"`
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Author      string   `json:"author"`
	Timestamp   string   `json:"timestamp"`
	Category    Category `json:"category"`
}
