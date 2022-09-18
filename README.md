## Proof-id backend 

## Run
```go
go run main.go
```

### Endpoints
```
/generateProof
```
```JSON
curl -i http://localhost:8080/generateProof \
-H 'Accept: application/json' \
-d '{"price":"500","times":"5","salary":"4000"}'
```

```
/verifyProof
```

```JSON
curl --location --request GET '127.0.0.1:8080/verifyProof' \
--header 'Content-Type: application/json' \
--data-raw '{"tokenid":"1"}'
```