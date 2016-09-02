package bob

import (
	"math"
	"math/rand"

	. "github.com/fogleman/pt/pt"
)

type Robot struct {
	Position         Vector
	Rotation         float64
	Tilt             float64
	Foot             float64
	LeftArm          float64
	RightArm         float64
	Neck             float64
	HeadRotation     float64
	HeadTilt         float64
	LeftEye          float64
	RightEye         float64
	LeftPupil        float64
	RightPupil       float64
	LeftPupilAspect  float64
	RightPupilAspect float64
	Antenna          float64
	Battery          float64
	Number           int
}

func NewRobot() *Robot {
	robot := &Robot{}
	robot.Battery = 1
	robot.LeftEye = 0
	robot.RightEye = 0
	robot.LeftPupil = 0.5
	robot.LeftPupilAspect = 1
	robot.RightPupil = 0.5
	robot.RightPupilAspect = 1
	robot.Antenna = 1
	robot.HeadRotation = 0
	robot.Neck = 0.25
	robot.Foot = 1
	robot.LeftArm = 0
	robot.RightArm = 0
	// robot.Sleep()
	// robot.Wake()
	return robot
}

func (robot *Robot) Sleep() {
	robot.Battery = 0
	robot.LeftPupil = 1
	robot.RightPupil = 1
	robot.LeftPupilAspect = 10
	robot.RightPupilAspect = 10
	robot.Neck = 0
	robot.Foot = 0
	robot.LeftArm = 0
	robot.RightArm = 0
	robot.Antenna = 0
	robot.HeadRotation = 0
	robot.HeadTilt = 0
	robot.LeftEye = 0
	robot.RightEye = 0
}

func (robot *Robot) Wake() {
	robot.LeftPupil = 0.5
	robot.RightPupil = 0.5
	robot.Neck = 0.5
	robot.Foot = 1
	robot.LeftArm = 0.25
	robot.RightArm = 0.25
	robot.Antenna = 1
	robot.HeadTilt = 0
	robot.HeadRotation = 0.25
}

func (robot *Robot) Random() {
	robot.LeftPupil = rand.Float64()*0.25 + 0.5
	robot.RightPupil = robot.LeftPupil
	robot.Battery = rand.Float64()
	robot.Rotation = rand.Float64() * 360
	robot.Tilt = 0 //rand.Float64()
	robot.Foot = rand.Float64()
	robot.LeftArm = (rand.Float64()*1 - 0) * 0.5
	robot.RightArm = (rand.Float64()*1 - 0) * 0.5
	robot.Neck = rand.Float64() * 0.75
	robot.HeadRotation = rand.Float64()*2 - 1
	robot.LeftEye = rand.Float64()*0.8 + 0.2
	robot.RightEye = rand.Float64()*0.8 + 0.2
	// robot.LeftPupil = rand.Float64()
	// robot.RightPupil = rand.Float64()
	// robot.LeftPupilAspect = rand.Float64()
	// robot.RightPupilAspect = rand.Float64()
	robot.Antenna = rand.Float64()
}

func (robot *Robot) PointHead(point Vector) {
	d := point.Sub(robot.HeadPosition())
	xy := Vector{d.X, d.Y, 0}.Normalize()

	tilt := math.Atan2(d.Z, Vector{d.X, d.Y, 0}.Length())
	robot.HeadTilt = -Degrees(tilt) / 30

	a := Radians(robot.Rotation)
	forward := Vector{math.Sin(a), math.Cos(a), 0}

	dot := Clamp(forward.Dot(xy), -1, 1)
	robot.HeadRotation = Degrees(math.Acos(dot)) / 90
	if xy.Cross(forward).Dot(V(0, 0, 1)) < 0 {
		robot.HeadRotation = -robot.HeadRotation
	}
}

func (robot *Robot) PointBody(point Vector) {
	d := point.Sub(robot.Position)
	robot.Rotation = Degrees(math.Atan2(d.X, d.Y))
}

func (robot *Robot) HeadPosition() Vector {
	z := robot.Foot*0.75 + robot.Neck + 3
	return robot.Position.Add(V(0, 0, z))
}

func (robot *Robot) Heading() Vector {
	a := Radians(robot.Rotation)
	return Vector{math.Sin(a), math.Cos(a), 0}
}

