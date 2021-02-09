package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var (
	// input directory
	input string

	// start index
	start int

	// prefix
	prefix string

	// forward 0 digits
	align int
)

func init()  {
	flag.StringVar(&input, "in", "./", "Input directory")
	flag.IntVar(&start, "index", 1, "Start Index")
	flag.IntVar(&align, "a", 4, "Fill sequence index to ... digits: 0001 is 4")
	flag.StringVar(&prefix, "prefix", "eaglc", "prefix of sequence")
}

func main()  {
	flag.Parse()

	dic, err := ioutil.ReadDir(input)
	if err != nil {
		fmt.Errorf("read directory error: %v", err)
		return
	}


	index := start
	for i := range dic {
		suffix := path.Ext(dic[i].Name())
		filepath := path.Join(input, dic[i].Name())
		newfile := path.Join(input, newName(prefix, align, index, suffix))
		os.Rename(filepath, newfile)

		index++
	}
}


func newName(prefix string, alig, index int, suffix string) string {
	return fmt.Sprintf("%s%0*d%s", prefix, alig, index, suffix)
}

