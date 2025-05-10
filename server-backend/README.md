# Bloody Moon - Backend

## ⚙️ Project Structure
Architecture MicroServices,
Api Gateway

Services:
1. Auth [Service](./Services/Auth/README.md)
2. Matchmaker [Service](./Services/Matchmaker/README.md)
3. Game Logic [Service](./Services/Gamelogic/README.md)

## 🔨 Technologies

Language / Frameworks
- Gin-gonic (golang)

Concepts
- Rest API
- WebSockets

DataBases
- MongoDB -> Users Data | Game Results
- Redis -> Cache | Game Stats

## ⚡ Quick Start

For better deployment, I use Docker.
[Install Docker](https://www.docker.com/get-started/)

Create [Image](https://www.geeksforgeeks.org/difference-between-docker-image-and-container/) :
```sh
docker build . -t app-gingonic
```

Create and Run Container :
```sh
docker run -d -p 8080:8080 --name running-app-gin app-gingonic
```