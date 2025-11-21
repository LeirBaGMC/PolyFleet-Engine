package transport

type ControlCenter struct {
	Vehicles []AutonomousVehicle
}

func (cc *ControlCenter) RegisterVehicle(v AutonomousVehicle) {
	cc.Vehicles = append(cc.Vehicles, v)
}

func (cc *ControlCenter) ExecuteAction(v AutonomousVehicle, action string) {
	switch action {
	case "start":
		v.Start()
	case "stop":
		v.Stop()
	default:
		println("Unknown action")
	}
}

func (cc *ControlCenter) GenerateReport() string {
	report := ""
	for _, v := range cc.Vehicles {
		report += "[" + v.ReportStatus() + "]\n"
	}
	return report
}
