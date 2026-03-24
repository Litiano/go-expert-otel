### Instruções

#### 1 - Copie o arquivo [config.env.example](temperature-service/config.env.example) para [config.env](temperature-service/config.env).
#### 2 - Adicione uma chave de API do serviço [Weather](https://www.weatherapi.com/).
#### 3 - Execute `docker compose up -d`
#### 4 - Realize uma chama na API:
```shell
curl --location 'http://localhost:8080/temperature' \
    --header 'Content-Type: application/json' \
    --data '{"cep": "12345678"}'
```
#### 5 - Visualize o trace em `http://localhost:9411/zipkin/`
Run Query
