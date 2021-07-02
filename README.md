# Transaction Account
****Dependencies:****
- Docker and Docker Compose
- SGBD of your choice (optional)
- cURL or Postman (optional)

****Technologies:****
- GO
- MySQL

****Architecture:****
- [Hexagonal](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/ "Hexagonal" )

**Playground**:
- For use in Postman, use the collection available at the root of the project called *account-transaction.postman_collection.json*

## Setting up the application:
```shell
make configure
```
Application requirements.

## Starting the application:
```shell
make up
```
After installation and configuration the application can be accessed http://localhost:3001

## Creating the database tables:
```shell
make migrate
```
Creation of the necessary tables for the application.
**NOTE: You will be asked for the bank password on migration.**

Connection to the database:
```
  HOST=localhost
  PORT=3306
  DATABASE=foo
  USERNAME=root
  PASSWORD=dev
```

------------


### Other commands:
- **Stopping the application:**
```shell
make down
```
- **Tests:**
```shell
make tests
```
- **Application log:**
```shell
make log
```

------------


# Using the API

#### Authentication:
All requests made to the APIs must be authenticated. The API uses the Basic Authentication approach. That means you have to send the HTTP requests with the Authorization header . For example:

`Authorization:2428892a8ab08c23be9c55177a0c7713`

To test API use this authorization:

`Authorization:0c7ee5a41bff7c8af4d4ff3740b0224d`

#### Routes:
| Endpoint | Method |
| :------------ | :------------ |
| /v1/accounts | POST |
| /v1/account/:account-id | GET |
| /v1/transaction| POST |


# Account [/accounts]
### Creating an account [POST]
Body:

| Parameters | Description | Type | Validation | Mandatory |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| document_number | Document identifier | string | Check Digit | YES |

Example:
```json
{
    "document_number": "49330192076"
}
```
Response: Status: 200 OK
```json
{
    "account_id": "30e11b25-89a3-403d-af06-26a837563c55"
}
```
Request example using cURL:
```shell
curl --location --request POST 'http://localhost:3001/v1/accounts' \
--header 'Authorization: 0c7ee5a41bff7c8af4d4ff3740b0224d' \
--header 'Content-Type: application/json' \
--data-raw '{
    "document_number": "49330192076"
}'
```
### Querying an account's information [GET /:account-id]
Response: Status: 200 OK
```json
{
    "account_id": "3e0147b4-b777-4994-be9c-921361d8c06d",
    "document_number": "15548617052"
}
```

------------


# Transaction [/transaction]
### Creating a transaction [POST]
Body:

| Parameters | Description | Type | Validation | Mandatory |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| account_id | Account identifier | string | uuid v4 | YES |
| operation_type | Type of operation | integer | Min.1, Max. 4 | YES |
| amount | Transaction amount | decimal | Min 0.1 | YES |

Example:
```json
{
    "account_id": "3e0147b4-b777-4994-be9c-921361d8c06d",
    "operation_type": 4,
    "amount": 10.01
}
```
Response: Status: 201 Created
```json
{
    "transaction_id": "5cd6f148-1e26-40f8-b4fc-b17dc98b1da1"
}
```
Request example using cURl:
```shell
curl --location --request POST 'http://localhost:3001/v1/transactions' \
--header 'Authorization: 0c7ee5a41bff7c8af4d4ff3740b0224d' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id": "3e0147b4-b777-4994-be9c-921361d8c06d",
    "operation_type": 4,
    "amount": 10.01
}'
```