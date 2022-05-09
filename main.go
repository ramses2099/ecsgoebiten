package main

// package ebiten
// go get github.com/hajimehoshi/ebiten/v2
// packge uuid
// go get github.com/google/uuid

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ramses2099/ecsgoebiten/components"
	cons "github.com/ramses2099/ecsgoebiten/constants"
	"github.com/ramses2099/ecsgoebiten/entities"
)

var (
	e entities.Entity
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// et := "Id Entity " + e.Id + "" + "\n component name " + e.GetComponent("Health").GetName() + ""
	// ebitenutil.DebugPrint(screen, et)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("ECS Game")
	//
	e = entities.NewEntity()
	c := components.NewComponentHealth(100)
	e.AddComponents(cons.ComponentHealth, c)

	p := components.NewComponentPosition(20, 25)
	e.AddComponents(cons.ComponentPosition, p)

	for _, cp := range e.GetComponents() {
		fmt.Println("Name component", cp.GetComponentName())
	}

	em := entities.NewEntityManager(e)

	fmt.Printf("Lenght Entity Manager %v \n", em.GetLength())

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
