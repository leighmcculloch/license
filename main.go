package main // import "4d63.com/license"

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"4d63.com/licenses"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "License provides license text for %d open source lisence.\n\n", licenses.Count())
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "  license <license>\n\n")
		fmt.Fprintf(os.Stderr, "Example:\n\n")
		fmt.Fprintf(os.Stderr, "  license mit > LICENSE\n\n")
		fmt.Fprintf(os.Stderr, "Licenses:\n\n")

		names := licenses.Names()
		sort.Strings(names)

		const cols = 4
		for i := 0; i < len(names); i += cols {
			fmt.Fprintf(os.Stderr, "  ")
			for k := 0; k < cols && i+k < len(names); k++ {
				name := names[i+k]
				fmt.Fprintf(os.Stderr, "%-20s", name)
			}
			fmt.Fprintf(os.Stderr, "\n")
		}
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "A single license only must be specified.\n")
		return
	}

	name := strings.ToLower(args[0])

	text, ok := licenses.Text(name)
	if !ok {
		fmt.Fprintf(os.Stderr, "License %s is not known.\n", name)
		alts := []string{}
		for _, n := range licenses.Names() {
			if strings.HasPrefix(n, name) {
				alts = append(alts, n)
			}
		}
		if len(alts) > 0 {
			sort.Strings(alts)
			fmt.Fprintf(os.Stderr, "Did you mean: %s?\n", strings.Join(alts, ", "))
		}
		return
	}

	os.Stdout.Write(text)
}
