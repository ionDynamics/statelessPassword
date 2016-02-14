package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"go.iondynamics.net/statelessPassword"
)

var (
	nameEnv = "ID_SLPW_FULLNAME"
	noEnv   = flag.Bool("noEnv", false, "Set this to true if no environment variables should be used")
)

func main() {
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	fullname := getFullname(r)
	bytMasterPw := getMasterpassword(r)

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
		site, ok := getSite(r)
		if !ok {
			continue
		}

		version, ok := getVersion(r)
		if !ok {
			continue
		}

		tlps, ok := getTemplates(r)
		if !ok {
			continue
		}

		params <- param{site, version, tlps}
		returnPw(pwd)
	}
}

type param struct {
	site    string
	version string
	tlps    []string
}
