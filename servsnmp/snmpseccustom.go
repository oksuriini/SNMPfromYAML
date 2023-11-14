package servsnmp

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gosnmp/gosnmp"
	"gopkg.in/yaml.v3"
)

// SwitchSNMP struct to hold information of each switch
type SwitchSNMP struct {
	SwitchName   string        `yaml:"name"`
	IpAddress    string        `yaml:"ip_address"`
	CommunityStr string        `yaml:"community_string"`
	PortCount    int           `yaml:"port_count"`
	snmpObj      gosnmp.GoSNMP `yaml:"snmp_obj"`
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
	return s, nil
}

// For debugging purposes
func (s SnmpPack) ListSwitches() {
	fmt.Println("Printing switches found in yaml file:")
	for _, v := range s.ArrSwitchSNMP {
		fmt.Printf("Switch: %s \nIP-Address: %s \nCommunityString: %s \nPort Count: %d \n\n", v.SwitchName, v.IpAddress, v.CommunityStr, v.PortCount)
	}
}

// TODO Method executes given Oids to switches in SnmpPack
func (s SnmpPack) GetOidsFromSwitches() {
}
