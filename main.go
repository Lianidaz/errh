package errh

import (
	"fmt"
	"os"
)

// ErrH - hanles errors for you. Pass error and level ( l(log), w(warning), f(fatal), p(panic), x(exit clean))
func ErrH(err error, lvl string) {
	if err == nil {
		return
	}
	switch lvl {
	case "l":
		fmt.Printf("LOG: %+v\n", err)
		return
	case "w":
		fmt.Printf("WARNING: %+v\n", err)
		return
	case "f":
		defer os.Exit(1)
		fmt.Printf("FATAL: %+v\n", err)
		return
	case "p":
		defer panic(fmt.Sprintf("%+v\n", err))
		return
	case "x":
		defer os.Exit(0)
		fmt.Printf("ERROR: %+v\nWill now exit\n", err)
		return
	}
}
