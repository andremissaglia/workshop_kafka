# Frontend

## Rodando localmente (Opcional)

```bash
npm install
ng serve --open
```
O backend precisa estar funcionando no endereço `localhost:8080`.

## Rodando com o docker

* Crie uma distribuição do frontend. (Opcional: Uma versão já está commitada no repositório.)

```bash
ng build
```

* Levante toda as imagens

```bash
docker-compose up -d frontend
```
Obs: Se o comando acima falhar, provavelmente as dependências demoraram para subir. Tente rodar novamente.
