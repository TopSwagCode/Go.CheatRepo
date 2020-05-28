package main

import (
	"flag"
	"fmt"
	"os"
)

var validMetrics = [3]string{"chars", "words", "lines"}

func main() {
	textPointer := flag.String("text", "", "Text <Text to parse> (Required)")
	metricPointer := flag.String("metric", "chars", "Metric <{chars|words|lines}> (Required)")
	uniquePointer := flag.Bool("unique", false, "Measure unique values of a metric.")

	flag.Parse()

	if *textPointer == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if stringInSlice(*metricPointer, validMetrics[0:3]) == false {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("textPointer: %s, metricPointer: %s, uniquePointer: %t\n", *textPointer, *metricPointer, *uniquePointer)
}


func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}