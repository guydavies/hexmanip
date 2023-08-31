package hexmanip

// Hex manipulation functions
//
// Usage:
// import {
//   "src/github.com/guydavies/hex"
// }
// bar = hex.Generate_hex(required_number_of_hex_digits)
//
// Author: Guy Davies
// Email: aguydavies@gmail.com

import (
	"fmt"
	"math/rand"
	"strings"
)

func prepend_to_length(new_hex_number string, max_hex_value_as_int int) string {
	// check the length of the produced string and prepend 0s to expected length
	// var desired_group_length = len(fmt.Sprintf("%x", max_hex_valueas_int))
	if len(new_hex_number) < len(fmt.Sprintf("%x", max_hex_value_as_int)) {
		var hex_strings = [2]string{"0", new_hex_number}
		new_hex_number = strings.Join(hex_strings[:], "")
		var padded_hex_number string = prepend_to_length(new_hex_number, max_hex_value_as_int)
		return padded_hex_number
	} else {
		return new_hex_number
	}
}

func Generate_hex(max_hex_value_as_int int) string {
	// function to generate a single random hex number from all zeroes upto max_value_as_int
	// lots of rather ugly type conversions in order to be able to manipulate the objects
	var rand_hex_str string = fmt.Sprintf("%x", (rand.Intn(max_hex_value_as_int + 1)))
	var padded_hex_number string = prepend_to_length(rand_hex_str, max_hex_value_as_int)
	return padded_hex_number
}
