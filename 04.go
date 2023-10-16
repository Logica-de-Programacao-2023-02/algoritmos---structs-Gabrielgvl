package main

import "fmt"

//Crie uma struct chamada Playlist com os campos "nome" e
//"músicas".
//O campo "músicas" deve ser um slice de structs,
//cada uma representando uma música com os campos
//"título", "artista" e "duração".
//Escreva uma função que receba uma Playlist como
//parâmetro e imprima o nome da playlist,
//o tempo total da playlist e as informações de cada música.

type Playlist struct {
	Nome    string
	Musicas []Musica
}

type Musica struct {
	Titulo  string
	Artista string
	Duracao int
}

func main() {
	p := Playlist{
		Nome: "Playlista Vrau",
		Musicas: []Musica{
			{
				Titulo:  "Vrau 1",
				Artista: "Vrauson",
				Duracao: 100,
			},
			{
				Titulo:  "Vrau 2",
				Artista: "Vrauson",
				Duracao: 150,
			},
			{
				Titulo:  "Vrau 3",
				Artista: "Vrausonelson",
				Duracao: 200,
			},
		},
	}
	PlaylistInfo(p)
}

func PlaylistInfo(p Playlist) {
	total := 0 // segundos
	for _, musica := range p.Musicas {
		total = total + musica.Duracao
	}
	fmt.Printf("A playlist %s tem %d:%d\n", p.Nome, total/60, total%60)
	for _, musica := range p.Musicas {
		fmt.Printf("A música %s do artista %s tem %d:%d\n",
			musica.Titulo,
			musica.Artista,
			musica.Duracao/60,
			musica.Duracao%60,
		)
	}
}
