package main

import (
	"errors"
)

type Stack struct {
	data []rune
}

func NewStack() *Stack {
	var s Stack
	return &s
}

func (s *Stack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *Stack) Pop() error {
	var err error
	dataLen := len(s.data)
	if dataLen > 0 {
		s.data = s.data[0 : dataLen-1]
		return nil
	}
	err = errors.New("")
	return err
}

func (s Stack) Top() (rune, error) {
	stackLen := len(s.data)
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

func (s Stack) Size() int {
	stackLen := len(s.data)
	return stackLen
}
