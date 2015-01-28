package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var _ = log.Print

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	flag.Usage = func() {
		fmt.Printf("Usage: jsonfilter [options] selector-string\n\nOptions:\n")
		flag.PrintDefaults()
	}

	asJson := flag.Bool("json", false, "display output as json")
	pretty := flag.Bool("pretty", false, "display pretty json (sets -json=true)")
	asValues := flag.Bool("values", false, "display each value of selected structure on its own line")
	asKeys := flag.Bool("keys", false, "display each key or index of selected structure on its own line")
	verbose := flag.Bool("v", false, "display errors")
	debug := flag.Bool("vv", false, "display errors and info")

	flag.Parse()

	if *pretty {
		asJson = pretty
	}
	if *debug {
		verbose = debug
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		if *debug {
			log.Println("stdin content detected")
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

	selector := ""
	for i := 0; i < flag.NArg(); i++ {
		selector += flag.Arg(i)
	}

	app := App{
		AsJson:   *asJson,
		Pretty:   *pretty,
		AsValues: *asValues,
		AsKeys:   *asKeys,
		Debug:    *debug,
		Input:    os.Stdin,
	}

	str, err := app.run(selector)
	if err != nil {
		if *verbose {
			log.Println(err)
		}
		os.Exit(1)
	}

	fmt.Print(str)
}
