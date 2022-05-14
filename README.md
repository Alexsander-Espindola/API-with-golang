# API-with-golang

# Modelo MVVM

## Camada model

    - Vai descreve a lógica de negócio da aplicação

## View

    -  Vai representar os componentes da inteface do usuário
    -  Vai exibir os dado vindo da Model

## ViewModel

    - Manter o estádo e ativar os eventos na view atravez de comando e data bindings
    - Vai fazer a ponte com a model e vai carregar a lógica da aplicação
    - Vinculação dos dados para a visualização definindo membros que são que controlam dados
    - Ponto de integraçãoo com outros serviços como o banco de dados

# GO

# Dockerizando o GoLang

Vídeo de base: https://www.youtube.com/watch?v=uDCzxwFT2-w&ab_channel=AprendaGolang

- Iniciando o Daemon do Docker
  Para consultar o status do daemon:

```
sudo systemctl status docker
```

Caso Active esteja parado:

```
sudo systemctl start docker)
```

- Adicionando um usuário ao docker

```
sudo usermod -aG docker $USER
```

Para ativar as alterações:

```
newgrp docker
```

- Ativar daemon do Docker para iniciar com o boot:

```
sudo systemctl enable docker
```

- Go mod:

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

- Rodando o docker:

```
docker run -p 8080:8080 my-server
```
