package api

import "strconv"

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
	Country   string
	Id        string
	FirstName string
	LastName  string
	Name      string
	Position  string
	Number    int
	Age       *StringedInt
}

func NewStringedInt(value int) *StringedInt {
	return &StringedInt{value}
}

type StringedInt struct {
	int
}

func (s *StringedInt) Int() int {
	return s.int
}

func (s *StringedInt) UnmarshalJSON(data []byte) error {
	var stringValue string
	if data[0] == '"' && data[len(data) - 1] == '"' {
		stringValue = string(data[1 : len(data) - 1])
	} else {
		stringValue = string(data)
	}

	intValue, err := strconv.Atoi(stringValue)
	if nil != err {
		return err
	}

	s.int = intValue
	return nil
}
