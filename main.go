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

	app := cli.NewApp()
	app.Name = "mayhem"
	app.Usage = "the attack of the clones !!!"
	app.Version = "0.1"

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Println("specify a level spec file")
			return nil
		}
		b, err := ioutil.ReadFile(c.Args().Get(0))
		if err != nil {
			fmt.Println(err)
		} else {
			var l Level
			json.Unmarshal(b, &l)
			Run(l.Board, l.Main, l.Bots)
			fmt.Println(l.Bots[0].X, l.Bots[0].Y, l.Bots[0].D)
		}
		return err
	}		

	app.Run(os.Args)
}
