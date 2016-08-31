package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}

	robot := NewRobot()
	mesh := robot.CreateMesh()
	scene.Add(mesh)

	material := DiffuseMaterial(White)
	scene.Add(NewPlane(V(0, 0, 0), V(0, 0, 1), material))

	light := LightMaterial(White, 80)
	scene.Add(NewSphere(V(0, 0, 10), 1, light))
	scene.Add(NewSphere(V(0, 10, 10), 1, light))

	camera := LookAt(V(-4, 8, 4), V(0, 0, 3), V(0, 0, 1), 50)
	sampler := NewSampler(4, 8)
	sampler.LightMode = LightModeAll
	sampler.SpecularMode = SpecularModeFirst
	IterativeRender("out%03d.png", 10000, &scene, &camera, sampler, 1024*2, 1024*2, -1)
}

type Robot struct {
}

func NewRobot() *Robot {
	return &Robot{}
}

func (robot *Robot) CreateMesh() Shape {
	whitePlastic := GlossyMaterial(White, 1.5, Radians(10))
	blackPlastic := GlossyMaterial(Black, 1.5, Radians(10))
	yellow := GlossyMaterial(HexColor(0xF7E967), 1.3, Radians(10))
	green := GlossyMaterial(HexColor(0xA9CF54), 1.3, Radians(10))
	darkGray := GlossyMaterial(HexColor(0x374140), 1.5, Radians(10))

	code, _ := LoadSTL("stl/code.stl", blackPlastic)
	antenna, _ := LoadSTL("stl/antenna.stl", darkGray)
	head, _ := LoadSTL("stl/head.stl", green)
	eye, _ := LoadSTL("stl/eye.stl", whitePlastic)
	pupil, _ := LoadSTL("stl/pupil.stl", blackPlastic)
	neck, _ := LoadSTL("stl/neck.stl", whitePlastic)
	body, _ := LoadSTL("stl/body.stl", yellow)
	arm, _ := LoadSTL("stl/arm.stl", whitePlastic)
	wheel, _ := LoadSTL("stl/wheel.stl", whitePlastic)

	mesh := NewMesh(nil)
	mesh.Add(wheel)
	mesh.Add(body.Transformed(Translate(V(0, 0, 0.75))))
	mesh.Add(head.Transformed(Translate(V(0, 0, 4))))
	mesh.Add(neck.Transformed(Translate(V(0, 0, 3.75))))
	mesh.Add(antenna.Transformed(Translate(V(0, 0, 6))))
	mesh.Add(code.Transformed(Translate(V(0, 1-0.05, 1.5))))
	mesh.Add(arm.Transformed(Translate(V(0, 0, 0)).Rotate(V(1, 0, 0), Radians(120)).Translate(V(-1.125, 0, 3))))
	mesh.Add(arm.Transformed(Translate(V(0, 0, 0)).Rotate(V(1, 0, 0), Radians(30)).Translate(V(1.125, 0, 3))))
	mesh.Add(eye.Transformed(Translate(V(-0.4, 0.6, 4.75))))
	mesh.Add(eye.Transformed(Translate(V(0.4, 0.6, 4.75))))
	mesh.Add(pupil.Transformed(Scale(V(0.75, 0.75, 0.75)).Translate(V(-0.4, 0.61+0.5, 4.75))))
	mesh.Add(pupil.Transformed(Scale(V(0.5, 0.5, 0.5)).Translate(V(0.4, 0.61+0.5, 4.75))))
	mesh.SmoothNormalsThreshold(Radians(30))
	return mesh
}
