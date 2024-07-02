# TODO API with Golang, Gin, and ScyllaDB

This project is a TODO API built using Golang, the Gin web framework, and ScyllaDB as the database. It supports basic CRUD operations, pagination, and filtering based on the TODO item status.

## Features
* User Signup: Users can sign up using an email, username, and password.
* Create TODO: Create a TODO item by providing a title and description.
* Update TODO: Update the details of a TODO item.
* Delete TODO: Delete a TODO item.
* Mark TODO as Completed: Mark a TODO item as completed.
* Get TODO by ID: Retrieve a TODO item by its ID.
* List TODOs: List all TODO items for a user with optional filters for status, limit, and offset for pagination.

## Getting Started

### Prerequisites
+ Golang
+ Docker
+ ScyllaDB

## Installation
  Clone the repository:
  
        git clone https://github.com/anjush-bhargavan/todo-api.git

  Install dependencies:
  
        go mod tidy
            
Set up ScyllaDB using Docker:

        docker run --name some-scylla -d scylladb/scylla

Run the application:

        go run main.go


## API Endpoints
- Signup User: POST /api/v1/signup
- Login User: POST /api/v1/login
- Create TODO: POST /api/v1/user/todos
- Get TODO by ID: GET /api/v1/user/todos/:id
- Mark as completed TODO: PATCH /api/v1/user/todos/:id
- Update TODO: PUT /api/v1/user/todos
- Delete TODO: DELETE /api/v1/user/todos/:id
- List TODOs: GET /api/v1/user/todos

  
## Postman Documentation
 You can find the Postman documentation for this API [here](https://documenter.getpostman.com/view/30219361/2sA3dvmD4i).

