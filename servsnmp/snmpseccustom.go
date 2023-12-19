package servsnmp

// REMINDER !!!!
// This was built to scan switch interfaces (ports if you may) so while other OID objects might work, the structures
// and functions are not assured to work nicely (or at all) with them.

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gosnmp/gosnmp"
	"gopkg.in/yaml.v3"
)

// SwitchSNMP struct to hold information of each switch
type SwitchSNMP struct {
	SwitchName   string `yaml:"name"`
	IpAddress    string `yaml:"ip_address"`
	CommunityStr string `yaml:"community_string"`
	PortCount    int    `yaml:"port_count"`

	// Holds GoSNMP struct used to make Get(oids) to switches
	snmpObj gosnmp.GoSNMP

	// Holds Get(oids) results
	Results gosnmp.SnmpPacket
}

// OidStruct struct to hold information of each oid
type OidStruct struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Oid         string `yaml:"oid"`
	Iterable    bool   `yaml:"iterable"`
}

// SnmpPack wraps OidStruct and SwitchSNMP structs within itself.
// Purpose is to enable nested yaml files as in the given template file
type SnmpPack struct {
	ArrSwitchSNMP []SwitchSNMP `yaml:"switches"`
	ArrOids       []OidStruct  `yaml:"oids"`
}

// This works as a wrapper to create our snmpObj, which allows us to make connections to snmp servers (w/e it is)
func NewSwitchSNMP(portCount int, communityStr string, ipAddress string, switchName string) SwitchSNMP {
	swSnmp := SwitchSNMP{
		PortCount:    portCount,
		CommunityStr: communityStr,
		IpAddress:    ipAddress,
		SwitchName:   switchName,
		snmpObj:      createsnmpObj(communityStr, ipAddress),
	}

	return swSnmp
}

func createsnmpObj(communityStr string, ipAddress string) gosnmp.GoSNMP {
	snmpObj := gosnmp.Default

	snmpObj.Target = ipAddress
	snmpObj.Community = communityStr
	snmpObj.Version = gosnmp.Version2c
	snmpObj.Port = 161

	// fmt.Println(snmpObj)

	return *snmpObj
}

// Function to create SnmpPack struct based on given yaml file
func CreateFromYaml(fileName string) (SnmpPack, error) {
	// data, err := os.Open(fmt.Sprintf("./%s", File))
	data, err := os.Open(fmt.Sprintf("./src/%s", fileName))
	if err != nil {
		return SnmpPack{}, err
	}
	defer data.Close()

	size, err := data.Stat()
	if err != nil {
		return SnmpPack{}, err
	}

	in := make([]byte, size.Size())

	_, err = bufio.NewReader(data).Read(in)
	if err != nil {
		return SnmpPack{}, err
	}

	s := SnmpPack{}

	err = yaml.Unmarshal(in, &s)
	if err != nil {
		return SnmpPack{}, err
	}
	for i, v := range s.ArrSwitchSNMP {
		s.ArrSwitchSNMP[i].snmpObj = createsnmpObj(v.CommunityStr, v.IpAddress)
	}
	return s, nil
}

// Get oids get all the oids in the structure and returns them in array of strings
func (s SnmpPack) getOids(iterator int) []string {
	var oids []string
	oidstrc := s.ArrOids
	for _, v := range oidstrc {
		// If oid is iterable then iterate iterator value at the end of the oid
		// As this is purely built for checking switch interface status this isn't very dynamic
		if v.Iterable {
			for i := 1; i <= iterator; i++ {
				oids = append(oids, fmt.Sprintf("%s.%s", v.Oid, fmt.Sprint(i)))
			}
		} else {
			oids = append(oids, v.Oid)
		}
	}
	return oids
}

// For debugging purposes <- redacted
// Not anymore uses CreateFileFromResults to save results
func (s SnmpPack) ListSwitches() {
	fmt.Println("Printing switches found in yaml file:")
	for _, v := range s.ArrSwitchSNMP {
		result := fmt.Sprintf("Switch: %s \nIP-Address: %s \nCommunityString: %s \nPort Count: %d \n\n", v.SwitchName, v.IpAddress, v.CommunityStr, v.PortCount)
		createFileFromResults(result)
	}
}

func (s *SwitchSNMP) setResult(result gosnmp.SnmpPacket) {
	s.Results = result
}

// Method executes given Oids to switches in SnmpPack, and returns results in gosnmp.SnmpPacket structure
// for further processing.
func (s *SnmpPack) GetOidsFromSwitches() {
	for i, v := range s.ArrSwitchSNMP {
		if err := v.snmpObj.Connect(); err != nil {
			fmt.Println("ERROR in connecting")
			fmt.Println(err)
			continue
		}
		result, err := v.snmpObj.Get(s.getOids(v.PortCount))
		if err != nil {
			fmt.Println("ERROR GETTING RESULTS")
			continue
		}

		s.ArrSwitchSNMP[i].setResult(*result)
	}
}

func createFileFromResults(s string) {
	os.WriteFile("results", []byte(s), 0666)
}
