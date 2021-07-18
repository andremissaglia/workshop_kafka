# Workshop kafka

Esta é uma aplicação para avaliação de gatos em tempo real. No frontend é exibido uma lista de gatos, onde o usuário pode dar uma nota de 1 a 5. A aplicação calcula em tempo real

Existem duas branches neste repositório:

 * `master`: O código está incompleto, com alguns TODOs. Útil para tentar acompanhar o workshop e fazer o exercício por conta.
 * `complete`: O projeto é funcional. Útil para servir de guia

## Rodando localmente (Opcional)

Procure o README de cada subsistema para verificar como rodar.

## Rodando com o docker

É preferível rodar a aplicação com o docker. Para isso, execute:

```
docker-compose up -d
```
Obs: Se o comando acima falhar, provavelmente as dependências demoraram para subir. Tente rodar novamente.

Obs2: Lembre-se sempre de reiniciar o container quando fizer alterações:
```bash
docker-compose restart backend
```

Para parar a aplicação:
```
docker-compose down
```

Este projeto usa docker volumes para persistência. Para limpar a base e recomeçar do zero:
```
docker-compose down -v
```
