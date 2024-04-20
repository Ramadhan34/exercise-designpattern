package main

import (
	"errors"
	"fmt"
)

// IGun interface
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Gun struct
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

// AK47 struct
type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// Musket struct
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// getGun function
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, errors.New("Wrong gun type passed")
}

func main() {
	// Membuat IGun menggunakan fungsi factory getGun
	ak47, err := getGun("ak47")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	musket, err := getGun("musket")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Memanggil metode printDetails pada IGun yang dibuat
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
