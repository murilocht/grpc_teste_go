package modelo

import "github.com/google/uuid"

type Pessoa struct {
	Id   string
	Nome string
}

type Pessoas struct {
	Pessoa []Pessoa
}

func (p *Pessoas) Add(pessoa *Pessoa) {
	p.Pessoa = append(p.Pessoa, *pessoa)
}

func AddPessoa() *Pessoa {
	pessoa := Pessoa{}
	pessoa.Id = string(uuid.NewString())

	return &pessoa
}

func AddPessoas() *Pessoas {
	return &Pessoas{}
}
