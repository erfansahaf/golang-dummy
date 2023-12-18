## Readme

First of all, run a MySQL server on port 3306 and create a database for this project.

After that you can run the web server (GIN) using the following command:

```shell
DB_USER='YOU USERNAME' DB_PASS='YOUR PASSWORD' DB_NAME='YOUR DATABASE' go run main.go serve
```

### Seeding the database

If you want to load 1 Million records into database, you can run:

```shell
POOL_SIZE=200 DB_USER='YOU USERNAME' DB_PASS='YOUR PASSWORD' DB_NAME='YOUR DATABASE' go run main.go seed
```

Please note that you can adjust the `POOL_SIZE` variable to speed up import process.