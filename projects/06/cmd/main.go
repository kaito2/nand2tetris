package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaito2/nand2tetris/internal"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gasm",
		Usage: "assembler generate .hack file",
		Action: func(c *cli.Context) error {
			filename := c.Args().First()
			if filename == "" {
				cli.ShowAppHelp(c)
				os.Exit(1)
			}
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				cli.ShowAppHelp(c)
				os.Exit(1)
			}
			base := filepath.Base(filename)
			err := internal.Parse(filename, binaryFilename(base))
			if err != nil {
				cli.ShowAppHelp(c)
				os.Exit(1)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// e.g. "Add.asm" => "Add.hack"
func binaryFilename(asmFilename string) string {
	extension := filepath.Ext(asmFilename)
	baseName := strings.TrimRight(asmFilename, extension)
	return fmt.Sprintf("%s.hack", baseName)
}
