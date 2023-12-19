package main

import (
	"SwitchSNMP/servsnmp"
	"flag"
	"fmt"
)

// File variable to hold flag value for input file
var (
	File      string
	Single    bool
	Server    string
	Community string
)

// Init flags
func init() {
	flag.StringVar(&File, "f", "servers.yaml", "Input the name of the file, when it is located in the same directory as the program itself. Default value: 'servers.yaml'.")
	flag.BoolVar(&Single, "m", false, "Input whether you want to handle single server or all servers from the YAML file")
	flag.StringVar(&Server, "s", "0.0.0.0", "When using single mode, input switch/device IP address")
	flag.StringVar(&Community, "c", "public", "Enter community string that works for the switch/device")
	flag.Parse()
}

func main() {
	if !Single {
		handleMultiple()
	} else {
		handleSingle(Server, Community)
	}
	fmt.Println("Program stopped")
}

// Read given yaml file into struct Switches, which holds an array of our switches
func handleMultiple() {
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
}

func handleSingle(server string, community string) {
	fmt.Println("LoL")
}

// r.ProcessResults() <-- not this simple

// YE WHO ENTER HERE, NOTHING BUT ARCHIVAL SHIT HERE
//
//
//
//
//
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
