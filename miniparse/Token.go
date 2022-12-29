package main

type Token struct{
	_type TokenType
	text []rune
}

func NewToken(_type TokenType, text []rune) *Token {
	return &Token{
		_type: _type,
		text: text,
	}
}


