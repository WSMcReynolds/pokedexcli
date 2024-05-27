package pokeapi

type locationResponse struct {
	Count    int               `json:"count"`
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Results  []locationResults `json:"results"`
}

type locationResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
