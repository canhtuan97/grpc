package structs


// Location is a single location.
type ResponseLocation struct {
	Areas []Areas
	GameIndices []GameIndices
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []Names
	Region Region
}
type Areas struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language Languages
	Name string `json:"name"`
}
type Languages struct {
	Name string `json:"name"`
	URL  string `json:"url"`

}
type Region struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationService interface {
	GetLocation() (*ResponsePokemon ,error)
}

type Location struct {
	client *Client
}


func (s *Location) GetLocation() (*ResponsePokemon , error) {

}