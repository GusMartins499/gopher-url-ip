package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Execute executes the CLI application.
func Execute() *cli.App {
	flags := []cli.Flag{
	cli.StringFlag{
		Name:  "host",
		Usage: "The URL to get the IP address for.",
	},
}

	app := cli.NewApp()
	app.Name = "gopher-url-ip"
	app.Usage = "A simple CLI tool to get the IP address of a given URL."

	app.Commands = []cli.Command{
		{
			Name:    "get-ip",
			Aliases: []string{"g"},
			Usage:   "Get the IP address of a given URL.",
			Flags:   flags,
			Action:  getIps,
		},
		{
			Name: "get-servers",
			Usage: "Get the IP addresses of the servers for a given URL.",
			Flags: flags,
			Action: getServers,
	},
}
	
	return app
}

func getServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func getIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
