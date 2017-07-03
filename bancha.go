package main

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"github.com/fatih/color"
)

var (
	port     string
	basePath string
)

func main() {

	app := cli.NewApp()
	app.Name = "bancha"

	app.Description = "Ban chá! A Simple HTTP Server with go. Debug/Log Porpose!"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "port, p",
			Value:       "1323",
			Usage:       "port for the server",
			Destination: &port,
		},
		cli.StringFlag{
			Name:        "basePath, b",
			Value:       "/",
			Usage:       "the base bath for the server",
			Destination: &basePath,
		},
	}

	app.Action = func(c *cli.Context) error {
		e := echo.New()
		e.Debug = true
		e.Any(basePath, printRequest)
		e.Start(":" + port)
		return nil
	}

	app.Run(os.Args)
}

func printRequest(c echo.Context) error {
	s, _ := ioutil.ReadAll(c.Request().Body)

	blue := color.New(color.BgBlue).SprintFunc()

	fmt.Printf("%s\n", blue("=============HEADERS============="))

	for k, v := range c.Request().Header {
		fmt.Printf("%s %s\n", k, v)
	}

	fmt.Printf("%s\n", blue("=============BODY============="))

	//fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))
	fmt.Printf("Json Received: %s\n", s)
	return c.String(204, "")
}
