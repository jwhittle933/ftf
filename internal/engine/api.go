package engine

import (
	"encoding/json"
	"github.com/jwhittle933/lext"
	"os"

	"github.com/jwhittle933/fattofit/internal/cli"
	"github.com/jwhittle933/fattofit/internal/empty"
)

type Validator func(flags *cli.Flags)

type Engine struct {
	Flags *cli.Flags
	Data  *Data
	l     lext.Logger
}

func Validate(args *cli.Flags, vds ...Validator) {
	for _, v := range vds {
		v(args)
	}
}

func ExitOnHelp(l lext.Logger) func(f *cli.Flags) {
	return func(f *cli.Flags) {
		if f.Help {
			l.Info(f.HelpText())
			os.Exit(0)
		}
	}
}

func ParseIfJSON(l lext.Logger) func(f *cli.Flags) {
	return func(f *cli.Flags) {
		if !empty.String(f.JSON) {
			l.Info("Parsing JSON")

			if err := json.Unmarshal([]byte(f.JSON), f); err != nil {
				l.Kill("Error parsing json: ", err.Error())
			}
		}

	}
}

func ExitIfMissingInfo(l lext.Logger) func(*cli.Flags) {
	return func(f *cli.Flags) {
		if empty.Int(f.Age) {
			l.Kill("Missing Age")
		}

		if empty.String(f.Height) {
			l.Kill("Missing Height")
		}

		if empty.Float(f.Weight) {
			l.Kill("Missing Weight")
		}

		if empty.String(f.Sex) {
			l.Kill("Missing info. Sex is required.")
		}
	}
}