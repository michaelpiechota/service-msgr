# service-msgr

## Introduction

A simple messenger API.

## Installation

Install Go: https://golang.org/doc/install

## Features

- A simple RESTful messenger API
- HTTP Server
- Configuration file imports
- Request logging
- Response encoding
- Graceful shutdown

The following third-party packages are used:
- [github.com/spf13/viper](https://github.com/spf13/viper) (config)
- [github.com/go-chi/chi](https://github.com/go-chi/chi) (routing and middleware) (v5)
- [github.com/uber-go/zap](https://github.com/uber-go/zap) (structured logging)
- [github.com/InVisionApp/go-health](https://github.com/InVisionApp/go-health) (health checking)

## Environment Variables
#### Runtime

| name                     | description                                                     | type    | optional | default     |
|--------------------------|-----------------------------------------------------------------|---------|----------|-------------|
| PORT                     | The server port                                                 | string  | yes      | 3000        |
| LOGGER                   | The logger type (TEST, DEVELOPMENT)                             | string  | yes      | DEVELOPMENT |
| ENV_FILE                 | location of the .env file                                       | string  | yes      | .env        |
| HEALTH_CHECK_ENDPOINT    | Endpoint to send a GET requeset for the health of service       | string  | yes      | /healthz    |
| READY_CHECK_ENDPOINT     | Endpoint to send a GET requeset for the services ready check    | string  | yes      | /readyz     |

## Run

```bash
make build
./bin/serve
```

## Clear Port (if needed)
```bash
kill -9 $(lsof -i:3000 -t) 2> /dev/null
```

## Notes
* app entry point is cmd/serve/main.go
* The "database" tables and functions are mocked in /messenger/db.go
* Currently, the GET /messages/{userID} only supports returning 1 entry. See TODO section.
* Limits of 100 messages or all messages in last 30 days have not been implemented for either global GET or {userID} GET endpoints.

## TODO:
* !!!ADD TESTS!!!
* Modify GET endpoint to return a slice of messages per userid. 
* Add 30 days/100 messages requirement logic to API endpoints. Maybe within Paginate func?
* Implement queues!?
* Implement a real DB!
* Finish docker-compose for ease of use

# API (Current State)

**Get All Messages**
----
  Returns all messages from all users.

* **URL**

  /messages/search

* **Method:**

  `GET`
  
*  **URL Params**

    N/A

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `[{"id":"1","user_id":100,"message":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.","date":"2018-07-04T02:09:29-06:00","user":{"id":100,"name":"Michael","role":"messengerUser"}}]`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{"status":"Resource not found."}`

* **Sample Call:**

  ```bash
  curl http://localhost:3000/messages/search
  ```

**Send A New Message**
----
Send a new message to a specific user

* **URL**

  /messages/create/{userID}

* **Method:**

  `POST`
*  **URL Params**

    {userID} where userID is the recipient's userID
* **Data Params**

  None

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** `{"id":"91","user_id":100,"message":"quick brown fox wow","date":"0001-01-01T00:00:00Z","user":{"id":100,"name":"Michael","role":"messengerUser"}}`
 
* **Error Response:**

* **Code:** 422, 400, or 404 

* **Sample Call:**

  ```bash
  curl -X POST -d '{"id":"100","message":"QUICK BROWN FOX WOW"}' http://localhost:3000/messages/create/100
  ```

**Get Messages From Specific User**
----
  Returns a message from a specific users.

* **URL**

  /messages/{userID}

* **Method:**

  `GET`
  
*  **URL Params**

    {userID} where userID is the user whose message you would like to retrieve

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"id":"1","user_id":100,"message":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.","date":"2018-07-04T02:09:29-06:00","user":{"id":100,"name":"Michael","role":"messengerUser"}}`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{"status":"Resource not found."}`

* **Sample Call:**

  ```bash
  curl http://localhost:3000/messages/100
  ```
