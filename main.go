package main

import (
	"fmt"
)

// IGun interface mendefinisikan metode yang harus dimiliki oleh setiap jenis senjata
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Gun adalah struktur dasar yang mengimplementasikan interface IGun
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

// AK47 adalah jenis senjata yang meng-embed struktur Gun dan secara tidak langsung mengimplementasikan semua metode IGun
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

// Musket adalah jenis senjata lain yang juga meng-embed struktur Gun dan secara tidak langsung mengimplementasikan semua metode IGun
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

// getGun adalah fungsi factory yang mengembalikan IGun berdasarkan tipe senjata yang diberikan
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	// Membuat IGun menggunakan fungsi factory getGun
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

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
