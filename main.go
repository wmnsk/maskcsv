package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Records [][]string
type FieldIndex map[string]int

var (
	infile   = flag.String("i", "", "File to be masked.")
	indelim  = flag.String("d", ",", "Delimiter of input file.")
	outfile  = flag.String("o", "", "File to be masked.")
	outdelim = flag.String("s", ",", "Delimiter of output file.")
	masklen  = flag.Int("l", 2, "Number of the letters to mask.")
	maskchar = flag.String("m", "X", "Character to be used as the mask.")
	fields   = flag.String("f", "no,header,val", "Path to the original CSV file to be masked.")
)

func readInputFile(inPath string, del rune) (Records, error) {
	inFile, err := os.Open(inPath)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	r := csv.NewReader(inFile)
	r.Comma = del
	r.LazyQuotes = true

	rcds, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return rcds, nil
}

func writeAsCSV(outPath string, del rune, masked Records) error {
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := csv.NewWriter(outFile)
	w.Comma = del

	w.WriteAll(masked)
	return nil
}

func getFieldIndex(flds []string, hdr []string) FieldIndex {
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

	records, err := readInputFile(*infile, indel)
	if err != nil {
		fmt.Printf("Failed reading CSV file: %s\n", err)
		flag.Usage()
		os.Exit(1)
	}

	indexes := getFieldIndex(fields, records[0])
	masked := make(Records, len(records))
	for i, line := range records {
		for _, v := range indexes {
			if (i == 0) || (len(line[v]) == 0) || (len(line[v]) < *masklen) {
			} else {
				line[v] = maskLastLetters(line[v], *maskchar, *masklen)
			}
		}
		masked[i] = line
	}

	if err := writeAsCSV(*outfile, outdel, masked); err != nil {
		fmt.Printf("Failed writing CSV file: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Masking successfully finished!")
}
