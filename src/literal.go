package main

import "fmt"

func main() {
	// Integer literals
	var i uint64
	i = 42
	i = 4_2
	i = 0600
	i = 0_600
	i = 0o600
	i = 0O600 // second character is capital letter 'O'
	i = 0xBadFace
	i = 0xBad_Face
	i = 0x_67_7a_2f_cc_40_c6
	i = 17014118346046923173
	i = 170_141183_460469_23173
	fmt.Println(i)

	// Floating-point literals
	var f float64
	f = 0.
	f = 72.40
	f = 072.40 // == 72.40
	f = 2.71828
	f = 1.e+0
	f = 6.67428e-11
	f = 1e6
	f = .25
	f = .12345e+5
	f = 1_5.        // == 15.0
	f = 0.15e+0_2   // == 15.0
	f = 0x1p-2      // == 0.25
	f = 0x2.p10     // == 2048.0
	f = 0x1.Fp+0    // == 1.9375
	f = 0x.8p-0     // == 0.5
	f = 0x15e - 2   // == 0x15e - 2 (integer subtraction)
	f = 0x_1FFFp-16 // == 0.1249847412109375
	fmt.Println(f)

	// Imaginary literals
	var c complex128
	c = 0i
	c = 123i   // == 123i for backward-compatibility
	c = 0o123i // == 0o123 * 1i == 83i
	c = 0xabci // == 0xabc * 1i == 2748i
	c = 0.i
	c = 2.71828i
	c = 1.e+0i
	c = 6.67428e-11i
	c = 1e6i
	c = .25i
	c = .12345e+5i
	c = 0x1p-2i // == 0x1p-2 * 1i == 0.25i
	fmt.Println(c)

	// Rune literals
	var r rune
	r = 'a'
	r = '本'
	r = '\'' // rune literal containing single quote character
	r = '\t'
	r = '\000'
	r = '\007'
	r = '\377'
	r = '\x07'
	r = '\xff'
	r = '\u12e4'
	r = '\U00101234'
	r = 'ä'
	fmt.Println(r)

	// String literals
	var s string
	s = `abc` // same as "abc"
	s = `\n
\n` // same as "\\n\n\\n"
	s = "\n"
	s = "\"" // same as `"`
	s = "Hello, world!\n"
	s = "日本語"
	s = "\u65e5本\U00008a9e"
	s = "\xff\u00FF"
	s = "日本語"                               // UTF-8 input text
	s = `日本語`                               // UTF-8 input text as raw literal
	s = "\u65e5\u672c\u8a9e"                   // the explicit Unicode code points
	s = "\U000065e5\U0000672c\U00008a9e"       // the explicit Unicode code points
	s = "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e" // the explicit UTF-8 bytes
	fmt.Println(s)
}
