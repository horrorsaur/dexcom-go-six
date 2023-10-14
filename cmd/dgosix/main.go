package main

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

// https://btprodspecificationrefs.blob.core.windows.net/assigned-numbers/Assigned%20Number%20Types/Assigned_Numbers.pdf
const DEXCOM_SERVICE_UUID uint16 = 0xFEBC       // Dexcom, Inc - Member Service UUID
const DEXCOM_COMPANY_IDENTIFIER uint16 = 0x00D0 // Dexcom, Inc - Company Identifier

var adapter = bluetooth.DefaultAdapter

func main() {
	// Enable BLE interface
	err := adapter.Enable()
	if err != nil {
		panic(err)
	}

	// Start scanning
	err = scanForDexcom(adapter)
	if err != nil {
		fmt.Println("hit timeout! did not find a dexcom device", err)
	}
}

func scanForDexcom(adapter *bluetooth.Adapter) error {
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if device.LocalName() == "DexcomSA" {
			adapter.StopScan()
			fmt.Printf("found device: %v, mac: %v", device.LocalName(), device.Address.MAC)
		}
	})

	if err != nil {
		panic("failed to scan: " + err.Error())
	}

	return nil
}
