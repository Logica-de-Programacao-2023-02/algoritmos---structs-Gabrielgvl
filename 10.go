package main

import "fmt"

//Crie uma struct chamada Filme com os campos "título",
//"diretor", "ano" e "avaliações". O campo "avaliações"
//deve ser um slice de inteiros representando as notas
//que o filme recebeu dos espectadores. Escreva funções
//que permitam adicionar ou remover avaliações do filme,
//calcular a média das avaliações e imprimir as informações
//do filme e sua média de avaliações.

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []Avaliacao
}

type Avaliacao struct {
	Nome string
	Nota int
}

func main() {
	f := Filme{
		Titulo:  "Matrix",
		Diretor: "N lembro",
		Ano:     1997,
	}

	f = AdicionarAvaliacao(f, "Gabriel", 5)
	f = AdicionarAvaliacao(f, "Gabriel", 5)
	f = AdicionarAvaliacao(f, "Gabriel", 3)
	f = AdicionarAvaliacao(f, "Gabriel", 10)

	FilmeInfo(f)

	f = RemoverAvaliacao(f, 2)

	FilmeInfo(f)
}

func AdicionarAvaliacao(f Filme, nomeUsuario string, nota int) Filme {
	f.Avaliacoes = append(f.Avaliacoes, Avaliacao{
		Nome: nomeUsuario,
		Nota: nota,
	})
	return f
}

func RemoverAvaliacao(f Filme, posicao int) Filme {
	f.Avaliacoes = append(f.Avaliacoes[:posicao], f.Avaliacoes[posicao+1:]...)
	return f
}

func CalcularMedia(f Filme) float64 {
	var total int

	for _, avaliacao := range f.Avaliacoes {
		total += avaliacao.Nota
	}

	return float64(total) / float64(len(f.Avaliacoes))
}

func FilmeInfo(f Filme) {
	fmt.Printf("O filme %s (%d) do diretor %s tem a média de avaliações de %.2f\n",
		f.Titulo,
		f.Ano,
		f.Diretor,
		CalcularMedia(f))
}
