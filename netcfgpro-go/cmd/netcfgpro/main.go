package main

import (
	"log"
	"netcfgpro-go/internal/generators"
	"netcfgpro-go/internal/models"
	"os"
)

func main() {
	// Create a sample DeviceConfig for demonstration.

	iface1 := models.NewInterface()
	iface1.Name = "GigabitEthernet0/1"
	iface1.Type = models.InterfaceTypeGigabit
	iface1.Description = "Uplink to a friendly server"
	iface1.IPAddress = strPtr("192.168.1.1")
	iface1.SubnetMask = strPtr("255.255.255.0")
	iface1.MTU = 9000

	iface2 := models.NewInterface()
	iface2.Name = "GigabitEthernet0/2"
	iface2.Type = models.InterfaceTypeGigabit
	iface2.Description = "Access Port for VLAN 10"
	iface2.SwitchportMode = models.SwitchportModeAccess
	iface2.AccessVLAN = intPtr(10)

	// Add a disabled interface to test the IsEnabled() helper method.
	disabled := false
	iface3 := models.NewInterface()
	iface3.Name = "GigabitEthernet0/3"
	iface3.Type = models.InterfaceTypeGigabit
	iface3.Description = "This port should be disabled"
	iface3.Enabled = &disabled

	sampleConfig := &models.DeviceConfig{
		Hostname: "CORE-SW-01",
		Vendor:   models.VendorCiscoIOS,
		Interfaces: []models.Interface{
			*iface1,
			*iface2,
			*iface3,
		},
		VLANs: []models.VLAN{
			{
				ID:   10,
				Name: "SERVERS",
			},
		},
	}

	// Generate the configuration and write it to standard output.
	err := generators.Generate(sampleConfig, os.Stdout)
	if err != nil {
		log.Fatalf("Error generating configuration: %v", err)
	}
}

// Helper functions to create pointers for optional fields.
func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
