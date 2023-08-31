package hex

// Hex manipulation functions
//
// Usage:
// import {
//   "src/github.com/guydavies/hex"
// }
// foo = hex.prepend_to_length(new_hex_number, hex_digits_per_group)
// bar = hex.generate_hex(required_number_of_hex_digits)
//
// Author: Guy Davies
// Email: aguydavies@gmail.com

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func prepend_to_length(new_hex_number string, hex_digits_per_group int) string {
	// check the length of the produced string and prepend 0s to expected length
	if len(new_hex_number) < hex_digits_per_group {
		var hex_strings = [2]string{"0", new_hex_number}
		new_hex_number = strings.Join(hex_strings[:], "")
		var padded_hex_number string = prepend_to_length(new_hex_number, hex_digits_per_group)
		return padded_hex_number
	} else {
		return new_hex_number
	}
}

func Generate_hex(hex_digits_per_group int) string {
	// function to generate a single hex number of length hex_digits_per_group
	// lots of rather ugly type conversions in order to be able to manipulate the objects
	var rand_hex_str string = fmt.Sprintf("%x", (rand.Intn(int(math.Round(math.Pow(2, (4.0*float64(hex_digits_per_group))))) - 1)))
	var padded_hex_number string = prepend_to_length(rand_hex_str, hex_digits_per_group)
	return padded_hex_number
}
