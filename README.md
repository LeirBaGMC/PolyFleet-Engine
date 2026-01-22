#  NeoQuito: Autonomous Fleet Simulation
> **Golang OOP Showcase** | Interface-Based Polymorphism | Composition over Inheritance

![Go](https://img.shields.io/badge/Language-Go_1.25+-00ADD8?logo=go&logoColor=white)
![Architecture](https://img.shields.io/badge/Pattern-Composition_&_Interfaces-purple)
![Status](https://img.shields.io/badge/Status-Educational_Prototype-success)

##  Overview
**NeoQuito** is a backend simulation engine designed to manage a heterogeneous fleet of autonomous vehicles (Buses and Drones). Unlike traditional OOP languages, this project leverages Go's type system to demonstrate **composition (struct embedding)** and **interface-based polymorphism** to orchestrate complex behaviors without class inheritance.

## Technical Architecture

The system is modularized within the `transport` package, enforcing strict separation of concerns:

### 1. Core Interfaces
The system relies on the `AutonomousVehicle` interface to decouple implementation from execution.
```go
type AutonomousVehicle interface {
    Start()
    Stop()
    AssignMission(mission string)
    ReportStatus() string
}
```
### 2. Composition & Embedding
Instead of inheritance trees, shared behaviors are injected via struct embedding, adhering to the "Has-a" vs "Is-a" principle:

* **GPS Unit:** Manages coordinate tracking (`Location`).
* **Battery System:** Handles energy consumption logic (`ConsumeEnergy`).
* **Mission Data:** State management for active tasks.

### 3. Polymorphic Control Center
The `ControlCenter` acts as the fleet orchestrator. It manages a slice of `[]AutonomousVehicle`, allowing it to batch-process commands (Start/Stop/Report) across different vehicle types (Bus vs. Drone) unaware of their underlying concrete implementation.

---

##  Quick Start

### Prerequisites
* Go 1.21 or higher

### Execution
1. **Clone the repository:**
   ```bash
   git clone [https://github.com/LeirBaGMC/NeoQuito-Autonomous-Fleet.git](https://github.com/LeirBaGMC/NeoQuito-Autonomous-Fleet.git)
   cd NeoQuito-Autonomous-Fleet

2. **Run the simulation:**
   ```bash
   go run main.go

NeoQuito-Autonomous-Fleet/
├── go.mod                  # Module definition
├── main.go                 # Simulation entry point
├── README.md               # Documentation
└── transport/              # Core Logic Package
    ├── vehicle.go          # Structs (Bus, Drone) & Embeds
    └── control.go          # ControlCenter Logic

## Key Learning Outcomes
This project demonstrates proficiency in:

Go Interfaces: Defining behavior contracts.

Struct Embedding: Reusing logic without inheritance.

Slice Management: Handling dynamic collections of interfaces.

Package Visibility: Managing exported/unexported identifiers in Go.

Developed by Gabriel Minda & Diego Ruiz as a Software Architecture POC.
