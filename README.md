# To generate api from proto-files:

```shell
make generate api
```

# To build the binary file of the project:

```shell
make build
```

# To run the whole command that necessary to run:

```shell
make run
```


# To create the migrations

```shell
migrate create -ext sql -dir db/migrations -seq {NAME}
```

