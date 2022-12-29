package main

type TokenReader interface {
	Read() Token
	Peek() Token
	Unread()
	GetPosition() int
	SetPositon(int)
}

//接口的实现
type SimpleTokenReader struct {
	tokens []Token
	pos int
}

func NewTokenReader(tokens []Token) *SimpleTokenReader {
	return &SimpleTokenReader{
		tokens: tokens,
		pos: 0,
	}
}

func (str *SimpleTokenReader) Read() Token {

}

func (str *SimpleTokenReader) Peek() Token {

}

func (str *SimpleTokenReader) Unread() Token {

}

