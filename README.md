# Instruções para Configuração do Ambiente

## Passo 1: Construir e Iniciar os Contêineres

Execute o comando abaixo para construir e iniciar os contêineres usando o Docker Compose:

```sh
docker-compose up -d --build
```


## Passo 2: Executar Script de Inicialização no Contêiner PostgreSQL

Após os contêineres estarem em execução, execute o seguinte comando para rodar o script de inicialização no contêiner PostgreSQL:

```sh
docker exec -it postgres /docker-entrypoint-initdb.d/at_startup/init.sh
```

```sh
google-chrome --disable-web-security --user-data-dir="/tmp/chrome_dev"
google-chrome --disable-web-security --user-data-dir="/tmp/chrome_dev"
google-chrome --disable-web-security --user-data-dir="/tmp/chrome_dev"
```# bem-te-vi
# bem-te-vi
