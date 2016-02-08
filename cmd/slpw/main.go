package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
	"golang.org/x/crypto/ssh/terminal"

	"go.iondynamics.net/statelessPassword"
)

var fullnameEnvName string = "ID_SLPW_FULLNAME"
var defaultVariantEnvName string = "ID_SLPW_DEFAULTVARIANT"

var noEnvPtr = flag.Bool("noEnv", false, "Set this to true if no environment variables should be used")

func main() {
	var defVar uint8 = 3
	defVarUsed := defVar
	defVarStr := os.Getenv(defaultVariantEnvName)
	if defVarStr != "" {
		i, err := strconv.Atoi(defVarStr)
		if err == nil {
			defVarUsed = uint8(i)
		}
	}

	var variantPtr = flag.Uint("variant", uint(defVarUsed), "Run slpw_perf to see the performance of each variant")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	var fullname string

	if !*noEnvPtr {
		fullname = os.Getenv(fullnameEnvName)
	} else {
		defVarUsed = defVar
	}

	if *variantPtr == 0 {
		fmt.Println("LEGACY MODE")
	}

	if fullname == "" {
		fmt.Println("Enter Full Name:")

		var err error
		fullname, err = r.ReadString('\n')
		if err != nil {
			fmt.Errorf("%s", err)
			os.Exit(1)
		}

		if !*noEnvPtr {
			err = os.Setenv(fullnameEnvName, fullname)
			if err != nil {
				fmt.Errorf("%s", err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("Found User:", fullname)
	}

	fmt.Println("Variant:", *variantPtr)

	fmt.Println("Enter Master Password:")
	bytMasterPw, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}

	fullname = strings.TrimSpace(fullname)

	algo := statelessPassword.New([]byte(fullname), bytMasterPw, uint8(*variantPtr))

	for {
		fmt.Println("Enter Site:")
		site, err := r.ReadString('\n')
		if err != nil {
			fmt.Errorf("%s", err)
			continue
		}

		fmt.Println("Enter Password Version (Default 1) :")
		version, _ := r.ReadString('\n')
		version = strings.TrimSpace(version)
		if version == "" {
			version = "1"
		}

		fmt.Println("Choose Template:")
		fmt.Println("1 default) 16 alphanumeric")
		fmt.Println("2) 20 alphanumeric")
		fmt.Println("3) 10 alphanumeric")
		fmt.Println("4) 8 alphanumeric")
		fmt.Println("5) 16 with special chars")
		fmt.Println("6) 25 with special chars")
		fmt.Println("7) 32 with special chars")
		fmt.Println("8) 4 digit PIN")
		fmt.Println("9) Billemont Long 14 with special chars")

		tpls := statelessPassword.Alphanumeric16Templates

		tplStr, _ := r.ReadString('\n')
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
		case "9":
			tpls = statelessPassword.BillemontLongPasswordTemplates
		}

		site = strings.TrimSpace(site)

		pwd, err := algo.Password(site, version, tpls)
		if err != nil {
			fmt.Errorf("%s", err)
			os.Exit(1)
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
		fmt.Println("\n\n...again? or CTRL+C")
	}
}
