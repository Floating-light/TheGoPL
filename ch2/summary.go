/* Chapter 2 Program Structure
   In this chapter we will go into more detail about the basic structural elements of a Go program.
   2.1 Names
	    1. 25 keywords cannot be used as names.
	    2. three dozen predelared names:
		   Constants: true false iota nil
		   Types: int int8 uint uintptr complex64 byte rune error string, etc...
		   Functions: make len new delete real imag panic recover, etc...
		   these names are not reserved, so you may use them in declarations.
		3. package names always in lower case.
		4. the name length in Go lean toward short names,especially for local variables with small scopes;
		   the larger the scope of a name, the longer and more meaningful it should be.
		5. use "camel case" for variable's name.
		   the letter of acronyms always in the same case, htmlEscape,HTMLEscape,or excapeHTML,but not escapeHtml.
*/
package ch2

