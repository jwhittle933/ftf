package cli

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jwhittle933/funked/term/colors"
)

type Flags struct {
	JSON   string  `short:"J" long:"json" description:"pass input as json" json:"-"`
	Height string  `short:"H" long:"height" description:"User's height" json:"height"`
	Weight float64 `short:"W" long:"weight" description:"User's weight" json:"weight"`
	Age    int     `short:"A" long:"age" description:"User's age" json:"age"`
	Sex    string  `short:"S" long:"sex" description:"User's sex/gender" json:"sex"`
	Level  float64 `short:"L" long:"level" description:"level of activity between 1.2 and 1.9" json:"level"`
	Help   bool    `short:"h" long:"help" description:"show help message" json:"-"`
}

func (f *Flags) HelpText() string {
	red := colors.NewRGB(170, 0, 120)
	green := colors.NewRGB(0, 140, 140)

	out := strings.Builder{}
	out.WriteString(
		fmt.Sprintf(
			"%s BMR Calculator © Jonathann Whittle, 2021\n\n",
			red.Sprintf("Fat To Fit"),
		),
	)

	t := reflect.TypeOf(*f)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.IsExported() {
			short, long, desc := tags(field.Tag)
			out.WriteString(
				fmt.Sprintf(
					"  [-%s | --%s]  %s\n",
					green.Sprintf(short),
					green.Sprintf(long),
					desc,
				),
			)
		}
	}

	out.WriteString("\n")
	return out.String()
}

func tags(t reflect.StructTag) (string, string, string) {
	return t.Get("short"), t.Get("long"), t.Get("description")
}
