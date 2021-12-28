package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/jwhittle933/fattofit/internal/cli"
	"os"

	"github.com/jwhittle933/lext"

	"github.com/jwhittle933/fattofit/internal/engine"
)

func main() {
	l := lext.New(os.Stdout)

	args := &cli.Flags{}
	flags.NewParser(args, 0).ParseArgs(os.Args)

	engine.Validate(
		args,
		engine.ExitOnHelp(l),
		engine.ParseIfJSON(l),
		engine.ExitIfMissingInfo(l),
	)
}

