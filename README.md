# graphql-learning using gqlgen

- example

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
