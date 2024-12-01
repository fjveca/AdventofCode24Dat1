package main

import (
	"fmt"
	"buffio"
	"errors"
	"os"
	"strings"
)

func main() {

  input, error := os.OpenFile("input", os.O_RDONLY)
  if error != nil {
	log.fatal("no valid input file found")
  }
  

}
