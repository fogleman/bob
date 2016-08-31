package main

import (
	"fmt"

	"github.com/fogleman/bob"
	. "github.com/fogleman/pt/pt"
)

func main() {
	scene := Scene{}

	points := PoissonDisc(-100, -100, 100, 100, 5, 32)
	fmt.Println(len(points))
	robots := NewMesh(nil)
	for _, point := range points {
		robot := bob.NewRobot()
		robot.Random()
		mesh := robot.CreateMesh()
		mesh.Transform(Translate(point))
		robots.Add(mesh)
	}
	scene.Add(robots)
	fmt.Println(len(robots.Triangles))

	material := DiffuseMaterial(White)
	scene.Add(NewPlane(V(0, 0, 0), V(0, 0, 1), material))

	light := LightMaterial(White, 30)
	scene.Add(NewSphere(V(0, 0, 100), 10, light))
	scene.Add(NewSphere(V(100, 10, 50), 10, light))

	camera := LookAt(V(100, 0, 20), V(60, 0, 0), V(0, 0, 1), 50)

	IterativeRender("out.png", 1, &scene, &camera, NewDirectSampler(), 1920/2, 1080/2, 4)

	sampler := NewSampler(4, 4)
	sampler.LightMode = LightModeAll
	sampler.SpecularMode = SpecularModeFirst
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
