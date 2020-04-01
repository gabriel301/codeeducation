# Desenvolvimento de Aplicações Moderas e Escaláveis com Microsserviços

Repositório dos entregáveis do curso _**Desenvolvimento de Aplicações Modernas e Escaláveis com Microsserviços**_ da [**Code.Education**](https://code.education/).

Esta descrição será atualizada no decorrer do curso.
## Dependências

- Composer 1.10.1
- Docker 19.03.8
- Docker-Compose 1.17.1
- Laravel Framework 7.3.0
- MySql 5.7
- Nginx 1.15.0
- PHP 7.3.6
- Redis 5.0.8

## Lista de Projetos Práticos

- DevOps
  - [Publicando imagem Laravel](#publicando-imagem-laravel)

## Projetos Práticos

### 1. Módulo DevOps

#### Publicando Imagem Laravel

   Para construir a imagem, executar o seguinte comando **no diretório raíz do repositório**:\
  ```docker build -t <nome da imagem> -f ./.docker/laravel/dockerfile.standalone .```

  Imagem Gerada: [Imagem](https://hub.docker.com/repository/docker/gabriel301/laravel)
