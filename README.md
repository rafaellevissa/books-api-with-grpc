# Books CRUD 📚

## Api de cadastro de livros.

**Stack estudada** 🛠️

 > * GoLang 
 > * Postgres 
 > * Docker
 > * GRPC 

 --- 

 ## Run (Kafka) 🏃
> $ cd ./apiKafka

> $ docker-compose up -d

 ## Run (Server Http or Grpc) 🏃
> $ docker-compose up -d

> $ go run main.go 
or
> $ make run

Obs: A colletion para ser usada no insomnia para realizar as requisições estão no utils/Colletion-Books-Api-With-Grpc.har, basta importar esse arquivo lá no insomnia

 ## Run Client Grpc
 > $ make run-client

 Obs: Para rodar o client grpc precisamos do servidor grpc rodando junto com o banco de dados. Caso queira trocar a função de criar para visualizar, vá em client/cmd/main.go e altere o valor da variável "routehttp". Deixei dessa forma estática e não tão usual para demonstrar apenas a utilização de um client Grpc.