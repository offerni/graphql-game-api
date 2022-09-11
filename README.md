# GraphQL Learning

Project spinning up concurrently an HTTP server with [echo](https://echo.labstack.com/) and a GraphQL server using [gqlgen](https://gqlgen.com/)

`go run github.com/99designs/gqlgen generate` generates schema resolvers and handlers based on `schema.graphqls`

WIP, only hardcoded values for now

Open `http://localhost:8080/` to use the REST api,

- Examples:

  - REST Endpoints:

    - GET

      ```
        http://localhost:8080/stores/steam
        http://localhost:8080/games/1
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

  - GraphQL:

    Open the GraphQL playground on `http://localhost:8081/` and use the query below

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
