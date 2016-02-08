package statelessPassword //import "go.iondynamics.net/statelessPassword"

const V = "AEIOU"
const C = "BCDFGHJKLMNPQRSTVWXYZ"
const v = "aeiou"
const c = "bcdfghjklmnpqrstvwxyz"
const A = "AEIOUBCDFGHJKLMNPQRSTVWXYZ"
const a = "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"
const n = "0123456789"
const o = "@&%?,=[]_:-+*$#!'^~;()/."
const X = "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()"
const x = X

const p = "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789"

var BillemontMaximumSecurityPasswordTemplates = []string{"anoxxxxxxxxxxxxxxxxx", "axxxxxxxxxxxxxxxxxno"}
var BillemontLongPasswordTemplates = []string{"CvcvnoCvcvCvcv", "CvcvCvcvnoCvcv", "CvcvCvcvCvcvno", "CvccnoCvcvCvcv", "CvccCvcvnoCvcv", "CvccCvcvCvcvno", "CvcvnoCvccCvcv", "CvcvCvccnoCvcv", "CvcvCvccCvcvno", "CvcvnoCvcvCvcc", "CvcvCvcvnoCvcc", "CvcvCvcvCvccno", "CvccnoCvccCvcv", "CvccCvccnoCvcv", "CvccCvccCvcvno", "CvcvnoCvccCvcc", "CvcvCvccnoCvcc", "CvcvCvccCvccno", "CvccnoCvcvCvcc", "CvccCvcvnoCvcc", "CvccCvcvCvccno"}
var BillemontMediumPasswordTemplates = []string{"CvcnoCvc", "CvcCvcno"}
var BillemontBasicPasswordTemplates = []string{"aaanaaan", "aannaaan", "aaannaaa"}
var BillemontShortPasswordTemplates = []string{"Cvcn"}
var BillemontPINTemplates = []string{"nnnn"}

var Alphanumeric8Templates = []string{"pppppppp"}
var Alphanumeric10Templates = []string{"pppppppppp"}
var Alphanumeric16Templates = []string{"pppppppppppppppp"}
var Alphanumeric20Templates = []string{"pppppppppppppppppppp"}
var Printable16Templates = []string{"xxxxxxxxxxxxxxxx"}
var Printable25Templates = []string{"xxxxxxxxxxxxxxxxxxxxxxxxx"}
var Printable32Templates = []string{"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
