package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/jackal-dao/canine/app"
)

func main() {
	rootCmd, _ := NewRootCmd()

	rootCmd.AddCommand(StartServer())
	rootCmd.AddCommand(SubmitProof())

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
