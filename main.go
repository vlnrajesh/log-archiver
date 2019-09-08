package main

import (
	"fmt"

	"github.com/vlnrajesh/golang-examples/string_utils"
	"github.com/vlnrajesh/log-archiver/utils"
)

func main() {
	fmt.Println("First commit")
	fmt.Println(string_utils.Reverse("Hello"))
	utils.CmdLineParser()
}
