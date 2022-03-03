package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eduardosnm/golang-make-http-request/internal"
)

const (
	APIEndpoint string = "https://swapi.dev/api/"
	APIResource string = "people/"
)

type characterRepo struct {
	url string
}

func NewRemoteRepository() internal.CharacterRepo {
	return &characterRepo{url: APIEndpoint}
}

func (p *characterRepo) GetCharacter(id string) (character *internal.Character, err error) {

	resp, err := http.Get(APIEndpoint + APIResource + id)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &character)
	if err != nil {
		return nil, err
	}

	character.Id = id

	return
}

func (p *characterRepo) AddCharacter(planet *internal.Character) error {
	fmt.Printf("Not implemented")
	panic(1)
}
