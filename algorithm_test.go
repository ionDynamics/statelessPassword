package statelessPassword

import (
	"fmt"
)

func Example_Preset0() {
	algo1, err1 := New([]byte("Example Full Name"), []byte("example password"), 0)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), BillemontLongPasswordTemplates))

	algo2, err2 := New([]byte("Example Full Name"), []byte("example password"), 0)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), BillemontLongPasswordTemplates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(err1, err2)

	// Output:
	// MaweHivr6?Boho <nil>
	//
	// true
	// <nil> <nil>
}

func Example_Preset1() {
	algo1, err1 := New([]byte("Example Full Name"), []byte("example password"), 1)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), BillemontMaximumSecurityPasswordTemplates))

	algo2, err2 := New([]byte("Example Full Name"), []byte("example password"), 1)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), BillemontMaximumSecurityPasswordTemplates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), BillemontMaximumSecurityPasswordTemplates)) == fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), BillemontMaximumSecurityPasswordTemplates)))
	fmt.Println(err1, err2)

	// Output:
	// XIA[LAcRhS=CbF$B;N0^ <nil>
	//
	// true
	// true
	// <nil> <nil>
}

func Example_Preset2() {
	algo1, err1 := New([]byte("Example Full Name"), []byte("example password"), 2)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Alphanumeric16Templates))

	algo2, err2 := New([]byte("Example Full Name"), []byte("example password_"), 2)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Alphanumeric16Templates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Alphanumeric16Templates)) == fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Alphanumeric16Templates)))
	fmt.Println(err1, err2)

	// Output:
	// Vfv4Ye3TtsMx74Q0 <nil>
	//
	// false
	// false
	// <nil> <nil>
}

func Example_Preset3() {
	algo1, err1 := New([]byte("Example Full Name"), []byte("example password"), 3)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Printable25Templates))

	algo2, err2 := New([]byte("Example Full Name"), []byte("example password"), 3)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Printable25Templates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Printable25Templates)) == fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Printable25Templates)))
	fmt.Println(err1, err2)

	// Output:
	// :FIobQ9nLmMqNyvnzM7x=?LZf <nil>
	//
	// true
	// true
	// <nil> <nil>
}

func Example_Preset5() {
	algo1, err1 := New([]byte("Example Full Name"), []byte("example password"), 5)
	out1 := fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Printable25Templates))

	algo2, err2 := New([]byte("Example Full Name"), []byte("example password"), 5)
	out2 := fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Printable25Templates))

	fmt.Println(out1)
	fmt.Println(out1 == out2)
	fmt.Println(fmt.Sprintln(algo1.Password("example.org", string([]byte{1}), Printable25Templates)) == fmt.Sprintln(algo2.Password("example.org", string([]byte{1}), Printable25Templates)))
	fmt.Println(err1, err2)

	// Output:
	// [kI0wE,V=z/u#g+ie'[^f/pW7 <nil>
	//
	// true
	// true
	// <nil> <nil>
}

func Example_Preset0_32() {
	algo, err := New([]byte("Example Full Name"), []byte("example password"), uint8(0))
	out1, _ := algo.Password("example.org", "1", Printable32Templates)
	out2, _ := algo.Password("example.com", "1", Printable32Templates)
	out3, _ := algo.Password("example.com", "2", Printable32Templates)
	fmt.Println(algo.Password("example.org", "1", Printable32Templates))
	fmt.Println(err)
	fmt.Println(out1 != out2)
	fmt.Println(out2 != out3)

	// Output:
	// G/q6Mu*_=''oI]+kZ@_*/wsNxJS8/:S[ <nil>
	// <nil>
	// true
	// true
}

func Example_NotInitialized() {
	algo := &algorithm{}
	fmt.Println(algo.Password("example.org", "1", Printable32Templates))

	// Output:
	// not initialized
}
