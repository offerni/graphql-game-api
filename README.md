# graphql-learning using gqlgen

- example
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
