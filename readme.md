**Create certificates:**

Public key 
```
 openssl genrsa -out app.rsa 1024
 ```
Private key
```
 openssl rsa -in app.rsa -pubout > app.rsa.pub
```


**Up server:**
```
  go run cmd/main.go
``` 