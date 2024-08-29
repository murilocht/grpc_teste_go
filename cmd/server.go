package main

import (
	"encoding/json"
	"grpc_teste/grpc/pb"
	"grpc_teste/grpc/service"
	"grpc_teste/modelo"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var lista = modelo.AddPessoas()

func main() {
	go startGrpc()

	http.HandleFunc("/", GerenciarPessoas)
	http.ListenAndServe(":3000", nil)
}

func startGrpc() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pessoaService := service.NovaPessoaGrpcService()
	pessoaService.Pessoas = lista

	pb.RegisterPessoaServiceServer(grpcServer, pessoaService)
	grpcServer.Serve(listener)
}

func GerenciarPessoas(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(lista)
	w.Write([]byte(json))
}
