/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” 
in the string s. The wc.Test function runs a test suite against the provided
 function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res_map := make(map[string]int)
	w_list := strings.Fields(s)
	for _, v := range w_list {
		res_map[v] += 1
	}
	return res_map
}

func main() {
	wc.Test(WordCount)
}
