package handler

import (
	"fmt"
	"sync"
	"time"
	"udp_sim/core"
	"udp_sim/util"
)

type LocalCarView struct {
	PositionLocal core.Pose
	VelocityLocal core.Velocity
	RelativeAngle float64
	LastUpdated   time.Time
}

type CarController struct {
	SelfID     string
	SelfPose   core.Pose
	Directory  map[string]LocalCarView
	DirectoryM sync.Mutex
}

func NewCarController(id string, pose core.Pose) *CarController {
	return &CarController{
		SelfID:    id,
		SelfPose:  pose,
		Directory: make(map[string]LocalCarView),
	}
}

func (c *CarController) UpdateSelfPose(pose core.Pose) {
	c.SelfPose = pose
}

func (c *CarController) Handle(msg core.PoseMessage) {
	if msg.CarID == c.SelfID {
		return
	}

	xLocal, yLocal := util.TransformToLocal(
		c.SelfPose.X, c.SelfPose.Y, c.SelfPose.Theta,
		msg.Position.X, msg.Position.Y,
	)

	vxLocal, vyLocal := util.TransformVelocityToLocal(
		c.SelfPose.Theta,
		msg.Velocity.Vx, msg.Velocity.Vy,
	)

	angle := util.ComputeRelativeBearing(xLocal, yLocal)

	c.DirectoryM.Lock()
	c.Directory[msg.CarID] = LocalCarView{
		PositionLocal: core.Pose{X: xLocal, Y: yLocal, Theta: msg.Position.Theta},
		VelocityLocal: core.Velocity{Vx: vxLocal, Vy: vyLocal},
		RelativeAngle: angle,
		LastUpdated:   time.Now(),
	}
	c.DirectoryM.Unlock()

	fmt.Printf("Updated %s → Pos(%.2f, %.2f), Vel(%.2f, %.2f), Angle=%.2f°\n",
		msg.CarID, xLocal, yLocal, vxLocal, vyLocal, angle)
}
