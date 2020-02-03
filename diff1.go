package main

import (
	"flag"
	"fmt"
//	"io"
	"os"
	"bufio"
	"sort"
)

func main() {

//define program usage in this context
		flag.Usage = func() {
			fmt.Printf("This program takes a minimum of 2 files.\n")
			fmt.Printf("Usage of %s: \n", os.Args[0])
			fmt.Println("catfile [options] file1 file2 \n")
			flag.PrintDefaults()
		}
	flag.Parse()
//Check if amount of arguments remaining is legal
	if flag.NArg() < 2  {
		flag.Usage()
		os.Exit(1)
	}
	filesToCompare := []string{flag.Arg(len(flag.Args())-2),flag.Arg(len(flag.Args())-1)}

	file1,err1 := os.Open(filesToCompare[0])
	file2,err2 := os.Open(filesToCompare[1])

	if err1 != nil { panic(err1) }	
	if err2 != nil { panic(err2) } 
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)
	
	scanner1.Split(bufio.ScanLines)
	scanner2.Split(bufio.ScanLines)

	var file1lines []string
	var file2lines []string

	for scanner1.Scan() {
		file1lines = append(file1lines, scanner1.Text())			
	}

	for scanner2.Scan() {
		file2lines = append(file2lines, scanner2.Text())
	}	

//	fmt.Print("The comparison!!!\n")

		var comparisonMatrix []int
	if len(file1lines) > len(file2lines) {
		for lineNumber:=1; lineNumber<len(file2lines); lineNumber++ {
			if file1lines[lineNumber-1] != file2lines[lineNumber-1] {
				comparisonMatrix = append(comparisonMatrix, lineNumber)
			}
		}
	} else {
		for lineNumber:=1; lineNumber<len(file1lines); lineNumber++ {
			if file2lines[lineNumber-1] != file1lines[lineNumber-1] {
				comparisonMatrix = append(comparisonMatrix, lineNumber)
			} 
		}
	}
	sort.Ints(comparisonMatrix)

	for lineNo,_ := range comparisonMatrix {
		fmt.Printf("%d \n", comparisonMatrix[lineNo])
	}

}
