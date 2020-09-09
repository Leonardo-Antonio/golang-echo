# Commands for generate certificates

- private certificate:

    ```shell
    openssl genrsa -out app.rsa 1024
    ```

- public certificate:
    ```shell
    openssl rsa -in app.rsa -pubout > app.rsa.pub
    ```

# Create database

```shell
> mysql -u (user) -p(password)
> sourse model/database/mysql.sql
```

# Run application

1. 
    ```go
    go run main.go
    ```

2. 
    ```shell
    ./run
    ```