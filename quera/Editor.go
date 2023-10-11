package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []rune
}

func (s Stack) Top() (rune, error) {
	stackLen := len([]rune(s.data))
	var lastItem rune
	var err error
	if stackLen > 0 {
		lastItem = s.data[stackLen-1]
		err = nil
	} else {
		lastItem = ' '
		err = errors.New("")
	}
	return lastItem, err
}

type TwoStackEditor struct {
	LeftStack  *Stack
	RightStack *Stack
	// rest
}

func NewTwoStackEditor(initialText string) *TwoStackEditor {
	leftStack := &Stack{data: []rune{}}
	rightStack := &Stack{data: []rune{}}

	reverseText := reverseString(initialText)
	input := []rune(reverseText)
	for _, str := range input {
		rightStack.data = append(rightStack.data, str)
	}
	return &TwoStackEditor{LeftStack: leftStack, RightStack: rightStack}
}

func (e *TwoStackEditor) InsertRune(r rune) {
	e.LeftStack.data = append(e.LeftStack.data, r)
}

func (e *TwoStackEditor) InsertString(text string) {
	e.LeftStack.data = append(e.LeftStack.data, []rune(text)...)
}

func (e *TwoStackEditor) CursorLeft() {
	leftLen := len([]rune(e.LeftStack.data))
	if leftLen > 0 {
		firstLeftItem := e.LeftStack.data[leftLen-1]
		e.RightStack.data = append(e.RightStack.data, firstLeftItem)
		e.LeftStack.data = e.LeftStack.data[:leftLen-1]
	}
}

func (e *TwoStackEditor) CursorRight() {
	rightLen := len([]rune(e.RightStack.data))
	if rightLen > 0 {
		firstLeftItem := e.RightStack.data[rightLen-1]
		fmt.Println(string(firstLeftItem), "<=====")
		e.LeftStack.data = append(e.LeftStack.data, firstLeftItem)
		e.RightStack.data = e.RightStack.data[:rightLen-1]
	}
}

func (e *TwoStackEditor) DeleteForward() {
	rightStackLen := len([]rune(e.RightStack.data))
	if rightStackLen > 0 {
		e.RightStack.data = e.RightStack.data[:rightStackLen-1]
	}
}

func (e *TwoStackEditor) DeleteBackward() {
	leftLen := len([]rune(e.LeftStack.data))
	if leftLen > 0 {
		e.LeftStack.data = e.LeftStack.data[:leftLen-1]
	}
}

func (e *TwoStackEditor) WholeText() string {
	finalResult := []rune{}
	rightStackText := reverseString(string(e.RightStack.data))
	finalResult = append(finalResult, e.LeftStack.data...)
	finalResult = append(finalResult, []rune(rightStackText)...)
	return string(finalResult)
}

func reverseString(input string) string {
	runes := []rune(input)
	length := len([]rune(runes))
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}
