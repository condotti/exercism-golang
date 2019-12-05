// Package variablelengthquantity implements a solution of the exercise titled `Variable Length Quantity'.
package variablelengthquantity

import "errors"

// EncodeVarint encodes an input with VLQ encoding.
func EncodeVarint(input []uint32) []byte {
	output := []byte{}
	for _, value := range input {
		base128 := []byte{byte(value & 0x7f)}
		for value >>= 7; value > 0; value >>= 7 {
			base128 = append([]byte{byte(value&0x7f) | 0x80}, base128...)
		}
		output = append(output, base128...)
	}
	return output
}

// DecodeVarint decodes an input with VLQ encoding, or an error when input is incomplete
func DecodeVarint(input []byte) ([]uint32, error) {
	output := []uint32{}
	if input[len(input)-1]&0x80 != 0 {
		return output, errors.New("incomplite input sequence")
	}
	var value uint32
	for _, octet := range input {
		if octet&0x80 == 0 {
			output = append(output, (value<<7)|uint32(octet))
			value = 0
		} else {
			value <<= 7
			value |= uint32(octet & 0x7f)
		}
	}
	return output, nil
}
