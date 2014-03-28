package kolonish

import (
  "testing"
  "github.com/lestrrat/go-lex"
  "github.com/lestrrat/go-xslate/parser"
)

func makeItem(t lex.LexItemType, p int, v string) lex.LexItem {
  return lex.NewLexItem(t, p, v)
}

var space     = makeItem(parser.ItemSpace, 0, " ")
var tagStart  = makeItem(parser.ItemTagStart, 0, "[%")
var tagEnd    = makeItem(parser.ItemTagEnd, 0, "[%")
var dollar    = makeItem(ItemDollar, 0, "")
func makeLexer(input string) *parser.Lexer {
  l := NewLexer(input)
  return l
}

func lexit(input string) *parser.Lexer {
  l := makeLexer(input)
  go l.Run(l)
  return l
}

func compareLex(t *testing.T, expected []lex.LexItem, l *parser.Lexer) {
  for n := 0; n < len(expected); n++ {
    i := l.NextItem()

    e := expected[n]
    if e.Type() != i.Type() {
      t.Errorf("Expected type %s, got %s", e.Type(), i.Type())
      t.Logf("   -> expected %s got %s", e, i)
    }
    if e.Type() == parser.ItemIdentifier || e.Type() == parser.ItemRawString {
      if e.Value() != i.Value() {
        t.Errorf("Expected.Value()ue %s, got %s", e.Value(), i.Value())
        t.Logf("   -> expected %s got %s", e, i)
      }
    }
  }

  i := l.NextItem()
  if i.Type() != parser.ItemEOF {
    t.Errorf("Expected EOF, got %s", i)
  }

}

func TestGetImplicit(t *testing.T) {
  tmpl  := `<: $foo :>`
  l     := lexit(tmpl)
  expected := []lex.LexItem {
    tagStart,
    space,
    dollar,
    makeItem(parser.ItemIdentifier, 0, "foo"),
    space,
    tagEnd,
  }
  compareLex(t, expected, l)
}

func TestLinewiseCode(t *testing.T) {
  tmpl := `
: "foo\n"
: for list -> i {
:    i
: }
`
  _ = lexit(tmpl)

}