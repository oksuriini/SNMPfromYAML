# SNMPfromYAML Project
Simple project to get information about switches/devices that have SNMP server capability.
This is primarily meant for scanning ethernet interfaces/ports on switches, but technically can work on any other property.



## Prerequisites:
Make sure you have golang installed:
https://go.dev/dl/

## Usage:
Clone repo:
```git clone https://github.com/oksuriini/SNMPfromYAML.git```

cd into repo:
```cd SNMPfromYAML/```

Run program from sources:
```go run ./src/```

Alternatively run program from sources with flags:
```go run ./src/ -f "testitiedosto.yaml"```

...or with single mode: ( NOT IMPLEMENTED)
```go run ./src/ -f -m -s "172.30.133.122" -c "public" -o "1.3.2.1.6.1.2.5.12"```

Change "servers.yaml" file properties to make sure it works in your environment.

## Features:

Parses YAML file for switches/devices, and then parses OIDs for which to check these devices for.
Currently works with both multiple devices and OIDs through which to scan.


