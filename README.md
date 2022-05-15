# Aprendizados do projeto

## Ferramentas
  - GO lang
  - gRPC

## Fontes que eu usei para o estudo

    - Playlist Go lang do zero

      https://www.youtube.com/watch?v=_MkQLDMak-4&list=PL5aY_NrL1rjucQqO21QH8KclsLDYu1BIg&index=1&ab_channel=FullCycle

    -

# GO

- Características:
  - Expansiva, limpa e eficiente
  - Simples: usa um paradigma procedural
  - Confiável e Eficiente: feita para garantir que as funções sejam executadas de maneira fácil, concorrente e rápida
  - Compilada e tipada: garante que não subirá muito dos erros de código para a produção
  - Pode gerar builds de difentes sistemas operaconais na sua aplicação

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

# gRPC

Fonte:

https://www.youtube.com/watch?v=F4t3ZBVMlvo&ab_channel=C%C3%B3digoFonteTV

- Caractrísticas:

  - Código aberto
  - Alto desempenho de chamadas ente sistemas
  - Tem o objetivo de ser mais leve e performático do que as APIs REST tradicionais
  - Usa versão HTTP2 para codificar o pacote de forma binária

- Padrão RPC:

  - Remote procedure call
  - Suporte para:
    1. Balanceamento de carga
    2. Rastreamento
    3. Verificação de integridade
    4. Autentificação
  - RPC Tradicional:
    - Modelo síncrono de cliente servidor
  - RPC assícrono
    - Modelo assíncrono de cliente servidor

- Protobuf
  - IDL: Linguagem de definição de interface
  - Toda a comunicação do sistema definida no sistema do protobuf
