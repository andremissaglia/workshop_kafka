# ksqlDB

Com o ksql, podemos trabalhar mais facilmente com streaming no kafka, usando uma linguagem de alto nível, muito próxima ao KSQL

## Rodando com o docker

```bash
# Iniciar o servidor
docker-compose up -d ksqldb-server

# Rodar o cli
docker-compose exec ksqldb-cli ksql http://ksqldb-server:8088
```
Obs: Se o comando acima falhar, provavelmente as dependências demoraram para subir. Tente rodar novamente.