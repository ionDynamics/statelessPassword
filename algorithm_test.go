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
	// XopmTegq7&Sizi <nil>
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
	// j_Kf6#8:tC+#5u'C-c4$ <nil>
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
	// 8dh5hmmiH6oWlnWf <nil>
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
	// yJ$#PLnF&y.l$~SR.8&0EBNtG <nil>
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
	// t'$1sX'_OOAyj?qOpR8%q2ryF <nil>
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
	// s0tzQE;SLKbjms+JuQ'T.LV%a5r1[N^3 <nil>
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
