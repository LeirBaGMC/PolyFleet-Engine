package main

import (
	"fmt"

	"mi-sistema/transport"
)

func main() {

	cc := &transport.ControlCenter{}

	bus := &transport.AutonomousBus{
		GPS:         transport.GPS{Location: "Depot"},
		Battery:     transport.Battery{Level: 100},
		MissionData: transport.MissionData{},
		Route:       "Line 5",
	}
	drone := &transport.DeliveryDrone{
		GPS:         transport.GPS{Location: "Warehouse"},
		Battery:     transport.Battery{Level: 100},
		MissionData: transport.MissionData{},
		Package:     "Package A",
	}

	cc.RegisterVehicle(bus)
	cc.RegisterVehicle(drone)

	bus.AssignMission("Transport passengers on Line 5")
	drone.AssignMission("Deliver package to address X")

	cc.ExecuteAction(bus, "start")
	cc.ExecuteAction(drone, "start")
	cc.ExecuteAction(bus, "stop")
	cc.ExecuteAction(drone, "stop")

	report := cc.GenerateReport()
	fmt.Println("Consolidated Report:")
	fmt.Print(report)
}
