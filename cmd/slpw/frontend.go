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

func getFullname(r *bufio.Reader) string {
	var fullname string
	if !*noEnv {
		fullname = os.Getenv(nameEnv)
	}

	if fullname == "" {
		fmt.Println("Enter Full Name:")

		var err error
		fullname, err = r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if !*noEnv {
			err = os.Setenv(nameEnv, fullname)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("Generating passwords for", fullname)
	}

	return strings.TrimSpace(fullname)
}

func getMasterpassword(r *bufio.Reader) []byte {
	fmt.Println("Enter Master Password:")
	bytMasterPw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bytMasterPw
}

func getSite(r *bufio.Reader) (string, bool) {
	fmt.Println("Enter Site:")
	site, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	return strings.TrimSpace(site), true
}

func getVersion(r *bufio.Reader) (string, bool) {
	fmt.Println("Enter Password Version (Default 1) :")
	version, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	version = strings.TrimSpace(version)
	if version == "" {
		version = "1"
	}

	return version, true
}

func getTemplates(r *bufio.Reader) ([]string, bool) {
	fmt.Println("Choose Template:")
	fmt.Println("1 default) 16 alphanumeric")
	fmt.Println("2) 20 alphanumeric")
	fmt.Println("3) 10 alphanumeric")
	fmt.Println("4) 8 alphanumeric")
	fmt.Println("5) 16 with special chars")
	fmt.Println("6) 25 with special chars")
	fmt.Println("7) 32 with special chars")
	fmt.Println("8) 4 digit PIN")

	tpls := statelessPassword.Alphanumeric16Templates

	tplStr, err := r.ReadString('\n')
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
		fmt.Println("Generating...")
		pwd = <-pwch
	}

	before, err := clipboard.ReadAll()
	clipboard.WriteAll(pwd)
	fmt.Println("\nPassword copied to clipboard! ")
	time.Sleep(5 * time.Second)
	fmt.Println("Cleaning clipboard in 5 seconds...")
	time.Sleep(5 * time.Second)
	if err != nil {
		clipboard.WriteAll("")
	} else {
		clipboard.WriteAll(before)
	}
	fmt.Println("\n...again? or CTRL+C\n")
}
