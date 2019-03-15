// Package errh - set o functions to avoud writing if err != nil...
package errh

import (
	"fmt"
	"os"
)

// Err - hanles errors for you. Pass error and level ( l(log), w(warning), f(fatal), p(panic), x(exit clean))
func Err(err error, lvl string) {
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

// Errl - simply log err
func Errl(err error) {
	fmt.Printf("LOG: %+v\n", err)
	return
}

// Errw - log err as a warning
func Errw(err error) {
	fmt.Printf("WARNING: %+v\n", err)
	return
}

// Errf - log and throw, exit code 1
func Errf(err error) {
	defer os.Exit(1)
	fmt.Printf("FATAL: %+v\n", err)
	return
}

// Errp - panic with err
func Errp(err error) {
	defer panic(fmt.Sprintf("%+v\n", err))
	return
}

// Errx - log err and clean exit
func Errx(err error) {
	defer os.Exit(0)
	fmt.Printf("ERROR: %+v\nWill now exit\n", err)
	return
}
