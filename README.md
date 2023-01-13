# NFE BOT API

## Commands

- Run api
  - `go run ./main.go`

## Request example

- **POST** /
  - Headers
    - Content-Type: application/json
  - Body
    - ```json
      {
        "employee": {
          "login": "",
          "password": ""
        },
        "company": {
          "document_number": "",
          "name": "",
          "zip_code": "",
          "street": "",
          "district": "",
          "city": "",
          "state": "",
          "email": ""
        },
        "nfe": {
          "code": "1",
          "description": "COMISSIONAMENTO",
          "quantity": 1,
          "unit_value": 10
        }
      }
      ```
