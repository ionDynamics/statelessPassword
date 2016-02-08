package statelessPassword

import (
	"fmt"
)

func Example_Variant0() {
	algo1 := New([]byte("Example Full Name"), []byte("example password"), 0)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), BillemontLongPasswordTemplates))

	algo2 := New([]byte("Example Full Name"), []byte("example password"), 0)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), BillemontLongPasswordTemplates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)

	// Output:
	// Modi6/MabqRucm <nil>
	//
	// true
}

func Example_Variant1() {

	algo1 := New([]byte("Example Full Name"), []byte("example password"), 1)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), BillemontMaximumSecurityPasswordTemplates))

	algo2 := New([]byte("Example Full Name"), []byte("example password"), 1)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), BillemontMaximumSecurityPasswordTemplates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), BillemontMaximumSecurityPasswordTemplates)) == fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), BillemontMaximumSecurityPasswordTemplates)))

	// Output:
	// JHi0nK8vPzNObTJX9k9~ <nil>
	//
	// true
	// true
}

func Example_Variant2() {

	algo1 := New([]byte("Example Full Name"), []byte("example password"), 2)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Alphanumeric16Templates))

	algo2 := New([]byte("Example Full Name"), []byte("example password_"), 2)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Alphanumeric16Templates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Alphanumeric16Templates)) == fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Alphanumeric16Templates)))

	// Output:
	// xX3Zo1QFf2XZRlA2 <nil>
	//
	// false
	// false
}

func Example_Variant3() {

	algo1 := New([]byte("Example Full Name"), []byte("example password"), 3)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Printable25Templates))

	algo2 := New([]byte("Example Full Name"), []byte("example password"), 3)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Printable25Templates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Printable25Templates)) == fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Printable25Templates)))

	// Output:
	// x7&faGvyyY*JKkKiwekjJvY3X <nil>
	//
	// true
	// true
}

func Example_Variant5() {

	algo1 := New([]byte("Example Full Name"), []byte("example password"), 5)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Printable25Templates))

	algo2 := New([]byte("Example Full Name"), []byte("example password"), 5)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Printable25Templates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Printable25Templates)) == fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Printable25Templates)))

	// Output:
	// kvAEXYb&TWELL0vsJSdkaKd5A <nil>
	//
	// true
	// true
}

func Example_Variant0_32() {
	algo := New([]byte("Example Full Name"), []byte("example password"), uint8(0))
	fmt.Println(algo.Password("example.org", "1", Printable32Templates))

	// Output:
	// bKIGeX0gQXW*6v%dO3^SFLLKFBR0bV0A <nil>
}
