package main

import (
	"flag"
	"fmt"
	"time"
	"udp_sim/core"
	"udp_sim/handler"
	"udp_sim/internal"
	"udp_sim/transport"
)

func main() {
	// CLI args
	id := flag.String("id", "CarX", "Unique car ID")
	x := flag.Float64("x", 0, "Initial X position")
	y := flag.Float64("y", 0, "Initial Y position")
	theta := flag.Float64("theta", 0, "Initial orientation in radians")
	vx := flag.Float64("vx", 1, "Velocity in X")
	vy := flag.Float64("vy", 0, "Velocity in Y")
	port := flag.Int("port", 9999, "UDP multicast port")
	group := flag.String("group", "224.0.0.1", "Multicast group IP")
	flag.Parse()
	pose := core.Pose{X: *x, Y: *y, Theta: *theta}
	velocity := core.Velocity{Vx: *vx, Vy: *vy}
	encoder := &core.JSONEncoder{}
	controller := handler.NewCarController(*id, pose)
	receiver := &transport.UDPReceiver{
		Port:    *port,
		GroupIP: *group,
		Encoder: encoder,
		Handler: controller,
	}
	sender := &transport.UDPSender{
		Addr:    fmt.Sprintf("%s:%d", *group, *port),
		Encoder: encoder,
	}
	car := internal.NewCar(*id, pose, velocity, sender)
	car.Controller = controller
	go receiver.Start()
	for {
		car.Tick()
		time.Sleep(1 * time.Second)
	}
}
