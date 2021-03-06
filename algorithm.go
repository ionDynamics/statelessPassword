// Package statelessPassword implements a derived Billemont's algorithm
// https://en.wikipedia.org/wiki/Master_Password
package statelessPassword //import "go.iondynamics.net/statelessPassword"

import (
	"bytes"
	"crypto/hmac"
	cRand "crypto/rand"
	"crypto/sha512"
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"golang.org/x/crypto/scrypt"
)

//ScryptParameter holds all parameters for the scrypt key derivation algorithm
type ScryptParameter struct {
	N      int
	R      int
	P      int
	KeyLen int
}

//Generator interface
type Generator interface {
	Init(fullname, masterpassword []byte) error
	Password(site string, version string, templates []string) (string, error)
}

type algorithm struct {
	init           bool
	fullname       []byte
	masterpassword []byte
	p              ScryptParameter
	saltPrefix     []byte
	key            []byte
}

//New calls NewCustom with preset-determined parameters
func New(fullname, masterpassword []byte, preset uint8) (Generator, error) {
	params := ScryptParameter{
		R:      8,
		KeyLen: 8096,
		P:      4,
	}

	switch preset {
	case 0:
		params.N = 32768
		params.P = 1
	case 1:
		params.N = 32768
		params.P = 2
	case 2:
		params.N = 32768
	case 3:
		params.N = 65536
	case 4:
		params.N = 131072
	case 5:
		params.N = 131072
		params.R = 12
	case 6:
		params.N = 1048576
		params.P = 1
	case 7:
		params.N = 1048576
	}

	return NewCustom(fullname, masterpassword, []byte("go.iondynamics.net/statelessPassword"), params)
}

//NewCustom creates and initializes a new Generator
func NewCustom(fullname, masterpassword, saltPrefix []byte, p ScryptParameter) (Generator, error) {
	algo := NewUninitalized(fullname, saltPrefix, p)
	return algo, algo.Init(fullname, masterpassword)
}

//NewUninitalized returns an algorithm without calling it's Init function
func NewUninitalized(fullname, saltPrefix []byte, p ScryptParameter) *algorithm {
	return &algorithm{
		p:          p,
		saltPrefix: saltPrefix,
		fullname:   fullname,
	}
}

//InitWithPrecalculatedKey initializes the algorithm with the provided input
func (algo *algorithm) InitWithPrecalculatedKey(fullname, key []byte) {
	algo.fullname = fullname
	algo.key = key
	algo.init = true
}

//Init initializes the algorithm with the provided input and precalculates the masterkey
func (algo *algorithm) Init(fullname, masterpassword []byte) error {
	algo.fullname = fullname
	err := algo.CalculateKey(masterpassword)
	if err != nil {
		return err
	}

	algo.init = true

	return err
}

//CalculateKey calculates the scrypt key
func (algo *algorithm) CalculateKey(masterpassword []byte) error {
	var err error
	algo.key, err = scrypt.Key(masterpassword,
		algo.ScryptSalt(),
		algo.p.N,
		algo.p.R,
		algo.p.P,
		algo.p.KeyLen)

	return err
}

//ScryptSalt returns the salt used by Scrypt
func (algo *algorithm) ScryptSalt() []byte {
	b := bytes.Buffer{}
	b.Write(algo.saltPrefix)
	b.WriteString(fmt.Sprint(len(algo.fullname)))
	b.Write(algo.fullname)

	return b.Bytes()
}

//Password returns the password for the given website, password version and templates
func (algo *algorithm) Password(site string, version string, templates []string) (string, error) {
	if !algo.init {
		return "", fmt.Errorf("%s", "not initialized")
	}

	if len(templates) < 1 {
		return "", fmt.Errorf("%s", "invalid template")
	}

	h := hmac.New(sha512.New, algo.key)

	_, err := h.Write([]byte(fmt.Sprint(string(algo.saltPrefix), len(site), site, version)))
	if err != nil {
		return "", err
	}

	seed, err := cRand.Int(bytes.NewReader(h.Sum(nil)), big.NewInt(math.MaxInt64))
	if err != nil {
		return "", err
	}

	prng := rand.New(rand.NewSource(seed.Int64()))
	template := templates[prng.Intn(len(templates))]
	password := []byte(template)

	for i, tplRune := range template {
		passchars, ok := tplChars[tplRune]
		if !ok {
			return "", fmt.Errorf("%s", "invalid template")
		}

		password[i] = passchars[prng.Intn(len(passchars))]
	}

	return string(password), err
}
