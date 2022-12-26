package main

type TokenReader interface {
	//返回下一个Token
	Read() Token

	//从stream读取下一个
	Peek() Token

}