package main

import (
	"SwitchSNMP/servsnmp"
	"flag"
	"fmt"

	"github.com/gosnmp/gosnmp"
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

	// Get results

	spack.GetOidsFromSwitches()

	// TODO process results

	for _, v := range spack.ArrSwitchSNMP {
		fmt.Printf("Switch %s results: \n\n", v.SwitchName)
		for i, j := range v.Results.Variables {
			fmt.Printf("%d: oid: %s ", i, j.Name)

			switch j.Type {
			case gosnmp.OctetString:
				bytes := j.Value.([]byte)
				fmt.Printf("string: %s\n", string(bytes))
			default:
				fmt.Printf("number: %d\n", gosnmp.ToBigInt(j.Value))
			}
		}

		fmt.Println("Program stopped")
	}
}

// r.ProcessResults() <-- not this simple
