package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Veiculo struct {
	gorm.Model
	Placa string
}

func bootstapDatabase() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("garagem.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Veiculo{})

	return db, err
}

func bootstrapApplication() {
	fmt.Println("Bem-vindo ao sistema de garagem!")
}

func entrance() string {
	var plate string

	fmt.Println("Entrada de veículos:")
	fmt.Scanln(&plate)

	return plate
}

func exit() string {
	var plate string

	fmt.Println("Saída de veículos:")
	fmt.Scanln(&plate)

	return plate
}

func search() string {
	var plate string

	fmt.Println("Buscar veículo:")
	fmt.Scanln(&plate)

	return plate
}

func list() {
	fmt.Println("Lista de veículos:")
}

func actions() (string, int) {
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Entrada")
	fmt.Println("2 - Saída")
	fmt.Println("3 - Busca")
	fmt.Println("4 - Lista")
	fmt.Println("5 - Sair")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		return entrance(), option
	case 2:
		return exit(), option
	case 3:
		return search(), option
	case 4:
		return "list", option
	case 5:
		return "sair", option
	default:
		return "", option
	}
}

func loop(db *gorm.DB, err error) {
	var plate, option = actions()

	switch option {
	case 1:
		db.Create(&Veiculo{Placa: plate})
		actions()
	case 2:
		db.Delete(&Veiculo{}, "placa = ?", plate)
		actions()
	case 3:
		var veiculo Veiculo
		db.First(&veiculo, "placa = ?", plate)
		fmt.Println("Veículo:", veiculo)
		actions()
	case 4:
		list()
		var veiculos []Veiculo
		db.Find(&veiculos)
		fmt.Println("Veículos:", veiculos)
		actions()
	default:
		return
	}
}

func main() {
	var db, err = bootstapDatabase()

	bootstrapApplication()

	loop(db, err)
}
