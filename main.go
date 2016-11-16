package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
	"encoding/csv"
)

type Records [][]string
type FieldIndice map[string]int

var (
	infile   = flag.String("i", "", "File to be masked.")
	indelim  = flag.String("d", ",", "Delimiter of input file.")
	outfile  = flag.String("o", "", "File to be masked.")
	outdelim = flag.String("s", ",", "Delimiter of output file.")
	masklen  = flag.Int("l", 2, "Number of the letters to mask.")
	maskchar = flag.String("m", "X", "Character to be used as the mask.")
	fields   = flag.String("f", "no,header,val", "Path to the original CSV file to be masked.")
)

func readInputFile(inPath string, del rune) Records {
	inFile, err := os.Open(inPath)
	if err != nil {
		fmt.Println("ERROR: No input file specified :-(")
		panic(err)
	}
	defer inFile.Close()

	r := csv.NewReader(inFile)
	r.Comma = del
	r.LazyQuotes = true

	rcds, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return rcds
}

func writeAsCSV(outPath string, del rune, masked Records) int {
	outFile, err := os.Create(outPath)
	if err != nil {
		fmt.Println("ERROR: No output filename specified :-(")
		panic(err)
	}
	defer outFile.Close()

	w := csv.NewWriter(outFile)
	w.Comma = del

	w.WriteAll(masked)
	return 0
}

func getFieldIndex(flds []string, hdr []string) FieldIndice {
	hdrs := make(map[string]int)
	for _, f := range flds {
		for n, r := range hdr {
			if f == r {
				hdrs[r] = n
			}
		}
	}
	return hdrs
}

func maskLastLetters(s string, m string, l int) string {
	rem := s[0 : len(s)-l]
	msk := strings.Repeat(m, l)
	return rem + msk
}

func main() {
	flag.Parse()
	indel := []rune(*indelim)[0]
	outdel := []rune(*outdelim)[0]
	fields := strings.Split(*fields, ",")

	records := readInputFile(*infile, indel)
	indice := getFieldIndex(fields, records[0])

	masked := make(Records, len(records))
	for i, line := range records {
		for _, v := range indice {
			if (i == 0) || (len(line[v]) == 0) || (len(line[v]) < *masklen) {
			} else {
				line[v] = maskLastLetters(line[v], *maskchar, *masklen)
			}
		}
		masked[i] = line
	}

	result := writeAsCSV(*outfile, outdel, masked)
	if result == 0 {
		fmt.Println("Masking successfully finished!")
	} else {
		fmt.Println("Something wrong happened...")
	}
}
