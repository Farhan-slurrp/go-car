# Car Listing

A microservices architecture for managing car listings and user information.

## Table of Contents

- [Repository Structure](#repository-structure)
- [API Documentation](#api-documentation)
- [Diagrams](#diagrams)

## Repository Structure

- `authentication/`: Directory for the authnentication related functions.
- `cmd/`: Main entry to compose different packages.
- `database/`: Directory for database related functions.
- `images/`: Dockerfile(s) folder.
- `internal/`: Directory for shared struct and configs.
- `pkg/`: Place to organize the services.
- `pkg/endpoints/`: Storing endpoints, request, and response types of the service
- `pkg/transport/`: Specifying service transport(s) like HTTP, gRPC
- `pkg/service`: Service request(s) handlers
- `vendor/`:  Directory to organize and manage third-party dependencies.

## API Documentation

### Car Listing Service

#### `GET /cars`

Retrieve the list of all car listings

#### `POST /cars`
Content-Type: application/json

Retrieve the list of car listings based on keys values in the filter

#### Request: 
```
{
  "filters": [
    { :key": "car_model", "value": "Honda" },
    { :key": "daily_price", "value": "100" },
    { :key": "available_from", "value": "2024-02-03T07:14:37.216Z" },
    { :key": "available_to", "value": "2024-02-06T07:14:37.216Z" },
  ]
}
```

#### Response:
```
{
  "cars": [
    {
      "ID": 1,
      "CreatedAt": "2024-02-03T09:43:25.863233+08:00",
      "UpdatedAt": "2024-02-03T09:43:25.863233+08:00",
      "DeletedAt": null,
      "car_model": "Bezza",
      "daily_price": 100,
      "available_from": "2024-01-30T22:35:39.982+08:00",
      "available_to": "2024-01-30T22:35:39.982+08:00",
      "owner_id": 3
    },
    {
      "ID": 2,
      "CreatedAt": "2024-02-03T14:41:38.869107+08:00",
      "UpdatedAt": "2024-02-03T14:44:32.904682+08:00",
      "DeletedAt": null,
      "car_model": "Honda",
      "daily_price": 200,
      "available_from": "2024-01-30T22:35:39.982+08:00",
      "available_to": "2024-01-30T22:35:39.982+08:00",
      "owner_id": 3
    }
  ]
}
```

#### `POST /cars/create`

Create new car listing and return the newly created listing ID

#### Request:
Content-Type: application/json

Header:
- Authorization: Bearer Token

```
{
    "car_listing": {
        "car_model": "Honda",
        "daily_price": 200,
        "available_from": "2024-01-30T22:35:39.982+08:00",
        "available_to": "2024-01-30T22:35:39.982+08:00"
    }
}
```

#### Response:
```
{
  "id": 1
}
```

#### `PUT or PATCH /cars/:id/update`

Update car listing info and return message

#### Request:
Content-Type: application/json

Header:
- Authorization: Bearer Token

```
{
    "car_listing": {
        "car_model": "Honda",
        "daily_price": 200,
        "available_from": "2024-01-30T22:35:39.982+08:00",
        "available_to": "2024-01-30T22:35:39.982+08:00"
    }
}
```

#### Response:
```
{
  "message": "Data updated"
}
```

### User Service

#### `GET /login`

Login with google account


#### `GET /user/:id`

Get User Data for the given id

#### Response:
```
{
  "user": {
    "ID": 3,
    "CreatedAt": "2024-02-02T21:55:19.488324+08:00",
    "UpdatedAt": "2024-02-03T08:46:41.703916+08:00",
    "DeletedAt": null,
    "name": "Farhan Ahmad Nurzi",
    "email": "farhanlmntrix@gmail.com",
    "role": "host"
  }
}
```

#### `PUT or PATCH /user/:id/update`
Content-Type: application/json

Update User Data for the given id

#### Request:
```
{
  "user": {
    "name": "Farhan Ahmad Nurzi",
    "email": "farhanlmntrix@gmail.com",
    "role": "host"
  }
}
```

#### Response:
```
{
  "message": "Data updated"
}
```

#### `GET /authorize/:token`

Get the user info for the given token

#### Response:
```
{
  "user": {
    "ID": 3,
    "CreatedAt": "2024-02-02T21:55:19.488324+08:00",
    "UpdatedAt": "2024-02-03T08:46:41.703916+08:00",
    "DeletedAt": null,
    "name": "Farhan Ahmad Nurzi",
    "email": "farhanlmntrix@gmail.com",
    "role": "host"
  }
}
```

## Diagrams

### Car Listing Table
![image](https://github.com/Farhan-slurrp/go-car/assets/58872254/42175c1b-8bbe-4600-ab53-8d51a18ca70e)

### User Table
![image](https://github.com/Farhan-slurrp/go-car/assets/58872254/2223b19f-bfbc-47c2-ad81-31996b8c6422)
