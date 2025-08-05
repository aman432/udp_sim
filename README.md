# UDP Car Simulation

## Introduction

UDP Car Simulation is a Go-based project for simulating multiple cars communicating their positions and velocities over UDP. Each car maintains a local view of other cars, transforming their global positions and velocities into its own local frame. 

## Features

- Simulate multiple cars with unique IDs and initial states.
- Real-time position and velocity updates over UDP.
- Local frame transformation for relative position, velocity, and bearing.
- Thread-safe directory of observed cars.

## Prerequisites

- **Go 1.20+** (or latest stable version)
- macOS, Linux, or Windows

## Installation

### 1. Download and Install Go

#### macOS

```sh
brew install go
```

#### Linux

```sh
sudo apt-get install golang-go
```

#### Windows
1. Download the Go installer from the [official Go website](https://golang.org/dl/).
2. Run the installer and follow the instructions.

## Verify Go Installation
To verify that Go is installed correctly, run the following command in your terminal or command prompt:

```sh
go version
```
If Go is installed correctly, you should see the version number printed in the terminal.

### 2. Clone the Repository

```sh
git clone https://github.com/amanbharadwajrzp/udp_sim.git
cd udp_sim
```
### 3. Download Dependencies
```sh
go mod tidy
```
### 4. Build the Project

```sh
go build -o udp_sim main.go
```

### 5. Run the Simulation
To run the simulation, execute the following command:
#### Terminal 1: Car A
```sh
go run cmd/car/main.go --id=CarA --x=0 --y=0 --vx=1 --vy=0
```
#### Terminal 2: Car B
```sh
go run cmd/car/main.go --id=CarB --x=10 --y=5 --vx=-1 --vy=0
```
#### Terminal 3: Car C
```sh
go run cmd/car/main.go --id=CarC --x=5 --y=10 --vx=0 --vy=-1
```

Each car will print updates about the other car's position, velocity, and relative angle in its local frame.

