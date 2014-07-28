// analyzer
package main

import (
	"C"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var query_field_id = -1

var query_field *string = flag.String("f", "field_name", "The field name to search for")
var query_values *string = flag.String("v", "field_values", "The values to search for, separated by commas")
var print_fields_str *string = flag.String("p", "print_fields", "The field(s) you want to get printed out, separated by commas")
var delimiter *string = flag.String("d", ",", "Field separation character")

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	flag.Parse()

	csvReader := csv.NewReader(os.Stdin)
	csvReader.Comma = []rune(*delimiter)[0]

	nextLine, e := csvReader.Read()

	out_fields := strings.Split(*print_fields_str, ",")

	target_fields := make(map[int]bool)

	for i, field := range nextLine {
		if *query_field == field {
			query_field_id = i
		}

		if stringInSlice(field, out_fields) {
			target_fields[i] = true
		} else {
			target_fields[i] = false
		}
	}

	if query_field_id == -1 {
		fmt.Fprintf(os.Stderr, "Can't find the fields to search for.\n")
		os.Exit(1)
	}

	query_values_slice := strings.Split(*query_values, ",")

	nextLine, e = csvReader.Read()

	for e == nil {
		if stringInSlice(nextLine[query_field_id], query_values_slice) {
			first := true
			for i, val := range nextLine {
				if target_fields[i] == true {
					if first == true {
						first = false
					} else {
						fmt.Print(",")
					}
					fmt.Print(val)
				}
			}
			fmt.Print("\n")
		}
		nextLine, e = csvReader.Read()
	}
}
