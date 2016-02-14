package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"go.iondynamics.net/statelessPassword"
)

var (
	nameEnv    = "ID_SLPW_FULLNAME"
	noEnv      = flag.Bool("no-env", false, "Set this to true if no environment variables should be used")
	noTerminal = flag.Bool("no-terminal", false, "Set this to true if your standard input isn't a terminal")
)

func main() {
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	fullname := getFullname(s)
	bytMasterPw := getMasterpassword(s)

	params := make(chan param, 1)
	pwd := make(chan string, 1)

	go func() {
		algo, err := statelessPassword.New([]byte(fullname), bytMasterPw, 5)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for {
			p := <-params
			pw, err := algo.Password(p.site, p.version, p.tlps)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			pwd <- pw
		}
	}()

	for {
		site, ok := getSite(s)
		if !ok {
			continue
		}

		version, ok := getVersion(s)
		if !ok {
			continue
		}

		tlps, ok := getTemplates(s)
		if !ok {
			continue
		}

		params <- param{site, version, tlps}
		returnPw(pwd)

		if *noTerminal {
			break
		}
	}
}

type param struct {
	site    string
	version string
	tlps    []string
}
