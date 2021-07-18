# Backend & Replication

Este projeto é um monorepositório com dois códigos em Go: o backend (servindo o usuário diretamente), e o worker de replicação

## Backend

O backend serve a aplicação principal diretamente. Possui dois endpoints:

* `/listCats`: lista todos os gatos e informações associadas, buscando informações no banco de dados
* `/vote`: registra um voto no kafka

## Replication

Este é um worker que irá ler da fila de médias, e salvar no banco de dados.

Input:
```
Tópico: ratings
Mensagem:
  Chave: ID do gato (string, ex: "4")
  Valor: Média de 1 a 5 (string, ex: "3.4")
```


## Rodando localmente (Opcional)

Recomenda-se o uso de IDE. Mas é possível rodar pelo terminal, usando:

```
#Linux:
./gradlew :run

#Windows:
gradlew.bat :run 
```

## Rodando com o docker

```bash
docker-compose up -d backend
```
Obs: Se o comando acima falhar, provavelmente as dependências demoraram para subir. Tente rodar novamente.

Obs2: Lembre-se sempre de reiniciar o container quando fizer alterações:
```bash
docker-compose restart backend
```