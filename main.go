package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "reho",
		Usage: "replace host of URL",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "location",
				Aliases: []string{"L"},
				Value:   false,
				Usage:   "location redirect",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() < 2 {
				return fmt.Errorf("shortage args")
			}

			src := c.Args().Get(0)
			dhost := c.Args().Get(1)

			l := c.Bool("location")

			replaced, err := ReplaceHost(src, dhost, l)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", replaced)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
