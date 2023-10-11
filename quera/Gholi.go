package main

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Phonebook struct {
	data map[string]string
}

func NewPhonebook() *Phonebook {
	return &Phonebook{
		data: make(map[string]string),
	}
}

func (p *Phonebook) GetPhoneNumber(contactName string) (string, error) {
	contact := p.data[contactName]
	if contact != "" {
		return contact, nil
	}

	err := errors.New(" ")
	return "", err
}

func (p *Phonebook) AddContact(contactName string, phoneNumber string) string {
	contact := p.data[contactName]
	if contact == phoneNumber {
		return "No Change"
	} else if contact != "" {
		p.data[contactName] = phoneNumber
		return "Phone Number Updated"
	} else {
		p.data[contactName] = phoneNumber
		return "Phone Number Added"
	}
}

func (p *Phonebook) GetAllContacts() []string {
	contactsName := []string{}
	for k, _ := range p.data {
		contactsName = append(contactsName, k)
	}
	return contactsName
}

func (p *Phonebook) GetAllContactsStartsWith(letter rune) []string {
	contactsName := []string{}
	for k, _ := range p.data {
		hasPrefix := strings.HasPrefix(k, string(letter))
		if hasPrefix {
			contactsName = append(contactsName, k)
		}
	}
	return contactsName
}

func (p *Phonebook) DeleteInvalidPhoneNumbers() {
	for k, v := range p.data {
		pattern := `^(?:\+|0)\d{10}$`
		regex := regexp.MustCompile(pattern)
		if !regex.MatchString(v) {
			delete(p.data, k)
		}
	}
}

func (p *Phonebook) DeleteDuplicatePhoneNumbers() {
	data := p.data
	allDuplicateContact := map[string][]string{}
	for k1, v1 := range data {
		duplicateContact := []string{}
		for k2, v2 := range data {
			firstValue := v1[1:]
			secondValue := v2[1:]
			if firstValue == secondValue && k1 != k2 {
				duplicateContact = append(duplicateContact, k1)
				duplicateContact = append(duplicateContact, k2)
			}
		}
		allDuplicateContact[v1[1:]] = duplicateContact
	}

	for _, v := range allDuplicateContact {
		duplicatedContact := v
		if len(v) > 0 {
			sort.Strings(duplicatedContact)
			for i, name := range duplicatedContact {
				if i >= 1 {
					delete(p.data, name)
				}
			}
		}
	}
}

// func (p *Phonebook) FindCatchyPhoneNumbers() []string {
// 	phoneBook := p.data
// 	catchyNumber := map[string]int{}
// 	for _, number := range phoneBook {
// 		catchyNumber[number] = 0
// 		trimNumber := number[2:]
// 		if trimNumber[0] == trimNumber[1] && trimNumber[1] == trimNumber[2] {
// 			catchyNumber[number] = catchyNumber[number] + 10
// 		} else if trimNumber[1] == trimNumber[2] && trimNumber[2] == trimNumber[3] {
// 			catchyNumber[number] = catchyNumber[number] + 10
// 		} else if trimNumber[2] == trimNumber[3] && trimNumber[3] == trimNumber[4] {
// 			catchyNumber[number] = catchyNumber[number] + 10
// 		} else if trimNumber[3] == trimNumber[4] && trimNumber[4] == trimNumber[5] {
// 			catchyNumber[number] = catchyNumber[number] + 10
// 		} else if trimNumber[4] == trimNumber[5] && trimNumber[5] == trimNumber[6] {
// 			catchyNumber[number] = catchyNumber[number] + 10
// 		} else if trimNumber[5] == trimNumber[6] && trimNumber[6] == trimNumber[7] {
// 			catchyNumber[number] = catchyNumber[number] + 10
// 		} else if trimNumber[6] == trimNumber[7] && trimNumber[7] == trimNumber[8] {
// 			catchyNumber[number] = catchyNumber[number] + 10
// 		}

// 		f1, _ := strconv.Atoi(trimNumber[0:2])
// 		f2, _ := strconv.Atoi(trimNumber[2:4])
// 		s1, _ := strconv.Atoi(trimNumber[1:3])
// 		s2, _ := strconv.Atoi(trimNumber[3:5])
// 		t1, _ := strconv.Atoi(trimNumber[2:4])
// 		t2, _ := strconv.Atoi(trimNumber[4:6])
// 		fo1, _ := strconv.Atoi(trimNumber[3:5])
// 		fo2, _ := strconv.Atoi(trimNumber[5:7])
// 		fi1, _ := strconv.Atoi(trimNumber[4:6])
// 		fi2, _ := strconv.Atoi(trimNumber[6:8])
// 		si1, _ := strconv.Atoi(trimNumber[5:7])
// 		si2, _ := strconv.Atoi(trimNumber[7:])
// 		if f1 < f2 && f1+1 == f2 || f1 > f2 && f1-1 == f2 {
// 			catchyNumber[number] = catchyNumber[number] + 15
// 		} else if s1 < s2 && s1+1 == s2 || s1 > s2 && s1-1 == s2 {
// 			catchyNumber[number] = catchyNumber[number] + 15
// 		} else if t1 < t2 && t1+1 == t2 || t1 > t2 && t1-1 == t2 {
// 			catchyNumber[number] = catchyNumber[number] + 15
// 		} else if fo1 < fo2 && fo1+1 == fo2 || fo1 > fo2 && fo1-1 == fo2 {
// 			catchyNumber[number] = catchyNumber[number] + 15
// 		} else if fi1 < fi2 && fi1+1 == fi2 || fi1 > fi2 && fi1-1 == fi2 {
// 			catchyNumber[number] = catchyNumber[number] + 15
// 		} else if si1 < si2 && si1+1 == si2 || si1 > si2 && si1-1 == si2 {
// 			catchyNumber[number] = catchyNumber[number] + 15
// 		}

