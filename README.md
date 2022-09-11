# GraphQL Learning

Project spinning up concurrently an HTTP server with [echo](https://echo.labstack.com/) and a GraphQL server using [gqlgen](https://gqlgen.com/)

To run the project run
```
  go run cmd/main.go
```
Or just `air` if you have it installed


`go run github.com/99designs/gqlgen generate` generates schema resolvers and handlers based on `schema.graphqls`

Open `http://localhost:8080/` to use the REST api,

- Examples:

  ##  REST:
  
  - Stores

    - GET

      ```
        http://localhost:8080/stores/steam
      ```

    - POST

      ```
        http://localhost:8080/stores
      ```

      - Body

      ```
        {
            "name": "Cool Store"
        }
      ```
  - Games

      - GET

        ```
          http://localhost:8080/games/1
        ```

      - POST

        ```
          http://localhost:8080/games
        ```

        - Body
        
        ```
        {
          "name": "The last of us",
          "store_id": "steam",
          "price": "25.98"
        }
        ```

  ## GraphQL:

    Open `http://localhost:8081/` and use the GraphQL playground

    ```
    query {
      store(id: "steam") {
        id,
        name,
        games{
          id,
          store_id
          name,
          price
        }
      }
    }
    ```

    ```
    mutation CreateStore {
      CreateStore(opts: {name: "xbox store"}) {
        id,
        name,
      }
    }
    ```

    ```
    mutation CreateGame {
      CreateGame(opts: {
        name: "The last of us part II"
        store_id: "steam",
        price: "58.66"
      }) {
        id,
        name,
        store_id,
        price
      }
    }
    ```
