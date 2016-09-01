package main

import (
	"fmt"
	"math/rand"

	"github.com/fogleman/bob"
	. "github.com/fogleman/pt/pt"
)

func main() {
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

	IterativeRender("out.png", 1, &scene, &camera, NewDirectSampler(), 1920, 1080, 4)

	sampler := NewSampler(4, 4)
	sampler.LightMode = LightModeAll
	sampler.SpecularMode = SpecularModeFirst
	IterativeRender("out%03d.png", 10000, &scene, &camera, sampler, 1920, 1080, -1)
}

func main2() {
	scene := Scene{}

	robot := bob.NewRobot()
	robot.PointHead(V(2, 10, 3))
	// robot.Sleep()
	mesh := robot.CreateMesh()
	scene.Add(mesh)

	material := DiffuseMaterial(White)
	scene.Add(NewPlane(V(0, 0, 0), V(0, 0, 1), material))

	light := LightMaterial(White, 50)
	scene.Add(NewSphere(V(5, 0, 10), 1, light))
	scene.Add(NewSphere(V(0, 10, 5), 1, light))

	camera := LookAt(V(2, 10, 3), V(0, 0, 3), V(0, 0, 1), 40)

	IterativeRender("out.png", 1, &scene, &camera, NewDirectSampler(), 1024, 1024, 4)

	// sampler := NewSampler(4, 4)
	// sampler.LightMode = LightModeAll
	// sampler.SpecularMode = SpecularModeFirst
	// IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
