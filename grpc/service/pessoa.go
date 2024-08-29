package service

import (
	"context"
	"grpc_teste/grpc/pb"
	"grpc_teste/modelo"
)

type PessoaGrpcService struct {
	pb.UnimplementedPessoaServiceServer
	Pessoas *modelo.Pessoas
}

func NovaPessoaGrpcService() *PessoaGrpcService {
	return &PessoaGrpcService{}
}

func (p *PessoaGrpcService) CreatePessoa(ctx context.Context, req *pb.Pessoa) (*pb.PessoaResut, error) {
	pessoa := modelo.AddPessoa()
	pessoa.Nome = req.Nome
	p.Pessoas.Add(pessoa)

	return &pb.PessoaResut{Id: pessoa.Id, Nome: pessoa.Nome}, nil
}
