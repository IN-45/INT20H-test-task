package main

import (
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/IN-45/INT20H-test-task/di"
	"github.com/IN-45/INT20H-test-task/pkg/cmd"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

//	@title		Backend API
//	@host		localhost:8000
//	@BasePath	/

func main() {
	if err := godotenv.Load("./config/.env"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err)
	}

	fxOptions := di.ProvideModules()

	app := &cli.App{
		Name: "backend",
		Commands: []*cli.Command{
			cmd.NewServerStartCommand(fxOptions),
			cmd.NewMigrationsCommand(fxOptions),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
