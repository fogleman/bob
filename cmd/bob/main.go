package main

import (
	"fmt"
	"math/rand"

	"github.com/fogleman/bob"
	. "github.com/fogleman/pt/pt"
)

func main2() {
	scene := Scene{}

	points := PoissonDisc(-120, -100, 120, 100, 6*1, 32)
	fmt.Println(len(points))
	robots := NewMesh(nil)
	for i, point := range points {
		robot := bob.NewRobot()
		robot.Number = i
		robot.Random()
		robot.Position = point
		robot.PointBody(V(0, 100, 20))
		robot.Rotation += rand.Float64()*90 - 45
		if rand.Float64() < 0.15 {
			robot.Sleep()
		} else {
			robot.PointHead(V(0, 100, 20))
		}
		mesh := robot.CreateMesh()
		robots.Add(mesh)
	}
	robots.SmoothNormalsThreshold(Radians(30))
	scene.Add(robots)
	fmt.Println(len(robots.Triangles))

	material := DiffuseMaterial(White)
	scene.Add(NewPlane(V(0, 0, 0), V(0, 0, 1), material))

	light := LightMaterial(White, 80)
	c := V(0, 50, 50)
	scene.Add(NewCube(V(c.X-5, c.Y-5, c.Z-0.1), V(c.X+5, c.Y+5, c.Z+0.1), light))

	camera := LookAt(V(0, 100, 20), V(0, 65, 0), V(0, 0, 1), 50)

	IterativeRender("out.png", 1, &scene, &camera, NewDirectSampler(), 2560, 1440, 9)

	sampler := NewSampler(4, 4)
	sampler.LightMode = LightModeAll
	sampler.SpecularMode = SpecularModeFirst
	IterativeRender("out%03d.png", 10000, &scene, &camera, sampler, 2560, 1440, -1)
}

func main() {
	scene := Scene{}

	robots := NewMesh(nil)
	for x := -1; x <= 1; x++ {
		robot := bob.NewRobot()
		robot.Position = V(float64(x)*4, 0, 0)
		robot.PointBody(V(2, 10, 3))
		if x == 0 {
			robot.LeftArm = 1
		}
		if x == -1 {
			robot.Neck = 0.75
			robot.Antenna = 0.1
		}
		robot.Number = -x + 2
		robot.PointHead(V(2, 10, 3))
		robots.Add(robot.CreateMesh())
	}
	robots.SmoothNormalsThreshold(Radians(30))
	scene.Add(robots)

	material := DiffuseMaterial(White)
	scene.Add(NewPlane(V(0, 0, 0), V(0, 0, 1), material))

	light := LightMaterial(White, 100)
	scene.Add(NewSphere(V(10, 0, 10), 1, light))
	scene.Add(NewSphere(V(0, 10, 10), 1, light))

	camera := LookAt(V(2, 10, 3), V(0, 0, 3), V(0, 0, 1), 50)

	IterativeRender("out.png", 1, &scene, &camera, NewDirectSampler(), 2560, 1440, 1)

	sampler := NewSampler(4, 4)
	sampler.LightMode = LightModeAll
	sampler.SpecularMode = SpecularModeFirst
	IterativeRender("out%03d.png", 10000, &scene, &camera, sampler, 2560, 1440, -1)
}
