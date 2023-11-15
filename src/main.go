package main

import (
	"SwitchSNMP/servsnmp"
	"flag"
	"fmt"
)

// File variable to hold flag value for input file
var File string

// Init flags
func init() {
	flag.StringVar(&File, "f", "servers.yaml", "Input the name of the file, when it is located in the same directory as the program itself. Default value: 'servers.yaml'.")
	// TODO flag for single switch lookup
	flag.Parse()
}

func main() {
	// Read given yaml file into struct Switches, which holds an array of our switches
	spack, err := servsnmp.CreateFromYaml(File)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO get results

	r := spack.GetOidsFromSwitches()

	// r := spack.GetResults()

	// TODO process results

	// r.ProcessResults()
}
