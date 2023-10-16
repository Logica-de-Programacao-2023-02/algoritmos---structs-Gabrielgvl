package main

import "fmt"

//Escreva uma função que receba um ponteiro para um
//booleano e altere o
//valor desse booleano para o seu inverso.

//Escreva uma função que receba um ponteiro para uma
//struct que contém informações de um produto
//(nome, preço e quantidade em estoque).
//A função deve atualizar o preço desse produto para
//um novo valor fornecido como argumento.

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {
	x := 42
	y := 55

	swap(&x, &y)
	fmt.Println(x, y)

	b := true
	invert(&b)
	fmt.Println(b)

	p := Produto{
		Nome:       "brinqudo",
		Preco:      10,
		Quantidade: 1,
	}

	atualizaPreco(&p, 100)

	fmt.Println(p)
}

func atualizaPreco(p *Produto, novoPreco float64) {
	p.Preco = novoPreco
}

func invert(b *bool) {
	*b = !*b
}

func swap(ptr1, ptr2 *int) {
	aux := *ptr1
	*ptr1 = *ptr2
	*ptr2 = aux
	// -------------------------
	*ptr1, *ptr2 = *ptr2, *ptr1
}
