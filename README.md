DBDiagram: https://dbdiagram.io/d/66f86a173430cb846cefa416

## Golang Migrate
- Create initial seq schema
```sh
    migrate create -ext sql -dir db/migration -seq init_schema
```

## Transactions
Transaction Goals
------------------------------
- Create a _transfer_ entry with amount 'X'
- Create an _entry_ in from_account with amount '-X'
- Create an _entry_ in to_account with amount 'X'
- Subtract 'X' from the balance of from_account
- Add 'X' to the balance of to_account


Transaction Isolation Levels
------------------------------

- Read Uncommitted - Not allowed in Postgres. [Read](https://www.postgresql.org/docs/current/transaction-iso.html#:~:text=PostgreSQL%27s%20Read%20Uncommitted%20mode%20behaves%20like%20Read%20Committed.%20This%20is%20because%20it%20is%20the%20only%20sensible%20way%20to%20map%20the%20standard%20isolation%20levels%20to%20PostgreSQL%27s%20multiversion%20concurrency%20control%20architecture.).
- Read Committed
- Repeatable Read
- Serializable

```sh
BEGIN;
 SET TRANSACTION ISOLATION LEVEL read committed;
 SHOW TRANSACTION ISOLATION LEVEL;
 ...

COMMIT;
```

> NB:
In postgres, the transaction isolation level can only be changed within transaction blocks. 

## References
> [select for update in postgresql](https://haril.dev/en/blog/2024/04/20/select-for-update-in-PostgreSQL)