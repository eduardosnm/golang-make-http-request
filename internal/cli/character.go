package cli

import (
	"fmt"

	. "github.com/eduardosnm/golang-make-http-request/internal"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

const idFlag string = "id"

func InitPlanetsCmd(remoteRepository CharacterRepo, csvRepository CharacterRepo) *cobra.Command {
	planetsCmd := &cobra.Command{
		Use:   "characters",
		Short: "Save character to csv file",
		Run:   runPlanetsFn(remoteRepository, csvRepository),
	}

	planetsCmd.Flags().StringP(idFlag, "i", "", "id of the character")

	return planetsCmd
}

const APIEndpoint string = "https://swapi.dev/api/"
const APIResource string = "people/"

func runPlanetsFn(remoteRepository CharacterRepo, csvRepository CharacterRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		IDResource, _ := cmd.Flags().GetString(idFlag)

		if IDResource == "" {
			IDResource = "1"
		}

		character, err := remoteRepository.GetCharacter(IDResource)
		if err != nil {
			fmt.Println("There was an error getting the character " + IDResource + " remotely")
		}

		err = csvRepository.AddCharacter(character)
		if err != nil {
			fmt.Println("There was an error saving the planet " + IDResource + "to disk")
		}
	}
}
