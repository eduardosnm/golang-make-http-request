package main

import (
	"github.com/eduardosnm/golang-make-http-request/internal"
	"github.com/eduardosnm/golang-make-http-request/internal/cli"
	"github.com/eduardosnm/golang-make-http-request/internal/storage/csv"
	"github.com/eduardosnm/golang-make-http-request/internal/storage/remote"
	"github.com/spf13/cobra"
)

func main() {

	var remoteRepository internal.CharacterRepo
	remoteRepository = remote.NewRemoteRepository()

	var csvRepository internal.CharacterRepo
	csvRepository = csv.NewCsvRepository()

	rootCmd := &cobra.Command{Use: "starwars-cli"}
	rootCmd.AddCommand(cli.InitPlanetsCmd(remoteRepository, csvRepository))
	rootCmd.Execute()
}
