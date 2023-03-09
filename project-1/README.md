# Go Intensivo 2 - Order consumer

> Evento Go Intensivo 2 FullCycle

## Run project

### Dependencies:

- Docker
- Docker Compose
- Internet access

### To run:

- terminal: docker compose up -d

- terminal: docker compose exec goapp bash

- terminal: go mod tidy

- browser:

        access: http://localhost:9021
        click: controlcenter.cluster
        click on menu: Topics
        click on button: Create topic
        fill on form: Topic.name: orders
        click on button: Create with defaults
        click tab: Messages

- browser new tab:

        access: http://localhost:15672
        fill login form: Username: guest and Password: guest
        click on button: Login
        click on tab: Queues
        click: Add a new queue
        fill on form: Name: orders
        click on button: Add queue
        click on table row: orders

- terminal: go run cmd/consumer/main.go

- browser:

        access:http://localhost:9021
        click: controlcenter.cluster
        click on menu: Topics
        click on table row: orders
        click tab: Messages
        click: + Produces a new message to this topic
        in Value insert this json:
          {
            "ID": "kafka123",
            "Price": 10,
            "Tax": 1
          }
        click on button: Produce

- browser new tab:

        access: http://localhost:15672
        if show login form:
          fill login form: Username: guest and Password: guest
          click on button: Login
        click on tab: Queues
        click on table row: orders
        click: Publish message
        in Payload insert this json:
          {
            "ID": "rabbitmq123",
            "Price": 10,
            "Tax": 1
          }
        click on button: Publish message

### To stop

- terminal: ctrl + c

- terminal: exit

- terminal: docker compose down
