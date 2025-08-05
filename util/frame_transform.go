package util

import (
	"math"
)

func TransformToLocal(xSelf, ySelf, theta float64, xOther, yOther float64) (float64, float64) {
	dx := xOther - xSelf
	dy := yOther - ySelf
	cosT := math.Cos(theta)
	sinT := math.Sin(theta)
	xLocal := dx*cosT + dy*sinT
	yLocal := -dx*sinT + dy*cosT
	return xLocal, yLocal
}

func TransformVelocityToLocal(theta float64, vx, vy float64) (float64, float64) {
	cosT := math.Cos(theta)
	sinT := math.Sin(theta)
	vxLocal := vx*cosT + vy*sinT
	vyLocal := -vx*sinT + vy*cosT
	return vxLocal, vyLocal
}

func ComputeRelativeBearing(x, y float64) float64 {
	angle := math.Atan2(y, x) * 180 / math.Pi
	for angle > 180 {
		angle -= 360
	}
	for angle < -180 {
		angle += 360
	}
	return angle
}
