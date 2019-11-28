// Package cipher implemets a solution of the exercise titled `Simple Cipher'.
package cipher

// Shift represents a key (shift distance) of shift cipher.
type Shift rune

// NewCaesar is the ctor of caesar cipher, which is actuall shift cipher.
func NewCaesar() Cipher {
	return Shift(3)
}

// NewShift is the ctor of shift cipher.
func NewShift(distance int) Cipher {
	if 1 <= distance*distance && distance*distance < 26*26 {
		return Shift(distance)
	}
	return nil
}

func isLower(c rune) bool {
	return 'a' <= c && c <= 'z'
}

func isUpper(c rune) bool {
	return isLower(c - 'A' + 'a')
}

func nonNegMod(a, b rune) rune {
	return (b + (a % b)) % b
}

// Encode encodes source string specified.
func (c Shift) Encode(source string) string {
	cipher := []rune{}
	for _, letter := range source {
		switch {
		case isLower(letter):
			cipher = append(cipher, nonNegMod(letter-'a'+rune(c), 26)+'a')
		case isUpper(letter):
			cipher = append(cipher, nonNegMod(letter-'A'+rune(c), 26)+'a')
		}
	}
	return string(cipher)
}

// Decode decodes cipher string specified.
func (c Shift) Decode(cipher string) string {
	plain := []rune{}
	// only lowercase alphabets exist in theory
	for _, letter := range cipher {
		plain = append(plain, nonNegMod(letter-'a'-rune(c), 26)+'a')
	}
	return string(plain)
}

// Vigenere represents a key (string to shift) of vigenere cipher.
type Vigenere []rune

// NewVigenere is the ctor of vigenere cipher.
func NewVigenere(key string) Cipher {
	rkey := []rune{}
	nonZero := false
	for _, letter := range key {
		if !isLower(letter) {
			return nil
		}
		rkey = append(rkey, letter-'a')
		nonZero = nonZero || letter-'a' != 0
	}
	if nonZero {
		return Vigenere(rkey)
	}
	return nil
}

// Encode encodes source string with vigenere cipher.
func (c Vigenere) Encode(source string) string {
	rkey := []rune(c)
	cipher := []rune{}
	index := 0
	for _, letter := range source {
		switch {
		case isLower(letter):
			cipher = append(cipher, nonNegMod(letter-'a'+rkey[index%len(rkey)], 26)+'a')
			index++
		case isUpper(letter):
			cipher = append(cipher, nonNegMod(letter-'A'+rkey[index%len(rkey)], 26)+'a')
			index++
		}
	}
	return string(cipher)
}

// Decode decodes cipher string encoded with vigenere cipher.
func (c Vigenere) Decode(cipher string) string {
	rkey := []rune(c)
	plain := []rune{}
	// only lowercase alphabets exist in theory
	for index, letter := range cipher {
		plain = append(plain, nonNegMod(letter-'a'-rkey[index%len(rkey)], 26)+'a')
	}
	return string(plain)
}