func (robot *Robot) CreateMesh() *Mesh {
	// left eye
	leftEye := robot.eye()
	leftPupil := robot.pupil()
	lx := robot.LeftPupil * robot.LeftPupilAspect
	lz := robot.LeftPupil
	if robot.LeftPupilAspect > 1 {
		lx = robot.LeftPupil
		lz = robot.LeftPupil / robot.LeftPupilAspect
	}
	leftPupil.Transform(Scale(V(lx, 1, lz)))
	leftPupil.Transform(Translate(V(0, 0.51, 0)))
	leftEye.Add(leftPupil)
	leftEye.Transform(Translate(V(-0.4, 0.6, 0.75)))
	leftEye.Transform(Translate(V(0, robot.LeftEye*0.5, 0)))

	// right eye
	rightEye := robot.eye()
	rightPupil := robot.pupil()
	rx := robot.RightPupil * robot.RightPupilAspect
	rz := robot.RightPupil
	if robot.RightPupilAspect > 1 {
		rx = robot.RightPupil
		rz = robot.RightPupil / robot.RightPupilAspect
	}
	rightPupil.Transform(Scale(V(rx, 1, rz)))
	rightPupil.Transform(Translate(V(0, 0.51, 0)))
	rightEye.Add(rightPupil)
	rightEye.Transform(Translate(V(0.4, 0.6, 0.75)))
	rightEye.Transform(Translate(V(0, robot.RightEye*0.5, 0)))

	// antenna
	antenna := robot.antenna()
	antenna.Transform(Translate(V(0, 0, 1.6)))
	antenna.Transform(Translate(V(0, 0, robot.Antenna*0.5)))

	// neck
	neck := robot.neck()

	// head
	head := robot.head()
	head.Add(leftEye)
	head.Add(rightEye)
	head.Add(antenna)
	head.Transform(Rotate(V(1, 0, 0), Radians(robot.HeadTilt*30)))
	head.Transform(Rotate(V(0, 0, 1), Radians(robot.HeadRotation*90)))
	head.Transform(Translate(V(0, 0, robot.Neck)))
	head.Add(neck)
	head.Transform(Translate(V(0, 0, 3)))

	// left arm
	leftArm := robot.arm()
	leftArm.Transform(Rotate(V(1, 0, 0), Radians(90-robot.LeftArm*90)))
	leftArm.Transform(Translate(V(-1.125, 0, 2.375)))

	// right arm
	rightArm := robot.arm()
	rightArm.Transform(Rotate(V(1, 0, 0), Radians(90-robot.RightArm*90)))
	rightArm.Transform(Translate(V(1.125, 0, 2.375)))

	// wheel
	wheel := robot.wheel()

	// body
	body := robot.body()
	body.Add(head)
	body.Add(leftArm)
	body.Add(rightArm)
	for i := 0; i < 8; i++ {
		color := Black
		// if float64(i) < robot.Battery*8 {
		// 	color = HexColor(0x04BFBF)
		// }
		if (1<<uint(i))&robot.Number != 0 {
			color = White
		}
		button := robot.button(color)
		button.Transform(Translate(V(0, 0.95, 0.625+float64(i)*0.25)))
		body.Add(button)
	}
	body.Transform(Translate(V(0, 0, robot.Foot*0.75)))
	body.Transform(Rotate(V(1, 0, 0), Radians(robot.Tilt*30)))
	body.Transform(Rotate(V(0, 0, 1), Radians(robot.Rotation)))
	body.Add(wheel)
	body.Transform(Translate(robot.Position))
	return body
}

func (robot *Robot) button(color Color) *Mesh {
	mesh := buttonMesh.Copy()
	mesh.SetMaterial(GlossyMaterial(color, 1.5, Radians(10)))
	return mesh
}

func (robot *Robot) antenna() *Mesh {
	mesh := antennaMesh.Copy()
	mesh.SetMaterial(GlossyMaterial(HexColor(0x374140), 1.5, Radians(10)))
	return mesh
}

func (robot *Robot) head() *Mesh {
	mesh := headMesh.Copy()
	mesh.SetMaterial(GlossyMaterial(HexColor(0xA9CF54), 1.3, Radians(10)))
	return mesh
}

func (robot *Robot) eye() *Mesh {
	mesh := eyeMesh.Copy()
	mesh.SetMaterial(GlossyMaterial(White, 1.5, Radians(10)))
	return mesh
}

func (robot *Robot) pupil() *Mesh {
	mesh := pupilMesh.Copy()
	mesh.SetMaterial(DiffuseMaterial(Black))
	return mesh
}

func (robot *Robot) neck() *Mesh {
	mesh := neckMesh.Copy()
	mesh.SetMaterial(GlossyMaterial(White, 1.5, Radians(10)))
	return mesh
}

func (robot *Robot) body() *Mesh {
	mesh := bodyMesh.Copy()
	mesh.SetMaterial(GlossyMaterial(HexColor(0xF7E967), 1.3, Radians(10)))
	return mesh
}

func (robot *Robot) arm() *Mesh {
	mesh := armMesh.Copy()
	mesh.SetMaterial(GlossyMaterial(White, 1.5, Radians(10)))
	return mesh
}

func (robot *Robot) wheel() *Mesh {
	mesh := wheelMesh.Copy()
	mesh.SetMaterial(GlossyMaterial(White, 1.5, Radians(10)))
	return mesh
}

// load STL meshes on init

var (
	buttonMesh  *Mesh
	antennaMesh *Mesh
	headMesh    *Mesh
	eyeMesh     *Mesh
	pupilMesh   *Mesh
	neckMesh    *Mesh
	bodyMesh    *Mesh
	armMesh     *Mesh
	wheelMesh   *Mesh
)

func init() {
	buttonMesh = loadMesh("stl/button.stl")
	antennaMesh = loadMesh("stl/antenna.stl")
	headMesh = loadMesh("stl/head.stl")
	eyeMesh = loadMesh("stl/eye.stl")
	pupilMesh = loadMesh("stl/pupil.stl")
	neckMesh = loadMesh("stl/neck.stl")
	bodyMesh = loadMesh("stl/body.stl")
	armMesh = loadMesh("stl/arm.stl")
	wheelMesh = loadMesh("stl/wheel.stl")
}

func loadMesh(path string) *Mesh {
	mesh, err := LoadSTL(path, DiffuseMaterial(Black))
	if err != nil {
		panic(err)
	}
	return mesh
}
