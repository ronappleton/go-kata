package languages

// DefaultLanguage is used by Lex when no language is supplied.
var DefaultLanguage = &Language{
	ID: "go", Name: "Go", Extensions: []string{".go"},
	SourceFilename: "solution.go", TestsFilename: "learner_test.go",
	LineComment: "//", BlockCommentStart: "/*", BlockCommentEnd: "*/",
	QuoteChars: "`\"'",
	AutoPairs: []AutoPair{
		{Open: '(', Close: ')'},
		{Open: '[', Close: ']'},
		{Open: '{', Close: '}'},
		{Open: '"', Close: '"'},
	},
	Keywords: []string{
		"func", "return", "if", "else", "for", "range", "package", "import",
		"type", "struct", "interface", "var", "const", "go", "defer",
		"select", "switch", "case", "default", "break", "continue",
		"map", "chan", "nil", "true", "false", "fallthrough", "goto",
	},
	Types: []string{
		"string", "bool", "byte", "rune", "int", "int8", "int16", "int32",
		"int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32",
		"float64", "complex64", "complex128", "uintptr", "any", "error",
	},
	Indent: "\t",
}
