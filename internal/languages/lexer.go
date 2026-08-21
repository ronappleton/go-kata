package languages

// Span is one highlighted region within source code.
type Span struct {
	Tag   string // highlight tag name: keyword, string, comment, number, type, func
	Start int
	End   int
}

// HighlightTags are the tag names the editor binds.
const (
	TagKeyword = "hl-keyword"
	TagString  = "hl-string"
	TagComment = "hl-comment"
	TagNumber  = "hl-number"
	TagType    = "hl-type"
	TagFunc    = "hl-func"
)

// Lex tokenizes src into highlightable spans using the language's keyword,
// type, comment, and quote definitions.
func Lex(src string, lang *Language) []Span {
	if lang == nil {
		lang = DefaultLanguage
	}

	keywords := make(map[string]bool, len(lang.Keywords))
	for _, k := range lang.Keywords {
		keywords[k] = true
	}
	types := make(map[string]bool, len(lang.Types))
	for _, t := range lang.Types {
		types[t] = true
	}

	quoteSet := make(map[byte]bool, len(lang.QuoteChars))
	for i := 0; i < len(lang.QuoteChars); i++ {
		quoteSet[lang.QuoteChars[i]] = true
	}

	var spans []Span
	lineComment := lang.LineComment
	blockStart := lang.BlockCommentStart
	blockEnd := lang.BlockCommentEnd

	isIdent := func(c byte) bool {
		return c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
	}
	isDigit := func(c byte) bool { return c >= '0' && c <= '9' }

	for i := 0; i < len(src); {
		c := src[i]

		// Line comments (//, #, -- etc).
		if lineComment != "" && stringsHasPrefix(src[i:], lineComment) {
			j := i + len(lineComment)
			for j < len(src) && src[j] != '\n' {
				j++
			}
			spans = append(spans, Span{TagComment, i, j})
			i = j
			continue
		}
		// Block comments.
		if blockStart != "" && stringsHasPrefix(src[i:], blockStart) {
			j := i + len(blockStart)
			if blockEnd != "" {
				for j < len(src) {
					if stringsHasPrefix(src[j:], blockEnd) {
						j += len(blockEnd)
						break
					}
					j++
				}
			} else {
				j = len(src)
			}
			spans = append(spans, Span{TagComment, i, j})
			i = j
			continue
		}
		// Strings and rune literals.
		if quoteSet[c] {
			j := i + 1
			for j < len(src) {
				if src[j] == '\\' && j+1 < len(src) {
					j += 2
					continue
				}
				if src[j] == c {
					j++
					break
				}
				j++
			}
			spans = append(spans, Span{TagString, i, j})
			i = j
			continue
		}
		// Numbers (dec, hex, floats, leading-dot floats).
		if isDigit(c) || (c == '.' && i+1 < len(src) && isDigit(src[i+1])) {
			j := i
			if c == '0' && i+1 < len(src) && (src[i+1] == 'x' || src[i+1] == 'X') {
				j = i + 2
				for j < len(src) && isIdent(src[j]) {
					j++
				}
			} else {
				for j < len(src) && (isDigit(src[j]) || src[j] == '.' || src[j] == '_' ||
					src[j] == 'e' || src[j] == 'E' || src[j] == '+' || src[j] == '-') {
					j++
				}
			}
			spans = append(spans, Span{TagNumber, i, j})
			i = j
			continue
		}
		// Identifiers and keywords.
		if isIdent(c) {
			j := i
			for j < len(src) && isIdent(src[j]) {
				j++
			}
			word := src[i:j]
			tag := ""
			switch {
			case keywords[word]:
				tag = TagKeyword
			case types[word]:
				tag = TagType
			}
			// Function calls: identifier immediately followed by "(".
			if tag == "" {
				k := j
				for k < len(src) && (src[k] == ' ' || src[k] == '\t') {
					k++
				}
				if k < len(src) && src[k] == '(' {
					tag = TagFunc
				}
			}
			if tag != "" {
				spans = append(spans, Span{tag, i, j})
			}
			i = j
			continue
		}
		i++
	}
	return spans
}

// stringsHasPrefix reports whether s starts with prefix.
func stringsHasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	return s[:len(prefix)] == prefix
}
