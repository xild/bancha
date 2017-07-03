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
	blue = color.New(color.BgBlue).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
)

const (
	name        = "Ban Chá! ʕ•ᴥ•ʔ"
	version     = "1.0.0"
	description = "Ban chá! A Simple HTTP Server with go. Debug/Log Porpose!"
	banner      = `
██████╗  █████╗ ███╗   ██╗ ██████╗██╗  ██╗ █████╗
██╔══██╗██╔══██╗████╗  ██║██╔════╝██║  ██║██╔══██╗
██████╔╝███████║██╔██╗ ██║██║     ███████║███████║
██╔══██╗██╔══██║██║╚██╗██║██║     ██╔══██║██╔══██║
██████╔╝██║  ██║██║ ╚████║╚██████╗██║  ██║██║  ██║
╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝
Simple http server to log!
Powered by Luis Vieira ʕ•ᴥ•ʔ
                                                  `
)

func main() {

	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Description = description
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "port, p",
			Value:       "0805",
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
	fmt.Printf("%s\n", green(banner))


	app.Action = func(c *cli.Context) error {
		e := echo.New()
		e.HideBanner = true
		e.Debug = true
		fmt.Printf("base path: %s\n", green(basePath))
		fmt.Printf("serving port: %s\n", green(port))
		e.Any(basePath, printRequest)
		e.Start(":" + port)
		return nil
	}

	app.Run(os.Args)
}

func printRequest(c echo.Context) error {
	s, _ := ioutil.ReadAll(c.Request().Body)

	fmt.Printf("%s\n", blue("=============HEADERS============="))

	for k, v := range c.Request().Header {
		fmt.Printf("%s %s\n", k, v)
	}

	fmt.Printf("%s\n", blue("=============BODY============="))

	//fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))
	fmt.Printf("Json Received: %s\n", s)
	return c.String(204, "")
}
