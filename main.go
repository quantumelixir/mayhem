package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"

	"github.com/codegangsta/cli"
)

type Level struct{
	Board [][]Color
	Bots []Robot
	Main string
}

func (l Level) Marshal() ([]byte, error) {
	return json.MarshalIndent(l, "", "\t")
}

// * Main
func main() {

	var filename string

	app := cli.NewApp()
	app.Name = "mayhem"
	app.Usage = "the attack of the clones !!!"
	app.Version = "0.1"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "load-spec-from-file, l",
			Value: "",
			Usage: "load spec from file",
			Destination: &filename,
		},
	}

	app.Action = func(c *cli.Context) {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("error:", err)
		} else {
			var l Level
			json.Unmarshal(b, &l)
			Run(l.Board, l.Main, l.Bots)
			fmt.Println(l.Bots[0].X, l.Bots[0].Y, l.Bots[0].D)
		}
	}		

	app.Run(os.Args)
}
