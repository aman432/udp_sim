package core

import "time"

type Pose struct {
    X     float64
    Y     float64
    Theta float64
}

type Velocity struct {
    Vx float64
    Vy float64
}

type PoseMessage struct {
    CarID     string
    Position  Pose
    Velocity  Velocity
    Timestamp time.Time
}
