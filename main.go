package main

import (
	"log"
	"os"

	"github.com/Jetereting/swaggo/parser"
	"github.com/urfave/cli"
)

const (
	version = "v1.0.1"
)

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Name = "swaggo"
	app.HelpName = "swaggo"
	app.Usage = "a utility for convert go annotations to swagger-doc"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dev, d",
			Usage: "develop mode",
		},
		cli.StringFlag{
			Name:  "swagger, s",
			Value: "./swagger.go",
			Usage: "where is the swagger.go file",
		},
		cli.StringFlag{
			Name:  "project, p",
			Value: "./",
			Usage: "where is the project",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "./",
			Usage: "the output of the swagger file that was generated",
		},
		cli.StringFlag{
			Name:  "type, t",
			Value: "json",
			Usage: "the type of swagger file (json or yaml)",
		},
	}
	app.Action = func(c *cli.Context) error {
		if err := parser.Parse(c.String("project"),
			c.String("swagger"),
			c.String("output"),
			c.String("type"),
			c.Bool("dev")); err != nil {
			return err
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Printf("[Error] %v", err)
	}
}
