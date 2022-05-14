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
- Go mod init:
```
go mod init github.com/Alexsander-Espindola...
```

- Rodando o docker:
```
docker run -p 8080:8080 my-server
```
