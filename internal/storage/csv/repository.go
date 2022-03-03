package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/eduardosnm/golang-make-http-request/internal"
)

type characterRepo struct {
}

func NewCsvRepository() internal.CharacterRepo {
	return &characterRepo{}
}

func (p *characterRepo) AddCharacter(planet *internal.Character) error {
	f, err := os.Create("planet." + planet.Id + ".csv")
	defer f.Close()

	csvWriter := csv.NewWriter(f)

	csvWriter.Write(planet.ToArray())
	csvWriter.Flush()

	return err
}

func (p *characterRepo) GetCharacter(id string) (planet *internal.Character, err error) {
	fmt.Printf("Not implemented")
	panic(1)
}
