DBDiagram: https://dbdiagram.io/d/66f86a173430cb846cefa416

## Golang Migrate
- Create initial seq schema
```sh
    migrate create -ext sql -dir db/migration -seq init_schema
```

Transaction Goals:
- Create a _transfer_ entry with amount 'X'
- Create an _entry_ in from_account with amount '-X'
- Create an _entry_ in to_account with amount 'X'
- Subtract 'X' from the balance of from_account
- Add 'X' to the balance of to_account

## References
> [select for update in postgresql](https://haril.dev/en/blog/2024/04/20/select-for-update-in-PostgreSQL)