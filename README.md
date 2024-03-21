# giphy-challenge

## Requisitos
Herramientas necesarias para la ejecucón de la aplicación:
- docker
- docker-compose

ó
- node 19.5.0
- go 1.18

## Poblar datos
Para poblar datos de prueba en el backend es necesario cambiar la variable `POPULATE` y agregar una `api_key` de Giphy valida en el archivo `backend/.env`, quedando:

```js
GIPHY_API_KEY=<example>
POPULATE=true
```

## Ejecución
### Backend
Cualquiera sea el método para levantar el proyecto se puede probar en `http://localhost:8080/playground`.

- Usando docker-compose:
```sh
cd backend
make docker-run-app
```

- Con go:
```sh
cd backend
go mod tidy
go run main.go
```

> Para este caso es necesario modificar las credenciales de MongoDB en el archivo `backend/.env`, quedando:
```
MONGO_URL=mongodb://exampleuser:examplepassword@database:27017/?authSource=admin&readPreference=primary&ssl=false&directConnection=true
MONGO_GIF_DATABASE=db
MONGO_GIF_COLLECTION=gifs
```

### Frontend
Cualquiera sea el método para levantar el proyecto se puede probar en `http://localhost:3000/`.

- Usando docker-compose:
```sh
cd frontend
docker-compose up
```

- Con yarn:
```sh
cd frontend
yarn add
yarn dev
```
