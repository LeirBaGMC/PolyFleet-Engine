package transport

import (
	"fmt"
)

type AutonomousVehicle interface {
	Start()
	Stop()
	ReportStatus() string
	AssignMission(mission string)
}

type GPS struct {
	Location string
}

func (g *GPS) UpdateLocation(loc string) {
	g.Location = loc
}

type Battery struct {
	Level int
}

func (b *Battery) ConsumeEnergy(amount int) {
	b.Level -= amount
	if b.Level < 0 {
		b.Level = 0
	}
}

type MissionData struct {
	CurrentMission string
}

func (m *MissionData) AssignMission(mission string) {
	m.CurrentMission = mission
}

type AutonomousBus struct {
	GPS
	Battery
	MissionData
	Route string
}

func (b *AutonomousBus) Start() {
	b.ConsumeEnergy(10)
	b.UpdateLocation("Starting route: " + b.Route)
	println("AutonomousBus starting on route:", b.Route)
}

func (b *AutonomousBus) Stop() {
	b.UpdateLocation("Stopped at depot")
	println("AutonomousBus stopped.")
}

func (b *AutonomousBus) ReportStatus() string {
	return "AutonomousBus - Mission: " + b.CurrentMission + " | Location: " + b.Location + " | Battery: " + fmt.Sprintf("%d%%", b.Level)
}

type DeliveryDrone struct {
	GPS
	Battery
	MissionData
	Package string
}

func (d *DeliveryDrone) Start() {
	d.ConsumeEnergy(5)
	d.UpdateLocation("Taking off for delivery")
	println("DeliveryDrone starting delivery of:", d.Package)
}

func (d *DeliveryDrone) Stop() {
	d.UpdateLocation("Landed at destination")
	println("DeliveryDrone stopped.")
}

func (d *DeliveryDrone) ReportStatus() string {
	return "DeliveryDrone - Mission: " + d.CurrentMission + " | Location: " + d.Location + " | Battery: " + fmt.Sprintf("%d%%", d.Level)
}
