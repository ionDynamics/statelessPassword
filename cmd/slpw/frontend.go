package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"golang.org/x/crypto/ssh/terminal"

	"go.iondynamics.net/statelessPassword"
)

func getFullname(s *bufio.Scanner) string {
	var fullname string
	if !*noEnv {
		fullname = os.Getenv(nameEnv)
	}

	if fullname == "" {
		say("Enter Full Name:")

		var err error
		fullname, err = getString(s)
		if err != nil {
			say(err)
			os.Exit(1)
		}

		if !*noEnv {
			err = os.Setenv(nameEnv, fullname)
			if err != nil {
				say(err)
				os.Exit(1)
			}
		}
	} else {
		say("Generating passwords for", fullname)
	}

	return strings.TrimSpace(fullname)
}

func getMasterpassword(s *bufio.Scanner) []byte {
	say("Enter Master Password:")
	var err error
	var mpw []byte

	if *noTerminal {
		var str string
		str, err = getString(s)
		mpw = []byte(str)
	} else {
		mpw, err = terminal.ReadPassword(int(os.Stdin.Fd()))
	}

	if err != nil {
		say(err)
		os.Exit(1)
	}
	return mpw
}

func getSite(s *bufio.Scanner) (string, bool) {
	say("Enter Site:")
	site, err := getString(s)
	if err != nil {
		say(err)
		return "", false
	}

	return strings.TrimSpace(site), true
}

func getVersion(s *bufio.Scanner) (string, bool) {
	say("Enter Password Version (Default 1) :")
	version, err := getString(s)
	if err != nil {
		say(err)
		return "", false
	}
	version = strings.TrimSpace(version)
	if version == "" {
		version = "1"
	}

	return version, true
}

func getTemplates(s *bufio.Scanner) ([]string, bool) {
	say("Choose Template:")
	say("1 default) 16 alphanumeric")
	say("2) 20 alphanumeric")
	say("3) 10 alphanumeric")
	say("4) 8 alphanumeric")
	say("5) 16 with special chars")
	say("6) 25 with special chars")
	say("7) 32 with special chars")
	say("8) 4 digit PIN")

	tpls := statelessPassword.Alphanumeric16Templates

	tplStr, err := getString(s)
	if err != nil {
		return tpls, false
	}

	switch strings.TrimSpace(tplStr) {
	case "2":
		tpls = statelessPassword.Alphanumeric20Templates
	case "3":
		tpls = statelessPassword.Alphanumeric10Templates
	case "4":
		tpls = statelessPassword.Alphanumeric8Templates
	case "5":
		tpls = statelessPassword.Printable16Templates
	case "6":
		tpls = statelessPassword.Printable25Templates
	case "7":
		tpls = statelessPassword.Printable32Templates
	case "8":
		tpls = statelessPassword.BillemontPINTemplates
	}

	return tpls, true
}

func returnPw(pwch chan string) {
	var pwd string
	select {
	case pw := <-pwch:
		pwd = pw
	case <-time.After(750 * time.Millisecond):
		say("Generating...")
		pwd = <-pwch
	}

	if *noTerminal {
		fmt.Print(pwd)
		return
	}

	before, err := clipboard.ReadAll()
	clipboard.WriteAll(pwd)
	say("\nPassword copied to clipboard! ")
	time.Sleep(5 * time.Second)
	say("Cleaning clipboard in 5 seconds...")
	time.Sleep(5 * time.Second)
	if err != nil {
		clipboard.WriteAll("")
	} else {
		clipboard.WriteAll(before)
	}
	say("\n...again? or CTRL+C\n")
}

func getString(s *bufio.Scanner) (string, error) {
	for s.Scan() {
		return s.Text(), nil
		break
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	return "", nil
}

func say(s ...interface{}) {
	if *noTerminal {
		return
	}

	fmt.Println(s...)
}
