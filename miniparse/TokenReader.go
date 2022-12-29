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

func (str *SimpleTokenReader) Read() Token {

}

func (str *SimpleTokenReader) Peek() Token {

}

func (str *SimpleTokenReader) Unread() Token {

}

