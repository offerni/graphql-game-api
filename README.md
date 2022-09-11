# graphqllearning using gqlgen

WIP, only hardcoded values for now

Open `http://localhost:8080/` to use the REST api,

- examples:
- REST:

```
  http://localhost:8080/stores/steam
  http://localhost:8080/stores/1
```

- GraphQL:

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
