package main

import (
	"fmt"
	"github.com/speedata/decorate"
	"github.com/speedata/optionparser"
	"io/ioutil"
	"log"
	"os"
)

func listinput() {
	input := decorate.InputFilters()
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}
	os.Exit(0)
}

func listoutput() {
	output := decorate.OutputFilters()
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
	os.Exit(0)
}

func main() {
	inputfilter := "text"
	outputfilter := "text"
	op := optionparser.NewOptionParser()
	op.Banner = "Usage: gohigh [options] inputfile [outputfile]"
	op.On("--list-input", "List input filter", listinput)
	op.On("--list-output", "List output filter", listoutput)
	op.On("-i", "--inputfilter FILTER", "Use input filter", &inputfilter)
	op.On("-o", "--outputfilter FILTER", "Use output filter", &outputfilter)
	op.Parse()

	if len(op.Extra) < 1 {
		fmt.Println("Need at least one argument.")
		op.Help()
		os.Exit(-1)
	}
	inputfile := op.Extra[0]

	ret, err := decorate.HighlightFile(inputfile, inputfilter, outputfilter)
	if err != nil {
		log.Fatal(err)
	}
	if len(op.Extra) < 2 {
		fmt.Println(ret)
	} else {
		ioutil.WriteFile(op.Extra[1], []byte(ret), 0644)
	}
}
