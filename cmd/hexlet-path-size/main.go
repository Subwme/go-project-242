package main

import (
	goproject242 "code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name: "hexlet-path-size",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"hm"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			human := c.Bool("human")
			all := c.Bool("all")

			path := c.Args().First()
			if path == "" {
				path = "."
			}

			result, err := goproject242.GetPathSize(path, false, human, all)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
			return nil
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}
}
