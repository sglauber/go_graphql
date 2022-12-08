<!-- BADGES/SHIELDS -->
<!--
*** Markdown "reference style" para facilitar a leitura.
-->

<!-- LOGO -->
<br />
<div align="center">
  <a href="https://www.studiosol.com.br/">
    <img src="https://i.ibb.co/j8DWsCT/Studio-Sol.png" alt="Logo" width="350" height="auto">
  </a>
</div>

<h2 align="center">Prova StudioSol - GraphQL API</h2>


<!-- REGRAS / REQUISITOS -->

## Regras do projeto 
Dada uma palavra contínua, e um conjunto de regras, Joaquim precisa verificar se a senha é válida baseada nas regras pedidas
  - minSize: tem pelo menos x caracteres.
  - minUppercase: tem pelo menos x caracteres maiúsculos
  - minLowercase: tem pelo menos x caracteres minúsculos
  - minDigit: tem pelo menos x dígitos (0-9)
  - minSpecialChars: tem pelo menos x caracteres especiais
    - Os caracteres especiais são os caracteres da seguinte string: "!@#$%^&*()-+\/{}[]"
  - noRepeted: (esta regra pode ser ignorada)

## Consulta de exemplo

```graphql
query {
  verify(password: "TesteSenhaForte!1223&", rules: [
    {rule: "minSize",value: 8},
		{rule: "minSpecialChars",value: 2},
		{rule: "noRepeted",value: 0},
		{rule: "minDigit",value: 4}
  ]){ noMatch, verify }
}
```


<!-- PRIMEIROS PASSOS PARA RODAR A APLICAÇÃO -->


## Estrutura de pastas

```
├── go.mod
├── go.sum
├── gqlgen.yml               - Configuração gqlgen, com as instruções e definições de diretórios
├── graph
│   ├── generated            - Pacotes que são gerados em tempo de execução pela lib
│   │   └── generated.go
│   ├── model                - Models da nossa API, sendo eles gerados automaticamente pelo gqlgen ou manualmente inseridos
│   │   └── models_gen.go
│   ├── resolver.go          - Resolver principal responsável pela injeção de dependência gerais da aplicação
│   ├── schema.graphqls      - Definição do schema, podendo ser dividido para facilitar a estruturação da API em projetos maiores
│   └── schema.resolvers.go  - O resolver para o schema definido acima com as funções necessárias para sua validação e manipulação
└── server.go                - Definição do servidor, endpoints, rota e portas
```

## Pré-requisitos

Para rodar o projeto é necessário que tenha o Docker instalado, você pode acessar o link [Get Docker][docker] para realizar o download.
Se preferir você poderá rodar utilizando `Go` em sua máquina local, cheque as instruções abaixo.

## Utilizando

O projeto funciona de maneira bem simples, ao iniciá-lo você terá acesso tanto ao servidor GraphQL quanto ao cliente através da interface do GraphiQL
onde você poderá visualizar e ralizar os testes da API.

Para realizar apenas a consulta você pode utilizar um cliente como [Postman][postman]

A api recebe uma consulta `verify` que aceita os parâmetros `password` uma string com a senha à ser validade e um array `rules` que deve ser preenchido com os objetos
de regras, contendo uma `rule` que é o nome da regra e um `value` que é o valor da regra à ser validada. Você pode seguir a consulta de exemplo para montar a sua consulta.

### Com Docker

Você pode buildar a imagem com `docker build`, ou apenas executar o container com `docker run`
É possível realizar o  tudo em um único comando, buildar e executar o contâiner, para isto utilize os comandos abaixo substituindo os valores

> Para facilitar realize os comandos de build na pasta do projeto onde se encontra o `Dockerfile`

- <nome_imagem> nome que deseja registrar para sua imagem
- <nome_container> nome que deseja registrar para seu container
- <porta:8080> você deve mapear a porta do host onde o container irá se conectar, como padrão pode utilizar `8080:8080`

```bash
  docker build -t <nome_imagem> . && docker run --name <nome_container> -p <porta:8080> <nome_imagem>
```

Você poderá apenas executar a imagem, isto fará com que o Docker realize o download da versão disponível no DockerHub, inicie a aplicação e remova o container quando o serviço for finalizado.

```bash
  docker run -p <porta:8080> --rm -d sglauber/go-graphql
```

Após executar o comando acima acesse a rota `localhost:<porta>/` para visualizar o GraphiQL.

### Go na máquina local

> Obs.: Por padrão o servidor será iniciado na porta: 8080

- Execute `go mod tidy` para adicionar automaticamente as dependências 
- Com `go run server.go` você irá iniciar o servidor 
- Então acesse a rota `localhost:8080/` para visualizar o GraphiQL
- Realize consultas na rota `localhost:8080/graphql`

## Experiência

Este foi meu primeiro contato com a linguagem `Go`, assim como APIs em graphql, foi prazeroso poder realizar este projeto e entender um pouco como funcionam estas tecnologias.
Decidi escolher `Go` para linguagem de implementação exatamente pelo desafio, acredito que o framework ou linguagem são ferramentas que se utilizadas corretamente nos auxiliam a chegar no objetivo de forma satisfatória e sou grato por poder tê-lo completado e feliz com o resultado, foi uma ótima oportunidades de aprendizado. 

A escolha da biblioteca `gqlen` se dá pelas funcionalidades de abstração, documentação e quantidade de exemplos que ela oferece se comparada à outras bibliotecas como `graphql-go` e `gophers`. Como podemos definir o schema-first ao invés de implementar interfaces e tipos diretamente no código ela facilitou o processo de aprendizagem e elimina boa parte dos problemas com relação à tipagem, porém, é um pouco chato de lidar quando se é necessário realizar o regenerate para mapear os models, já que ela irá também reescrever os resolvers.

## Referências

- [gqlgen docs][gqlgen]
- [qlgen-examples][qlgen-examples]
- [graphql Docs][graphql]
- [go.dev][go.dev]



<!-- Markdown Reference Style -->
[gqlgen]: https://gqlgen.com/getting-started/
[qlgen-examples]: https://github.com/99designs/gqlgen/tree/master/_examples
[graphql]: https://graphql.org/learn/
[go.dev]: https://go.dev/doc/
[docker]: https://docs.docker.com/get-docker/
[postman]: https://learning.postman.com/docs/sending-requests/graphql/graphql/
