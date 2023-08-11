# Translation Playground

A playground application to experiment with the [go-i18n framework](https://github.com/nicksnyder/go-i18n/tree/main).

Usage:

Run the application:
```sh
go run cmd.go
```

In a separate terminal run either of the following commands:

* To receive English responses:
```sh
curl "localhost:3000/?name=your_name" | jq .
```

* To receive Spanish responses:
```sh
curl "localhost:3000/?name=your_name&lang=es" | jq .
```