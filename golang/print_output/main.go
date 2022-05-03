package main

import (
	"errors"
	"fmt"
)

func main() {

	var errorTest = errors.New("Error test") // error example
	fmt.Printf("%v\n", errorTest)

	var number = 100 // int
	fmt.Printf("%d\n", number)

	var number2 = 100.12 // float 64 example
	fmt.Printf("%v\n", number2)

	var string1 = "This is the example string" // string example
	fmt.Printf("%s\n", string1)

	var binary1 = 4
	var binary2 = 5
	fmt.Printf("Binary: %b\\%b\n", binary1, binary2) // binary example '%b' symbol refer to annotation verb formats for binary, '\\' => \ a special value for backslash and '\n' => new line

	var percentage = 50
	fmt.Printf("Percentage: %d %%\n", percentage) // Percentage example '%d' decimal and the verb %%, which consumes no argument, produces a percent sign

	//var percentage2 = 50
	//fmt.Printf("Percentage 2: %d %%", percentage2, binary1) // An extra argument is added to the statement.
	// Result --> Percentage: 50 %Percentage 2: 50 %%!(EXTRA int=4) It will compile and then catch this type of errors early
	//
	//
	//
	//
	//
	//
	// finding suspicious codes
	// you can use the vet command – it can find calls whose arguments do not align with the format string.
	// If you try to compile and run this incorrect line of code

	list := []int64{0, 1}
	fmt.Printf("list: %v\n", list)      // Default format 'list: [0 1]'
	fmt.Printf("list #: %#v\n", list)   // Go-syntax format 'list #: []int64{0, 1}'
	fmt.Printf("list type: %T\n", list) // The type of the value 'list type: []int64'

	// Integer
	intValue := 94
	fmt.Printf("base 10: %d\n", intValue)                     // Base 10 'base: 15'
	fmt.Printf("add sign added: %+d\n", intValue)             // Always show sign 'add sign added: +15'
	fmt.Printf("Space width right = 4: %4d\n", intValue)      // Pad with spaces (width 4 right) 'Space width right = 4:   15'
	fmt.Printf("Space width left = 4: %-4d\n", intValue)      // Pad with spaces (width 4 left) 'Space width left = 4: 15    '
	fmt.Printf("Pad with zeroes (width 4): %04d\n", intValue) // Pad with zeroes (width 4) 'Pad with zeroes (width 4): 0015'
	fmt.Printf("Base 2: %b\n", intValue)                      // Base 2 'Base 2: 10001' base 2 == normal binary '%b'

	// reference (octal number list): https://www.electrostudy.com/2015/07/octal-number-system-1-100.html
	octalExample := 8
	fmt.Printf("Base 8: %o\n", octalExample) // Base 8 'Base 8: 21' octal with '%o'

	base16Value := 13
	// This is the base 16 hex values, https://www.electronics-tutorials.ws/binary/bin_3.html
	// https://www.electrostudy.com/2015/08/hexadecimal-number-system-1-100.html
	fmt.Printf("Base 16 (lowercase): %x\n", base16Value)       // Base 16  hex(lowercase) 'Base 16 (lowercase): d'
	fmt.Printf("Base 16 (uppercase): %X\n", base16Value)       // Base 16  hex (uppercase) 'Base 16 (uppercase): D'
	fmt.Printf("Base 16, with leading 0x: %#x\n", base16Value) // Base 16 hex (uppercase) 'Base 16 (uppercase): 0xd'

	// character (quoted, Unicode)
	unicodeLetter := 33
	fmt.Printf("Character: %c\n", unicodeLetter)               // Base 16  hex(lowercase) 'Base 16 (lowercase): d'
	fmt.Printf("Quoted character: %q\n", unicodeLetter)        // Base 16  hex(lowercase) 'Base 16 (lowercase): d'
	fmt.Printf("Unicode: %U\n", unicodeLetter)                 // Base 16  hex(lowercase) 'Base 16 (lowercase): d'
	fmt.Printf("Unicode with character: %#U\n", unicodeLetter) // Base 16  hex(lowercase) 'Base 16 (lowercase): d'

	// euro sign
	// '\u20AC' -> euro sign
	//codeTest := '\u20AC'
	codeTest := '\u0021'
	fmt.Printf("Code test: %q\n", codeTest)

	valueOnly := fmt.Sprintf("%q", codeTest) // Return the resulting string.
	fmt.Println(valueOnly)

	/*
	     * Windows-1252 or CP-1252 (code page 1252) is a single-byte character encoding of the Latin alphabet,
	     * used by default in the legacy components of Microsoft Windows for English and many European languages including Spanish, French, and German.

	       0x80: '\u20AC', // EURO SIGN
		   0x81: '\uFFFD', // UNDEFINED
		   0x82: '\u201A', // SINGLE LOW-9 QUOTATION MARK
		   0x83: '\u0192', // LATIN SMALL LETTER F WITH HOOK
		   0x84: '\u201E', // DOUBLE LOW-9 QUOTATION MARK
		   0x85: '\u2026', // HORIZONTAL ELLIPSIS
		   0x86: '\u2020', // DAGGER
		   0x87: '\u2021', // DOUBLE DAGGER
		   0x88: '\u02C6', // MODIFIER LETTER CIRCUMFLEX ACCENT
		   0x89: '\u2030', // PER MILLE SIGN
		   0x8A: '\u0160', // LATIN CAPITAL LETTER S WITH CARON
		   0x8B: '\u2039', // SINGLE LEFT-POINTING ANGLE QUOTATION MARK
		   0x8C: '\u0152', // LATIN CAPITAL LIGATURE OE
		   0x8D: '\uFFFD', // UNDEFINED
		   0x8E: '\u017D', // LATIN CAPITAL LETTER Z WITH CARON
		   0x8F: '\uFFFD', // UNDEFINED
		   0x90: '\uFFFD', // UNDEFINED
		   0x91: '\u2018', // LEFT SINGLE QUOTATION MARK
		   0x92: '\u2019', // RIGHT SINGLE QUOTATION MARK
		   0x93: '\u201C', // LEFT DOUBLE QUOTATION MARK
		   0x94: '\u201D', // RIGHT DOUBLE QUOTATION MARK
		   0x95: '\u2022', // BULLET
		   0x96: '\u2013', // EN DASH
		   0x97: '\u2014', // EM DASH
		   0x98: '\u02DC', // SMALL TILDE
		   0x99: '\u2122', // TRADE MARK SIGN
		   0x9A: '\u0161', // LATIN SMALL LETTER S WITH CARON
		   0x9B: '\u203A', // SINGLE RIGHT-POINTING ANGLE QUOTATION MARK
		   0x9C: '\u0153', // LATIN SMALL LIGATURE OE
		   0x9D: '\uFFFD', // UNDEFINED
		   0x9E: '\u017E', // LATIN SMALL LETTER Z WITH CARON
		   0x9F: '\u0178', // LATIN CAPITAL LETTER Y WITH DIAERESIS
	*/

	// Boolean example
	// Use %t to format a boolean as true or false.
	booleanExample := true
	fmt.Printf("Boolean test: %t\n", booleanExample)

	// Pointer hex example
	// Use %p to format a pointer in base 16 notation with leading 0x.
	p := pointerStruct{1, 2}
	fmt.Printf("pointer format example:  %p\n", &p)

	// Float example
	// Float (indent, precision, scientific notation)
	floatValueExample := 123.456
	fmt.Printf("Scientific notation example:  %e\n", floatValueExample)
	fmt.Printf("Decimal point, no exponent example:  %f\n", floatValueExample)
	fmt.Printf("Default width, precision 2 example:  %.2f\n", floatValueExample)
	fmt.Printf("Width 8, precision 2 example:  %8.2f\n", floatValueExample)
	fmt.Printf("Exponent as needed, necessary digits only example:  %g\n", floatValueExample)

	// String or byte slice (quote, indent, hex)
	stringSliceTestValue := "value"
	fmt.Printf("Plain string example:  %s\n", stringSliceTestValue)
	fmt.Printf("Width of 6, with spacing from the left example: %6s\n", stringSliceTestValue)
	fmt.Printf("Width of 6, with spacing from the right example: %-6s\n", stringSliceTestValue)
	fmt.Printf("Quotes in string example: %q\n", stringSliceTestValue)
	fmt.Printf("Hex dump of byte values example: %x\n", stringSliceTestValue)
	fmt.Printf("Hex dump of byte values with spacing example: % x\n", stringSliceTestValue)

	// Special values
	fmt.Println("Alert or bell example: \u0008 test\u000A")
	fmt.Println("Form feed example: \u000C test\n")
	fmt.Println("carriage return example: \u000D test\n")
	fmt.Println("vertical tab example: \u000b test\n")

}

type pointerStruct struct {
	x, y int
}
