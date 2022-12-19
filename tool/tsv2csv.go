package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	bakf, err := os.Create(flag.Arg(0) + ".bak")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(bakf, f); err != nil {
		log.Fatal(err)
	}
	bakf.Close()

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		log.Fatal(err)
	}

	w := &bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		tokens := strings.Fields(sc.Text())
		w.WriteString(strings.Join(tokens, ",") + "\n")
	}
	f.Close()

	os.Remove(flag.Arg(0))

	f2, err := os.Create(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(f2, w); err != nil {
		log.Fatal(err)
	}
	f2.Close()
}
