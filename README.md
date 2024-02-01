# 🐳 GraphQL usando go e docker

Explorando o GraphQL com o uso da linguagem Go.


## 💻 Como executar o projeto

1. execute o comando docker-compose para iniciar os containers:

```bash
docker-compose up -d
```

2. execute o comando docker-compose exec para acessar o contêiner:


```bash
docker-compose exec -it graphql bash
```

3. execute o comando para iniciar o servidor em Go:


```bash
go run cmd/server/server.go
```

4. Acesse a aplicação em seu navegador:

```bash
http://localhost:8080
```