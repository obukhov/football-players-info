package api

type Response struct {
	Status  string
	Code    int
	Message string
}

type TeamResponse struct {
	Response
	Data struct {
		     Team *Team
	     }
}

type Team struct {
	Id      int
	Name    string
	Players []Player
}

type Player struct {
	Country      string
	Id           string
	FirstName    string
	LastName     string
	Name         string
	Position     string
	Number       int
	BirthDate    string
	Age          string
	Height       int
	Weight       int
	ThumbnailSrc string
}
