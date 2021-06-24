# Budreck

An API that will track all your transactions, basically a budget tracker.
The Budreck app itself will help you to be less budrek on managing your finance.

#### How to run?

Make sure you have these following depedencies:
- postgresql
- make
- go

Once you have that, create a database.

```
createdb yourdbname
```

Then modify `config/environment.go`, adjust it with your own configuration.

```
DATABASE_USERNAME string = "root"
DATABASE_PASSWORD string = nil
DATABASE_NAME     string = "budreck"
DATABASE_HOST     string = "localhost"
DATABASE_PORT     int32  = 5432
```

You've set up, now you can run `make install` && `make start`

#### Migrating Database Schema

To migrate the database, you can use `golang-migrate` CLI, but it is easier to just run:

```sh
go run main.go db:migrate

# Or if you have the compiled version.
bin/budrack db:migrate
```

To downgrade the database version by one version, you can run:

```sh
go run main.go db:rollback

# Or if you have the compiled version.
bin/budrack db:rollback
```

#### Generating Migration

To generate migration file, all you need to run is just:

```sh
go run main.go generate migration your_migration_name_in_snake_case

# Or if you have the compiled version.
bin/budrack migrate your_migration_name_in_snake_case

```

## LICENSE

MIT
