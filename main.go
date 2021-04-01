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
		Action: func(c *cli.Context) error {
			if c.NArg() < 2 {
				return fmt.Errorf("shortage args")
			}

			src := c.Args().Get(0)
			dhost := c.Args().Get(1)

			replaced, err := ReplaceHost(src, dhost)
			if err != nil {
				return err
			}
			fmt.Printf("%s", replaced)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
