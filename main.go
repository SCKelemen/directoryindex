package main

import (
	"flag"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	filepath.Walk(root, visit)
}

func visit(path string, f os.FileInfo, err error) error {

}
