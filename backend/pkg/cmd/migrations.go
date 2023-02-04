package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/IN-45/INT20H-test-task/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

func NewMigrationsCommand(fxOptions []fx.Option) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					fxOptions = append(
						fxOptions,
						fx.Invoke(func(db *bun.DB) error {
							migrator := migrate.NewMigrator(db, migrations.Migrations)
							ctx := context.Background()

							if err := migrator.Init(ctx); err != nil {
								return err
							}

							_, err := migrator.Migrate(ctx)

							return err
						}),
					)
					app := fx.New(fxOptions...)

					return app.Start(context.Background())
				},
			},
			{
				Name:  "create",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					fxOptions = append(
						fxOptions,
						fx.Invoke(func(db *bun.DB) error {
							migrator := migrate.NewMigrator(db, migrations.Migrations)
							name := strings.Join(c.Args().Slice(), "_")

							files, err := migrator.CreateSQLMigrations(c.Context, name)
							if err != nil {
								return err
							}

							for _, mf := range files {
								fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
							}

							return nil
						}),
					)
					app := fx.New(fxOptions...)

					return app.Start(context.Background())
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					fxOptions = append(
						fxOptions,
						fx.Invoke(func(db *bun.DB) error {
							migrator := migrate.NewMigrator(db, migrations.Migrations)
							return migrator.Unlock(c.Context)
						}),
					)
					app := fx.New(fxOptions...)

					return app.Start(context.Background())
				},
			},
		},
	}
}
