package internal

import (
	"time"
	"udp_sim/core"
	"udp_sim/handler"
	"udp_sim/transport"
)

type Car struct {
	ID         string
	Pose       core.Pose
	Velocity   core.Velocity
	Controller *handler.CarController
	Sender     transport.Sender
}

func NewCar(id string, pose core.Pose, velocity core.Velocity, sender transport.Sender) *Car {
	controller := handler.NewCarController(id, pose)
	return &Car{
		ID:         id,
		Pose:       pose,
		Velocity:   velocity,
		Controller: controller,
		Sender:     sender,
	}
}

func (c *Car) Tick() {
	// Move
	c.Pose.X += c.Velocity.Vx
	c.Pose.Y += c.Velocity.Vy
	c.Controller.UpdateSelfPose(c.Pose)

	// Broadcast
	msg := core.PoseMessage{
		CarID:     c.ID,
		Position:  c.Pose,
		Velocity:  c.Velocity,
		Timestamp: time.Now(),
	}
	c.Sender.Send(msg)
}
