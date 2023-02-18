package cities

type CreateCity struct {
	Name string `json:"cities"`
}

type CityUpdate struct {
	ID   int
	Name string `json:"name"`
}

type City struct {
	ID int
}
