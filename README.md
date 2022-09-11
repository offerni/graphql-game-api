# graphqllearning using gqlgen

WIP, only hardcoded values for now

- example
  Open `http://localhost:8080/` to use the REST api,
  -- examples:

  ```
    http://localhost:8080/stores/asdfg
    http://localhost:8080/stores/qwerty
  ```

  Open the GraphQL playground on `http://localhost:8081/` and use the query below

```
query {
  store(id: "asdfg") {
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
