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
				Name:  "human",
				Usage: "human-readable sizes (auto-select unit)",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			human := c.Bool("human")
			path := c.Args().First()
			if path == "" {
				path = "."
			}
			result, err := goproject242.GetPathSize(path, false, human, false)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
			return nil

		},
	}

	app.Run(context.Background(), os.Args)
}
