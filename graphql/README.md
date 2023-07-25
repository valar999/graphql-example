# github.com/graphql-go/graphql Example

Insert transaction:

```sh
mongosh test --eval 'db.transactions.insertMany([{ "amount": 100.50, "date": new Date() }, { "amount": 200.25, "date": new Date() }])'
```

Run the server and request it with:

```sh
http POST localhost:8080/graphql query='{list{_id, amount, date}}'
```
