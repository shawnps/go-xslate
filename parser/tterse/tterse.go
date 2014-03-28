package tterse

import (
  "github.com/lestrrat/go-lex"
  "github.com/lestrrat/go-xslate/parser"
)

// SymbolSet contains TTerse specific symbols
var SymbolSet = parser.DefaultSymbolSet.Copy()
func init() {
 // "In" must come before Include
  SymbolSet.Set("INCLUDE",  parser.ItemInclude, 2.0)
  SymbolSet.Set("IN",       parser.ItemIn,      1.5)
  SymbolSet.Set("WITH",     parser.ItemWith)
  SymbolSet.Set("CALL",     parser.ItemCall)
  SymbolSet.Set("END",      parser.ItemEnd)
  SymbolSet.Set("WRAPPER",  parser.ItemWrapper)
  SymbolSet.Set("SET",      parser.ItemSet)
  SymbolSet.Set("GET",      parser.ItemGet)
  SymbolSet.Set("IF",       parser.ItemIf)
  SymbolSet.Set("ELSIF",    parser.ItemElseIf)
  SymbolSet.Set("ELSE",     parser.ItemElse)
  SymbolSet.Set("UNLESS",   parser.ItemUnless)
  SymbolSet.Set("FOREACH",  parser.ItemForeach)
  SymbolSet.Set("WHILE",    parser.ItemWhile)
  SymbolSet.Set("MACRO",    parser.ItemMacro)
  SymbolSet.Set("BLOCK",    parser.ItemBlock)
  SymbolSet.Set("END",      parser.ItemEnd)
}

// TTerse is the main parser for TTerse
type TTerse struct {
  items []lex.LexItem
}

// NewLexer creates a new lexer
func NewLexer(template string) *parser.Lexer {
  l := parser.NewLexer(template, SymbolSet)
  l.SetTagStart("[%")
  l.SetTagEnd("%]")

  return l
}

// New creates a new TTerse parser
func New() *TTerse {
  return &TTerse {}
}

// Parse parses the given template and creates an AST
func (p *TTerse) Parse(template []byte) (*parser.AST, error) {
  return p.ParseString(string(template))
}

// ParseString is the same as Parse, but receives a string instead of []byte
func (p *TTerse) ParseString(template string) (*parser.AST, error) {
  b := parser.NewBuilder()
  lex := NewLexer(template)
  return b.Parse("foo", template, lex)
}
