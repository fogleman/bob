package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/fogleman/bob"
	. "github.com/fogleman/pt/pt"
)

func createScene(t, prev float64) (*Scene, *Camera) {
	scene := Scene{}

	changed := false
	rnd := rand.New(rand.NewSource(2))

	mesh := NewMesh(nil)
	for y := 0; y < 3; y++ {
		for x := -2; x <= 2; x++ {
			robot := bob.NewRobot()
			robot.Position = V(float64(x)*4, float64(y)*-4, 0)

			if x == 0 && y == 0 {
				s := rnd.Float64()*5 + 0.5
				s = 2.5
				var o float64
				var v bob.Float64Variable

				o, v = 0.0, nil
				v.Add(s+o, s+o+1, 0, 1, bob.EaseInOutCubic)
				changed = changed || v.Changed(t, prev)
				robot.Antenna = v.Get(t) * 0.5

				o, v = 0.75, nil
				v.Add(s+o, s+o+1, 0, 1, bob.EaseInOutCubic)
				changed = changed || v.Changed(t, prev)
				robot.LeftPupil = v.Get(t) * 0.5
				robot.RightPupil = v.Get(t) * 0.5

				o, v = 1.5, nil
				v.Add(s+o, s+o+1, 0, 1, bob.EaseInOutCubic)
				changed = changed || v.Changed(t, prev)
				robot.Neck = v.Get(t) * 0.25

				o, v = 2.25, nil
				v.Add(s+o, s+o+1, 0, 1, bob.EaseInOutCubic)
				changed = changed || v.Changed(t, prev)
				robot.Foot = v.Get(t) * 1

				o, v = 3.0, nil
				v.Add(s+o, s+o+1, 0, 1, bob.EaseInOutCubic)
				changed = changed || v.Changed(t, prev)
				robot.LeftArm = v.Get(t) * 0.5
				robot.RightArm = v.Get(t) * 0.5
			} else {
				robot.Sleep()
				robot.LeftPupil = 0
				robot.RightPupil = 0
			}

			mesh.Add(robot.CreateMesh())
		}
	}

	var v bob.Float64Variable
	v.Add(0, 6, 10, 50, bob.EaseInOutQuad)
	changed = changed || v.Changed(t, prev)
	fovy := v.Get(t)

	mesh.SmoothNormalsThreshold(Radians(30))
	scene.Add(mesh)

	if !changed {
		return nil, nil
	}

	material := DiffuseMaterial(White)
	scene.Add(NewPlane(V(0, 0, 0), V(0, 0, 1), material))
	// scene.Add(NewPlane(V(0, -20, 0), V(0, 1, 0), material))

	light := LightMaterial(White, 100)
	scene.Add(NewSphere(V(0, 10, 20), 2, light))

	camera := LookAt(V(0, 10, 4), V(0, 0, 3.5), V(0, 0, 1), fovy)
	return &scene, &camera
}

func main() {
	prevT := -1.0
	prevPath := ""
	t0 := 0.0
	t1 := 8.0
	fps := 25.0

	var wg sync.WaitGroup
	for i := 0; ; i++ {
		t := t0 + float64(i)/fps
		if t > t1 {
			break
		}
		path := fmt.Sprintf("frame%06d.png", i)
		fmt.Println(path)
		scene, camera := createScene(t, prevT)
		if scene == nil {
			os.Symlink(prevPath, path)
			continue
		}
		sampler := NewSampler(4, 4)
		sampler.LightMode = LightModeAll
		sampler.SpecularMode = SpecularModeFirst
		// sampler := NewDirectSampler()
		FrameRender(path, 4, scene, camera, sampler, 1920, 1080, 1, &wg)
		prevT = t
		prevPath = path
	}
	wg.Wait()
}
