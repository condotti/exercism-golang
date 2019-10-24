// Package romannumerals implements a solution of exercise titled `Roman Numerals'
package romannumerals

import "errors"

// ToRomanNumeral returns given N as roman numeral representation, or an error.
func ToRomanNumeral(n int) (string, error) {
	if 0 < n && n <= 3000 {
		roman := ""
		for i := 0; n > 0; i, n = i+1, n/10 {
			roman = func(n int, rd string) string {
				i, v, x := rd[:1], rd[1:2], rd[2:3]
				return []string{"", i, i + i, i + i + i, i + v,
					v, v + i, v + i + i, v + i + i + i, i + x}[n]
			}(n%10, "IVXLCDM  "[i*2:]) + roman
		}
		return roman, nil
	}
	return "", errors.New("invalid input")
}
