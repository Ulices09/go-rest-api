# go-rest-api

## ORM 

Install the ent cli to create new schemas

```
go get entgo.io/ent/cmd/ent
```

Create a new schema with the following command 

```
go run entgo.io/ent/cmd/ent init <SchemaName>
```

After adding the corresponding fields, edges, etc. Generate the code with the following command

```
make db-gen
```

## API Documentation

Install swagger cli

```
go install github.com/swaggo/swag/cmd/swag@latest
```

More information [here](https://github.com/swaggo/swag)

## References

- [Domain Driven Hexagon](https://awesomeopensource.com/project/Sairyss/domain-driven-hexagon)

- [Domain Driven Hexagon Repository](https://github.com/Sairyss/domain-driven-hexagon)

- [DDD, Hexagonal, Onion, Clean, CQRS, … How I put it all together](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/)

- [go-structure-examples](https://sourcegraph.com/github.com/katzien/go-structure-examples@master)