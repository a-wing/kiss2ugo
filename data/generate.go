// +build ignore

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shurcooL/vfsgen"
)

func main() {
	os.Chdir("data")
	err := vfsgen.Generate(http.Dir("../docs"), vfsgen.Options{
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Docs",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
