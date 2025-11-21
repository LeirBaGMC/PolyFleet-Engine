# Sistema de Transporte Autónomo - NeoQuito

Este proyecto implementa una simulación de gestión de vehículos autónomos utilizando el lenguaje de programación Go. El diseño se centra en la Programación Orientada a Objetos (POO) aplicando interfaces, polimorfismo y composición.

## 📋 Integrantes del Grupo
* Gabriel Minda
* Diego Ruiz

## 🚀 Cómo ejecutar la simulación

1. Asegúrate de tener instalado Go (versión 1.25 o superior).
2. Ubícate en la carpeta raíz del proyecto (donde está el archivo `go.mod`).
3. Ejecuta el siguiente comando en la terminal:

```bash
go run .

O alternativamente:


go run main.go

🛠️ Diseño del Sistema
El sistema está modularizado dentro del paquete transport y se estructura de la siguiente manera:

Interfaces Usadas
AutonomousVehicle: Es la interfaz principal que define el comportamiento obligatorio para cualquier vehículo en el sistema.

Start(): Inicia el vehículo y consume energía.

Stop(): Detiene el vehículo.

AssignMission(mission string): Asigna una tarea específica.

ReportStatus() string: Devuelve el estado actual del vehículo.

Tipos Implementados (Structs)
Se han implementado dos tipos de vehículos concretos:

AutonomousBus: Simula un autobús de transporte público.

DeliveryDrone: Simula un dron de entrega de paquetes.

Composición (Modularidad)
Para evitar la repetición de código, se utilizaron structs reutilizables mediante composición (embedding):

GPS: Gestiona la ubicación actual (Location).

Battery: Gestiona el nivel de energía y el consumo (ConsumeEnergy).

MissionData: Gestiona la descripción de la misión actual.

Control Center
El ControlCenter actúa como el orquestador del sistema:

Utiliza un slice de interfaces AutonomousVehicle para almacenar diferentes tipos de vehículos.

Ejecuta acciones (start, stop) de manera polimórfica sin conocer el tipo exacto del vehículo.

Genera reportes consolidados del estado de toda la flota.


---

### 3. Estructura Final de Archivos
Solo para confirmar, tu carpeta debe verse así antes de entregar:

```text
TALLER_MINDA_RUIZ/
│
├── go.mod               <-- Define el módulo "mi-sistema"
├── main.go              <-- Tu simulación principal
├── README.md            <-- El archivo que te acabo de dar
│
└── transport/           <-- Carpeta del paquete
    ├── vehicle.go       <-- Interfaces, Structs (Bus, Drone, GPS, Battery)
    └── control.go       <-- Struct ControlCenter y sus métodos
