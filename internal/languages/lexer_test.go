package languages

import "testing"

func spansOf(tag string, spans []Span) []int {
	var out []int
	for _, s := range spans {
		if s.Tag == tag {
			out = append(out, s.Start)
		}
	}
	return out
}

func TestLexGoKeywords(t *testing.T) {
	src := "func main() {\n\treturn 42\n}"
	spans := Lex(src, DefaultLanguage)
	kws := spansOf(TagKeyword, spans)
	if len(kws) != 2 {
		t.Fatalf("expected 2 keyword spans (func, return), got %d (%v)", len(kws), kws)
	}
	if src[kws[0]:kws[0]+4] != "func" {
		t.Fatalf("first keyword should be func")
	}
}

func TestLexFunctionCall(t *testing.T) {
	src := "result := strings.TrimSpace(input)"
	spans := Lex(src, DefaultLanguage)
	funcs := spansOf(TagFunc, spans)
	if len(funcs) != 1 {
		t.Fatalf("expected 1 function span, got %d (%v)", len(funcs), funcs)
	}
	if got := src[funcs[0]:]; len(got) < 9 || got[:9] != "TrimSpace" {
		t.Fatalf("expected TrimSpace highlighted as func, got %q", got)
	}
}

func TestLexStringsAndComments(t *testing.T) {
	src := "// a line comment\ns := \"hello \\\" world\"\n/* block\ncomment */"
	spans := Lex(src, DefaultLanguage)

	if len(spansOf(TagComment, spans)) != 2 {
		t.Fatalf("expected 2 comment spans, got %v", spansOf(TagComment, spans))
	}
	strs := spansOf(TagString, spans)
	if len(strs) != 1 {
		t.Fatalf("expected 1 string span (escaped quote handled), got %d (%v)", len(strs), strs)
	}
	want := "\"hello \\\" world\""
	got := src[strs[0]:]
	if len(got) < len(want) || got[:len(want)] != want {
		t.Fatalf("string span should cover the full escaped literal, got %q", got)
	}
}

func TestLexTypesAndNumbers(t *testing.T) {
	src := "var x int64 = 0x1F\nf := 3.14"
	spans := Lex(src, DefaultLanguage)
	if len(spansOf(TagType, spans)) != 1 {
		t.Fatalf("expected int64 highlighted as type")
	}
	if len(spansOf(TagNumber, spans)) != 2 {
		t.Fatalf("expected 2 number spans (0x1F, 3.14)")
	}
}

func TestLexRust(t *testing.T) {
	rust := NewRegistry().Lookup("rust")
	src := "fn main() {\n    let x: i32 = 42;\n    // comment\n    println!(\"hi\");\n}"
	spans := Lex(src, rust)
	if len(spansOf(TagKeyword, spans)) < 2 { // fn, let
		t.Fatalf("expected rust keywords, got %v", spansOf(TagKeyword, spans))
	}
	if len(spansOf(TagType, spans)) < 1 { // i32
		t.Fatalf("expected i32 as type, got %v", spansOf(TagType, spans))
	}
	if len(spansOf(TagComment, spans)) != 1 {
		t.Fatalf("expected rust comment")
	}
	if len(spansOf(TagString, spans)) != 1 {
		t.Fatalf("expected rust string")
	}
}

func TestLexPythonHashComment(t *testing.T) {
	py := NewRegistry().Lookup("python")
	src := "# comment\ndef f(x):\n    return x"
	spans := Lex(src, py)
	if len(spansOf(TagComment, spans)) != 1 {
		t.Fatalf("expected python # comment")
	}
	if len(spansOf(TagKeyword, spans)) < 2 { // def, return
		t.Fatalf("expected python keywords, got %v", spansOf(TagKeyword, spans))
	}
}

func TestLexCSharpKeywordsAndStrings(t *testing.T) {
	cs := NewRegistry().Lookup("csharp")
	src := "public class Foo {\n    string s = \"hi\";\n    return;\n}"
	spans := Lex(src, cs)
	if len(spansOf(TagKeyword, spans)) < 2 { // public, class
		t.Fatalf("expected csharp keywords, got %v", spansOf(TagKeyword, spans))
	}
	if len(spansOf(TagString, spans)) != 1 {
		t.Fatalf("expected csharp string")
	}
}

func TestLexNilLanguageUsesDefault(t *testing.T) {
	src := "func f() {}"
	spans := Lex(src, nil)
	if len(spansOf(TagKeyword, spans)) != 1 {
		t.Fatalf("nil language should lex as Go")
	}
}

func TestLexBacktickRawString(t *testing.T) {
	src := "s := `raw\nstring`\n// after"
	spans := Lex(src, DefaultLanguage)
	strs := spansOf(TagString, spans)
	if len(strs) != 1 {
		t.Fatalf("expected raw string span, got %v", strs)
	}
	if len(spansOf(TagComment, spans)) != 1 {
		t.Fatalf("expected trailing comment")
	}
}
