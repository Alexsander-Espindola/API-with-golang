## Sumário

[Rodando a aplicação](#Rodando a aplicação)

# Rodando a aplicação

## Rodando com gRPC server

1. Iniciando mongodb com docker

```
docker run --name mongodb-golang -d -p 27017:27017 mongo
```

- Se já tver rodado o comando acima antes rode:

```
docker container start mongodb-golang
```

2. Iniciando o servidor para o user

```
go run src/proto/cmd/server/server.go
```

3. Iniciando o cliente para o user

```
go run src/proto/cmd/client/client.go
```

## Rodando os testes

# Aprendizados do projeto

- GO lang
- gRPC

# GO

## Características:

- Expansiva, limpa e eficiente
- Simples: usa um paradigma procedural
- Confiável e Eficiente: feita para garantir que as funções sejam executadas de maneira fácil, concorrente e rápida
- Compilada e tipada: garante que não subirá grande parte dos possíveis erros de código para a produção
- Pode gerar builds de diferentes sistemas operacionais na sua aplicação

## Go mod:

- Go mod init

```
go mod init [module-path]
```

- Exemplo:

```
go mod init github.com/Alexsander-Espindola...
```

O comando go mod init inicializa e grava um novo arquivo go.mod no diretório atual, na verdade criando um novo módulo enraizado no diretório atual.

- Go mod tidy

```
go mod tidy
```

Garante que o arquivo go.mod corresponda ao código-fonte no módulo. Ele adiciona quaisquer requisitos de módulo ausentes necessários para construir os pacotes e dependências do módulo atual e remove os requisitos em módulos que não fornecem nenhum pacote relevante. Ele também adiciona as entradas ausentes ao go.sum e remove as entradas desnecessárias.

- Go mod vendor

```
go mod vendor
```

O comando go mod vendor constrói um diretório chamado vendor no diretório raiz do módulo principal que contém cópias de todos os pacotes necessários para suportar compilações e testes de pacotes no módulo principal. Os pacotes que são importados apenas por testes de pacotes fora do módulo principal não são incluídos. Tal como acontece com go mod arrumado e outros comandos de módulo, as restrições de construção, exceto para ignorar, não são consideradas ao construir o diretório do fornecedor.

- Go mod graph

```
go mod graph [-go=version]
```

O comando go mod graph imprime o gráfico de requisitos do módulo (com substituições aplicadas) em forma de texto.

# Dockerizando o GoLang

Vídeo de base: https://www.youtube.com/watch?v=uDCzxwFT2-w&ab_channel=AprendaGolang

## Iniciando o Daemon do Docker

- Para consultar o status do Daemon do Docker:

```
sudo systemctl status docker
```

- Caso Active esteja parado:

```
sudo systemctl start docker)
```

- Adicionando um usuário ao docker

```
sudo usermod -aG docker $USER
```

- Para ativar as alterações:

```
newgrp docker
```

- Ativar daemon do Docker para iniciar com o boot:

```
sudo systemctl enable docker
```

- Rodando o docker:

```
docker run -p 8080:8080 my-server
```

## Docker com Mongo

- Fonte:

```
https://www.mongodb.com/compatibility/docker
```

- Iniciando

```
docker run --name mongodb -d mongo
```

- Criando porta

```
docker run --name mongodb -d -p 27017:27017 mongo
```

# Conectando ao MongoDB com Go

- Instalação

```
go get go.mongodb.org/mongo-driver/mongo
```

- Conexão com o Mongo

```
import (
   "context"
   "time"

   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
   "go.mongodb.org/mongo-driver/mongo/readpref"
)

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
```

- Work with BSON

```
https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/

```

# gRPC

- Um artigo incrível sobre o gRPC (ainda não li tudo)

```
https://www.linkedin.com/pulse/grpc-jose-r-f-junior/?originalSubdomain=pt
```

## Protoc

- Instalação (Linux)

```
$ sudo apt install protobuf-compiler
```

- Verificação da instalação

```
$ protoc --version
```

- Resposta

```
libprotoc *.*
```

- Go get para o gRPC

```
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH="$PATH:$(go env GOPATH)/bin"
```

- Iniciando o protocol buffer
- Cria uma pasta proto com um arquivo inicializador do proto:

```
// caminho: src/proto/protoUser.proto
syntax = "proto3";
option go_package="proto/pb";
```

- Cria algumas messages e um service

```
message Request {
 string message = 1;
}
message Response {
 int32 status = 1;
}
service SendMessage {
 rpc RequestMessage (Request) returns (Response){}
}
```

- Comando para inicializar o proto

```
cd src
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

- Características:

- Código aberto
- Alto desempenho de chamadas ente sistemas
- Tem o objetivo de ser mais leve e performático do que as APIs REST tradicionais
- Usa versão HTTP 2 para codificar o pacote de forma binária

- Padrão RPC:

- Remote procedure call
- Suporte para:
  1.  Balanceamento de carga
  2.  Rastreamento
  3.  Verificação de integridade
  4.  Autentificação
- RPC Tradicional:
  - Modelo síncrono de cliente servidor
- RPC assíncrono

  - Modelo assíncrono de cliente servidor

- Protobuf
- IDL: Linguagem de definição de interface
- Toda a comunicação do sistema definida no sistema do protobuf

## Fontes que eu usei para o estudo

- Golang (A Linguagem do Futuro?) // Dicionário do Programador

```
https://www.youtube.com/watch?v=2kyNEf9IsBQ&ab_channel=C%C3%B3digoFonteTV
```

- Playlist Golang do zero

```
https://www.youtube.com/watch?v=_MkQLDMak-4&list=PL5aY_NrL1rjucQqO21QH8KclsLDYu1BIg&index=1&ab_channel=FullCycle
```

- O que é gRPC?

```
https://www.youtube.com/watch?v=AxYEyvX3xxI&ab_channel=FullCycle
```

- gRPC (Nem toda comunicação usa API) // Dicionário do Programador

```
https://www.youtube.com/watch?v=F4t3ZBVMlvo&ab_channel=C%C3%B3digoFonteTV
```

- Como fazer um Dockerfile otimizado para Golang

```
https://www.youtube.com/watch?v=uDCzxwFT2-w
```

- Quick Start: Golang & MongoDB

```
https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
```

- GoLang - Inserindo e buscando dados do MongoDB

```
https://www.youtube.com/watch?v=TtnI6nPhzgQ&ab_channel=HunCoding
```

- Mongo Go Drive

```
https://github.com/mongodb/mongo-go-driver
```

- Quick Start: Golang & MongoDB

```
https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
```

- [GOLANG] Como utilizar o gRPC com GO!

```
https://www.youtube.com/watch?v=LuS59XHdKG8&ab_channel=Yuri
https://github.com/hyperyuri/grpc-with-go
```

- gRPC na Prática

```
https://www.youtube.com/watch?v=VMCRRznA3g0&ab_channel=FullCycle
```

- Plataforma de Desafios em Go Lang - Parte 1

```
https://www.youtube.com/watch?v=1ig4JWzn30g&t=1434s
```

- Go Lang: Validação e testes de usuário - Plataforma de desafios - Parte 2

```
https://www.youtube.com/watch?v=JP-D1In0juw&t=2085s&ab_channel=FullCycle
```
