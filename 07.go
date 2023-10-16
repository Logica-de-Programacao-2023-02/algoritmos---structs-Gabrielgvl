package main

import "fmt"

//Crie uma struct chamada Animal com os campos "nome",
//"espécie", "idade" e "som". Escreva funções que
//permitam modificar o som que o animal faz e uma função que
//imprima as informações do animal e o som que ele faz.

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func main() {
	animal := Animal{
		Nome:    "Galinha",
		Especie: "Frango",
		Idade:   10,
		Som:     "Pó",
	}

	AnimalInfo(animal)

	animal = AlteraSom(animal, "Có")

	AnimalInfo(animal)
}

func AlteraSom(animal Animal, novoSom string) Animal {
	animal.Som = novoSom
	return animal
}

func AnimalInfo(banana Animal) {
	fmt.Printf("O animal %s da espécie %s e com a idade de %d faz o seguinte som: %s\n",
		banana.Nome,
		banana.Especie,
		banana.Idade,
		banana.Som,
	)
}
