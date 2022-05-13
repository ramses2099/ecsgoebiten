package main

// package ebiten
// go get github.com/hajimehoshi/ebiten/v2
// packge uuid
// go get github.com/google/uuid

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ramses2099/ecsgoebiten/ecs"
)

var (
	entitymanager *ecs.ManagerEntity
	shipImage     *ebiten.Image
)

// init function
func init() {
	var err error
	shipImage, _, err = ebitenutil.NewImageFromFile("asset/ship_0001.png")
	if err != nil {
		log.Fatal(err)
	}

}

type Game struct{ inited bool }

func (g *Game) Init() {
	defer func() {
		g.inited = true
	}()
}

func (g *Game) Update() error {
	if !g.inited {
		g.Init()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// Write your game's rendering.
	msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	//draw image
	screen.DrawImage(shipImage, nil)

	// ecs.Each(entitymanager, ecs.Transform{}, func(id ecs.Id, a interface{}) {
	// 	transform := ecs.Transform{}
	// 	ok := ecs.Read(entitymanager, id, &transform)
	// 	if ok {
	// 		transform.X += 1
	// 		fmt.Printf(" x: %v y %v", transform.X, transform.Y)
	// 	}
	// })

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("ECS Game")

	//entity manager
	entitymanager = ecs.NewManagerEntity()

	entityId := entitymanager.NewId()
	ecs.Write(entitymanager, entityId, ecs.Transform{X: 10, Y: 20})

	transform := ecs.Transform{}
	ok := ecs.Read(entitymanager, entityId, &transform)
	if ok {
		fmt.Printf(" x: %v y %v", transform.X, transform.Y)
	}

	//component

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
