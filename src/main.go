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
		fmt.Println("ERROR:  Create from yaml")
		fmt.Println(err)
		return
	}

	// Get results
	spack.GetOidsFromSwitches()
	// TODO process results

	for _, v := range spack.ArrSwitchSNMP {
		fmt.Printf("Switch %s results: \n\n", v.SwitchName)
		for _, v := range v.Results.Variables {
			fmt.Println(v)
		}
	}

	fmt.Println("Program stopped")

	// snmp := gosnmp.Default
	// snmp.Target = "172.30.133.159"
	// snmp.Port = 161
	// snmp.Version = gosnmp.Version2c
	// snmp.Community = "public"
	//
	// fmt.Println("Connecting")
	//
	// fmt.Println(snmp)
	//
	// err := snmp.Connect()
	//
	//	if err != nil {
	//		fmt.Println("Error making connection")
	//		fmt.Println(err)
	//		return
	//	}
	//
	// defer snmp.Conn.Close()
	//
	// result, err := snmp.Get([]string{".1.3.6.1.2.1.2.2.1.7.3"})
	//
	//	if err != nil {
	//		fmt.Println("Error gettings")
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	for _, v := range result.Variables {
	//		fmt.Printf("%+v\n", v)
	//		if v.Type != gosnmp.BitString {
	//			fmt.Printf("Numero %s\n", v.Name)
	//			fmt.Println(v.Value)
	//		}
	//	}
}

// r.ProcessResults() <-- not this simple
