# Deepfit Backend Api

Deepfit is a fitness application.

## Tech Stack

This project created with `Golang` and `MongoDB` and try to follow `Domain Driven Design` pattern.

## Installation

Use the go default package manager to install dependencies or use docker command to run application at docker.

```bash
go mod download
```

## Folder Structure
`Golang Standart Project Layout` followed in this project.

### `/internal`
Domain of the application is located in this folder for example repositories, entities, services etc.

### `/tools`
All 3th party tools are located in this folder.

### `/pkg`
Pkg folder is a helper folder that have functions which can be imported everywhere of the application.

### `/cmd`
The main.go file is located in this folder.

### `/scripts`
All .sh files to run, test, dockerize application is located in this folder.

### `/api`
Postman collection for this project is located in this folder.

### `/configs`
All env variables and application configs are located in this folder.

## License
[MIT](https://choosealicense.com/licenses/mit/)
