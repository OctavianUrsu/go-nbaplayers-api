package gonbaplayersapi

type Player struct {
	PlayerId  int    `json:"playerId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	TeamId    int    `json:"teamId"`
}
