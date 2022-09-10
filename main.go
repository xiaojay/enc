package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "enc",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "input", Usage: "input file", Aliases: []string{"i"}},
			&cli.StringFlag{Name: "output", Usage: "output file", Aliases: []string{"o"}},
			&cli.StringFlag{Name: "password", Usage: "password", Aliases: []string{"p"}},
			&cli.BoolFlag{Name: "decode", Value: false, Usage: "decode file or not", Aliases: []string{"d"}},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) (err error) {
	file, err := os.Open(c.String("input"))
	if err != nil {
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	s := string(data)
	if c.Bool("decode") {
		s = Decrypt(c.String("password"), s)
	} else {
		s = Encrypt(c.String("password"), s)
	}

	by := []byte(s)
	ioutil.WriteFile(c.String("output"), by, 0644)
	return

}
