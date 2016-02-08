package statelessPassword

import (
	"fmt"
)

func Example_V() {
	fmt.Println(V)
	// Output:
	// AEIOU
}

func Example_C() {
	fmt.Println(C)
	// Output:
	// BCDFGHJKLMNPQRSTVWXYZ
}

func Example_c() {
	fmt.Println(c)
	// Output:
	// bcdfghjklmnpqrstvwxyz
}

func Example_v() {
	fmt.Println(v)
	// Output:
	// aeiou
}

func Example_A() {
	fmt.Println(A)
	// Output:
	// AEIOUBCDFGHJKLMNPQRSTVWXYZ
}

func Example_a() {
	fmt.Println(a)
	// Output:
	// AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz
}

func Example_n() {
	fmt.Println(n)
	// Output:
	// 0123456789
}

func Example_o() {
	fmt.Println(o)
	// Output:
	// @&%?,=[]_:-+*$#!'^~;()/.
}

func Example_X() {
	fmt.Println(X)
	// Output:
	// AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()
}

func Example_x() {
	fmt.Println(x)
	// Output:
	// AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()
}

func Example_p() {
	fmt.Println(p)
	// Output:
	// AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789
}

func Example_BillemontMaximumSecurityPasswordTemplates() {
	fmt.Println(BillemontMaximumSecurityPasswordTemplates)
	// Output:
	// [anoxxxxxxxxxxxxxxxxx axxxxxxxxxxxxxxxxxno]
}

func Example_BillemontLongPasswordTemplates() {
	fmt.Println(BillemontLongPasswordTemplates)
	// Output:
	// [CvcvnoCvcvCvcv CvcvCvcvnoCvcv CvcvCvcvCvcvno CvccnoCvcvCvcv CvccCvcvnoCvcv CvccCvcvCvcvno CvcvnoCvccCvcv CvcvCvccnoCvcv CvcvCvccCvcvno CvcvnoCvcvCvcc CvcvCvcvnoCvcc CvcvCvcvCvccno CvccnoCvccCvcv CvccCvccnoCvcv CvccCvccCvcvno CvcvnoCvccCvcc CvcvCvccnoCvcc CvcvCvccCvccno CvccnoCvcvCvcc CvccCvcvnoCvcc CvccCvcvCvccno]
}

func Example_BillemontMediumPasswordTemplates() {
	fmt.Println(BillemontMediumPasswordTemplates)
	// Output:
	// [CvcnoCvc CvcCvcno]
}

func Example_BillemontBasicPasswordTemplates() {
	fmt.Println(BillemontBasicPasswordTemplates)
	// Output:
	// [aaanaaan aannaaan aaannaaa]
}

func Example_BillemontShortPasswordTemplates() {
	fmt.Println(BillemontShortPasswordTemplates)
	// Output:
	// [Cvcn]
}

func Example_BillemontPINTemplates() {
	fmt.Println(BillemontPINTemplates)
	// Output:
	// [nnnn]
}

func Example_Alphanumeric8Templates() {
	fmt.Println(Alphanumeric8Templates)
	// Output:
	// [pppppppp]
}

func Example_Alphanumeric10Templates() {
	fmt.Println(Alphanumeric10Templates)
	// Output:
	// [pppppppppp]
}

func Example_Alphanumeric16Templates() {
	fmt.Println(Alphanumeric16Templates)
	// Output:
	// [pppppppppppppppp]
}

func Example_Alphanumeric20Templates() {
	fmt.Println(Alphanumeric20Templates)
	// Output:
	// [pppppppppppppppppppp]
}

func Example_Printable16Templates() {
	fmt.Println(Printable16Templates)
	// Output:
	// [xxxxxxxxxxxxxxxx]
}

func Example_Printable25Templates() {
	fmt.Println(Printable25Templates)
	// Output:
	// [xxxxxxxxxxxxxxxxxxxxxxxxx]
}

func Example_Printable32Templates() {
	// Output:
	fmt.Println(Printable32Templates)
	// [xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx]
}