// 		if trimNumber[0] == trimNumber[3] && trimNumber[1] == trimNumber[2] {
// 			catchyNumber[number] = catchyNumber[number] + 20
// 		} else if trimNumber[2] == trimNumber[5] && trimNumber[3] == trimNumber[4] {
// 			catchyNumber[number] = catchyNumber[number] + 20
// 		} else if trimNumber[4] == trimNumber[7] && trimNumber[5] == trimNumber[6] {
// 			catchyNumber[number] = catchyNumber[number] + 20
// 		} else if trimNumber[1] == trimNumber[4] && trimNumber[2] == trimNumber[3] {
// 			catchyNumber[number] = catchyNumber[number] + 20
// 		} else if trimNumber[3] == trimNumber[6] && trimNumber[4] == trimNumber[5] {
// 			catchyNumber[number] = catchyNumber[number] + 20
// 		} else if trimNumber[5] == trimNumber[8] && trimNumber[6] == trimNumber[7] {
// 			catchyNumber[number] = catchyNumber[number] + 20
// 		}
// 	}

// 	var catchyNumberSlice []struct {
// 		Key   string
// 		Value int
// 	}

// 	for k, v := range catchyNumber {
// 		catchyNumberSlice = append(catchyNumberSlice, struct {
// 			Key   string
// 			Value int
// 		}{k, v})
// 	}

// 	sort.Slice(catchyNumberSlice, func(i, j int) bool {
// 		return catchyNumberSlice[i].Value > catchyNumberSlice[j].Value
// 	})

// 	finalResult := []string{}
// 	var thirdItem int
// 	catchyNumberSliceLen := len(catchyNumberSlice)
// 	for k, v := range catchyNumberSlice {
// 		if catchyNumberSliceLen <= 3 {
// 			finalResult = append(finalResult, v.Key)
// 		} else {
// 			if k <= 2 {
// 				finalResult = append(finalResult, v.Key)
// 			}
// 			if k == 2 {
// 				thirdItem = v.Value
// 			}
// 			if k > 2 && v.Value == thirdItem {
// 				finalResult = append(finalResult, v.Key)
// 			}
// 		}
// 	}

// 	return finalResult
// }

func (p *Phonebook) FindCatchyPhoneNumbers() []string {
	catchyNumber := make(map[string]int)
	for _, number := range p.data {
		trimmedNumber := number[2:]
		score := 0

		checkRepeatingDigits := func(s string) bool {
			for i := 0; i < len(s)-2; i++ {
				if s[i] == s[i+1] && s[i] == s[i+2] {
					return true
				}
			}
			return false
		}

		checkSequentialDigits := func(s string) bool {
			for i := 0; i < len(s)-1; i++ {
				digit1, _ := strconv.Atoi(string(s[i]))
				digit2, _ := strconv.Atoi(string(s[i+1]))
				if abs(digit1-digit2) == 1 {
					return true
				}
			}
			return false
		}

		checkMirrorPattern := func(s string) bool {
			for i := 0; i < len(s)-3; i++ {
				if s[i] == s[i+3] && s[i+1] == s[i+2] {
					return true
				}
			}
			return false
		}

		if checkRepeatingDigits(trimmedNumber) {
			score += 10
		}

		if checkSequentialDigits(trimmedNumber) {
			score += 15
		}

		if checkMirrorPattern(trimmedNumber) {
			score += 20
		}

		catchyNumber[number] = score
	}

	var catchyNumberSlice []struct {
		Key   string
		Value int
	}

	for k, v := range catchyNumber {
		catchyNumberSlice = append(catchyNumberSlice, struct {
			Key   string
			Value int
		}{k, v})
	}

	sort.Slice(catchyNumberSlice, func(i, j int) bool {
		return catchyNumberSlice[i].Value > catchyNumberSlice[j].Value
	})

	finalResult := []string{}
	for k, v := range catchyNumberSlice {
		if k < 3 {
			finalResult = append(finalResult, v.Key)
		} else if k == 3 {
			break
		}
	}
	return finalResult

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//https://go.dev/doc/faq#closures_and_goroutines
//https://www.google.com/search?sca_esv=561558033&q=parallel+vs+concurrent&tbm=isch&source=lnms&sa=X&ved=2ahUKEwjipa-vr4aBAxVCTaQEHTt7DZMQ0pQJegQICRAB&biw=1920&bih=963&dpr=1#imgrc=I_hqMnOw7XXWDM
//https://go.dev/doc/effective_go
//https://www.youtube.com/@7bjar/videos
//https://www.sqlite.org/testing.html
// https://www.google.com/search?q=instroment+coverage&oq=instroment+coverage&aqs=chrome..69i57.10615j0j1&sourceid=chrome&ie=UTF-8
