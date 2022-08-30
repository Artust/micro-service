package validate

import (
	"net/mail"
	"strings"
	"unicode"
)

func ValidPassword(p string) bool {
	const minPassLength = 8
	const maxPassLength = 255
	const spChar = "+-^$*.[]{}()?\"!@#%&/\\,<>':;|_~`"
	var hasNumber bool
	var hasSpecialChar bool
	var hasUpper bool
	var hasLower bool
	var hasCharLen bool
	var passLen int

	for _, ch := range p {
		switch {
		case unicode.IsNumber(ch):
			hasNumber = true
			passLen++
		case unicode.IsUpper(ch):
			hasUpper = true
			passLen++
		case unicode.IsLower(ch):
			hasLower = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			if strings.Contains(spChar, string(ch)) {
				hasSpecialChar = true
			}
			passLen++
		}
	}
	if minPassLength <= passLen && passLen <= maxPassLength {
		hasCharLen = true
	}
	return hasUpper && hasLower && hasNumber && hasSpecialChar && hasCharLen
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
