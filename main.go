package main

import (
	"fmt"
	"strconv"
)

const NumberTokenType = "Number"
const PlusTokenType = "Plus"
const EOFTokenType = "EOF"

type Token struct {
	Typo  string
	Value string
}

type Interpretor struct {
	Input        string
	pos          int
	currentToken *Token
}

func (i *Interpretor) Expr() (string, error) {
	left := i.getNextToken()
	i.eat(NumberTokenType)
	i.eat(PlusTokenType)
	right := i.getNextToken()
	i.eat(NumberTokenType)
	leftNumber, err := strconv.Atoi(left.Value)
	if err != nil {
		return "", err
	}
	rightNumber, err := strconv.Atoi(right.Value)
	if err != nil {
		return "", err
	}
	result := leftNumber + rightNumber
	return strconv.Itoa(result), nil
}

func (i *Interpretor) getNextToken() *Token {
	if i.pos == len(i.Input) {
		i.currentToken = &Token{
			Typo: EOFTokenType,
		}
		return i.currentToken
	}
	if string(i.Input[i.pos]) == "+" {
		i.currentToken = &Token{
			Typo:  PlusTokenType,
			Value: "+",
		}
		return i.currentToken
	}
	result := ""
	for ; i.pos < len(i.Input); i.pos++ {
		if _, err := strconv.Atoi(string((i.Input[i.pos]))); err != nil {
			break
		}
		result += string((i.Input[i.pos]))
	}
	i.currentToken = &Token{
		Typo:  NumberTokenType,
		Value: result,
	}
	return i.currentToken
}

func (i *Interpretor) eat(typo string) error {
	if i.currentToken.Typo != typo {
		return fmt.Errorf("expect %s, but get %s", typo, i.currentToken.Typo)
	}
	i.pos++
	return nil
}

func main() {
	input := "12+35"
	interpretor := &Interpretor{
		Input: input,
	}
	result, err := interpretor.Expr()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%s = %s\n", input, result)
}
