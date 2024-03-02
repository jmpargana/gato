package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func readWrite(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(data)
	return err
}

func readWriteBuffered(r io.Reader) error {
	reader := bufio.NewReader(r)
	w := bufio.NewWriter(os.Stdout)
	if _, err := reader.WriteTo(w); err != nil {
		return err
	}
	return w.Flush()
}

func readWriteNumbered(r io.Reader) error {
	yellow := color.New(color.FgYellow).Printf
	sc := bufio.NewScanner(r)
	i := 1
	for ; sc.Scan(); i++ {
		line := sc.Text()
		yellow(fmt.Sprintf("%d: ", i))
		os.Stdout.Write([]byte(line))
		os.Stdout.Write([]byte{'\n'})
	}
	return nil
}

func parseArgs(ctx *cli.Context) []string {
	args := ctx.Args()
	if args.Len() == 0 || args.Len() == 1 && args.First() == "-" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		return strings.Split(string(data), "\n")
	}

	return args.Slice()
}

func main() {
	app := &cli.App{
		Name:  "gato",
		Usage: "cat utility but in go and simpler",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "numbered",
				Aliases: []string{"n"},
				Usage:   "Include numbered lines",
				Value:   false,
			},
		},
		Action: func(ctx *cli.Context) error {
			red := color.New(color.FgRed).Println
			args := parseArgs(ctx)
			for _, arg := range args {
				if arg == "" {
					continue
				}
				f, err := os.Open(arg)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				fmt.Println()
				red(arg)
				if ctx.Bool("numbered") {
					err = readWriteNumbered(f)
				} else {
					err = readWriteBuffered(f)
				}
				if err != nil {
					return nil
				}
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
