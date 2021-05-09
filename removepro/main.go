package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	input := ScanTargets()
	var result []string
	if !ScanFlag() {
		for _, elem := range input {
			result = append(result, RemoveHeaders(elem))
		}
		result = removeDuplicateValues(result)
		for _, elem := range result {
			fmt.Println(elem)
		}
	} else {
		for _, elem := range input {
			sub := RemovePort(RemoveHeaders(GetOnlySubs(elem)))
			if sub != "" {
				result = append(result, sub)
			}
		}
		result = removeDuplicateValues(result)
		for _, elem := range result {
			fmt.Println(elem)
		}
	}

}

//ScanInput return the array of elements
//taken as input on stdin.
func ScanTargets() []string {

	var result []string

	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return result
}

//removeDuplicateValues
func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//RemoveHeaders
func RemoveHeaders(input string) string {
	res := strings.Index(input, "://")
	if res >= 0 {
		return input[res+3:]
	} else {
		return input
	}
}

//GetOnlySubs
func GetOnlySubs(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	return u.Host
}

//ScanFlag
func ScanFlag() bool {
	subsPtr := flag.Bool("subs", false, "Return only subdomains without protocols.")
	flag.Parse()
	return *subsPtr
}

//RemovePort
func RemovePort(input string) string {
	res := strings.Index(input, ":")
	if res >= 0 {
		return input[:res-1]
	}
	return input
}
