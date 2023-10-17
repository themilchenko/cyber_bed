package models

type SearchResponse struct {
	Data []struct {
		ID       int    `json:"id"`
		Slug     string `json:"slug"`
		ScName   string `json:"scientific_name"`
		ImageURL string `json:"image_url"`
	} `json:"data"`
}
