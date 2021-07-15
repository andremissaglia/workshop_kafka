# Aggregator

Este projeto irá calcular em tempo real a média de notas para cada gato.

Input:
```
Tópico: votes
Mensagem:
  Chave: ID do gato (string, ex: "4")
  Valor: Voto de 1 a 5 (string, ex: "3")
```

Output:
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
docker-compose up -d aggregator
```
Obs: Se o comando acima falhar, provavelmente as dependências demoraram para subir. Tente rodar novamente.