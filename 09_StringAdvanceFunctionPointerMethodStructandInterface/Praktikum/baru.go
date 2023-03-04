package main

import (
	"fmt"
	"strings"
)

type SubstitutionCipher struct {
	key string
}

func NewSubstitutionCipher(key string) *SubstitutionCipher {
	return &SubstitutionCipher{key: key}
}

func (sc *SubstitutionCipher) Encode(plainText string) string {
	var sb strings.Builder
	for _, ch := range plainText {
		if ch >= 'a' && ch <= 'z' {
			ch = rune(sc.key[ch-'a'])
		} else if ch >= 'A' && ch <= 'Z' {
			ch = rune(sc.key[ch-'A'] - 32)
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func (sc *SubstitutionCipher) Decode(cipherText string) string {
	invKey := make([]byte, 26)
	for i, ch := range sc.key {
		if ch >= 'a' && ch <= 'z' {
			invKey[ch-'a'] = byte('a' + i)
		} else if ch >= 'A' && ch <= 'Z' {
			invKey[ch-'A'] = byte('A' + i)
		}
	}
	var sb strings.Builder
	for _, ch := range cipherText {
		if ch >= 'a' && ch <= 'z' {
			ch = rune(invKey[ch-'a'])
		} else if ch >= 'A' && ch <= 'Z' {
			ch = rune(invKey[ch-'A'] + 32)
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

type student struct {
	name string
}

type Chiper interface {
	Encode(string) string
	Decode(string) string
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = NewSubstitutionCipher("qwertyuiopasdfghjklzxcvbnm")

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode(a.name))
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode(a.name))
	}
}
