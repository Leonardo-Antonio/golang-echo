# Commands for generate certificates

- private certificate:

    ```shell
    openssl genrsa -out app.rsa 1024
    ```

- public certificate:
    ```shell
    openssl rsa -in app.rsa -pubout > app.rsa.pub
    ```