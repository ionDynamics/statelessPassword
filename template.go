package statelessPassword //import "go.iondynamics.net/statelessPassword"

var (
	tplChars = map[rune]string{
		'V': "AEIOU",
		'C': "BCDFGHJKLMNPQRSTVWXYZ",
		'v': "aeiou",
		'c': "bcdfghjklmnpqrstvwxyz",
		'A': "AEIOUBCDFGHJKLMNPQRSTVWXYZ",
		'a': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz",
		'n': "0123456789",
		'o': "@&%?,=[]_:-+*$#!'^~;()/.",
		'X': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789@&%?,=[]_:-+*$#!'^~;()/.",
		'x': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789@&%?,=[]_:-+*$#!'^~;()/.",
		'p': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789",
	}

	BillemontMaximumSecurityPasswordTemplates = []string{"anoxxxxxxxxxxxxxxxxx", "axxxxxxxxxxxxxxxxxno"}
	BillemontLongPasswordTemplates            = []string{"CvcvnoCvcvCvcv", "CvcvCvcvnoCvcv", "CvcvCvcvCvcvno", "CvccnoCvcvCvcv", "CvccCvcvnoCvcv", "CvccCvcvCvcvno", "CvcvnoCvccCvcv", "CvcvCvccnoCvcv", "CvcvCvccCvcvno", "CvcvnoCvcvCvcc", "CvcvCvcvnoCvcc", "CvcvCvcvCvccno", "CvccnoCvccCvcv", "CvccCvccnoCvcv", "CvccCvccCvcvno", "CvcvnoCvccCvcc", "CvcvCvccnoCvcc", "CvcvCvccCvccno", "CvccnoCvcvCvcc", "CvccCvcvnoCvcc", "CvccCvcvCvccno"}
	BillemontMediumPasswordTemplates          = []string{"CvcnoCvc", "CvcCvcno"}
	BillemontBasicPasswordTemplates           = []string{"aaanaaan", "aannaaan", "aaannaaa"}
	BillemontShortPasswordTemplates           = []string{"Cvcn"}
	BillemontPINTemplates                     = []string{"nnnn"}

	Alphanumeric8Templates  = []string{"pppppppp"}
	Alphanumeric10Templates = []string{"pppppppppp"}
	Alphanumeric16Templates = []string{"pppppppppppppppp"}
	Alphanumeric20Templates = []string{"pppppppppppppppppppp"}
	Printable16Templates    = []string{"xxxxxxxxxxxxxxxx"}
	Printable25Templates    = []string{"xxxxxxxxxxxxxxxxxxxxxxxxx"}
	Printable32Templates    = []string{"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
)
