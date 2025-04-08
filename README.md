# Memory-Efficient Load Balancer

## Overview
This project is a **memory-efficient load balancer** with forward and reverse proxy capabilities, designed to handle **heavy traffic** efficiently, similar to **AWS ELB**. The implementation is done using **Golang** with a focus on **performance, security, and scalability**. The primary focus is on **minimizing memory overhead** while maintaining high efficiency.

## Features
- **Forward & Reverse Proxy Support**
- **Optimized Memory Management**
- **High-Performance Load Distribution**
- **Security Enhancements**
- **Machine Learning-based Traffic Optimization**
- **Web UI for Monitoring & Configuration**
- **Efficient Concurrency Handling using Goroutines**

## Packages Used
This project leverages the following Golang packages for performance optimization:
- **[GoFiber](https://github.com/gofiber/fiber)** - A fast, lightweight web framework inspired by Express.js
- **[fasthttp](https://github.com/valyala/fasthttp)** - A high-performance HTTP engine
- **Goroutines** - For efficient concurrency handling and parallel processing

## Installation
```bash
# Clone the repository
git clone https://github.com/siddhu949/Load-Balancer.git
cd Load-Balancer

git config --global user.name "Your GitHub Username"
git config --global user.email "your-email@example.com"

# Install dependencies
go mod tidy

# Build the project
go build -o load-balancer

# Run the load balancer
./load-balancer
```

## Usage
Modify the configuration file to specify backend servers and balancing strategies. Example configuration:
```json
{
  "load_balancing_strategy": "round_robin",
  "servers": [
    {"host": "server1.example.com", "port": 8080},
    {"host": "server2.example.com", "port": 8081}
  ]
}
```

Start the load balancer with the config:
```bash
./load-balancer -config config.json
```

## GitHub Workflow
```bash
git pull origin main            # Get latest changes
git checkout -b feature-branch  # Create a new branch
git add .
git commit -m "Commit message"
git push origin feature-branch  # Push changes
```

## Merging Changes
```bash
git checkout main
git pull origin main
git merge feature-branch
git push origin main
```

## Cleanup
```bash
git branch -d feature-branch
git push origin --delete feature-branch
```

## Useful Git Commands
```bash
git status                     # Check status
git stash && git stash pop     # Save & restore changes
git reset --hard HEAD          # Reset changes
git push origin main -u
```

## Running Performance Tests & Monitoring
To run **Apache Bench** benchmark:
```sh
.\ab.exe -n 5000 -c 100 http://localhost:8080/
```

To run **Prometheus**:
```sh
.\prometheus.exe --config.file=prometheus.yml
```

To run **the main load balancer**:
```sh
go run ./cmd/server/main.go
```

To run **backend servers**:
```sh
go run backendX.go
```

To run **Grafana**:
```sh
.\grafana-server.exe
```

## Running with Docker (supports)
To run the load balancer using Docker:
```sh
docker run -p 8080:8080 leanbalancer
```

To run backend servers in Docker:
```sh
docker run -d -p 9001:80 --name backend1 nginx
docker run -d -p 9002:80 --name backend2 nginx
docker run -d -p 9003:80 --name backend3 nginx
docker run -d --network=mynetwork -p 8080:8080 --name leanbalancer leanbalancer
```
## dahboard :
- created in web folder
```sh
npx create-react-app .
For Material UI:

npm install @mui/material @emotion/react @emotion/styled @mui/icons-material axios

```
```sh
For Material UI:

npm install @mui/material @emotion/react @emotion/styled @mui/icons-material axios

```
```sh
For boot strap CSS:

npm install bootstrap react-bootstrap

```
```sh
npm install react-router-dom


```
```
```sh
npm install react-icons

```
```
#through api:
backend adding:
```sh
 Invoke-RestMethod -Uri "http://localhost:9090/api/v1/backends" `
>>   -Method Post `
>>   -Headers @{ "Content-Type" = "application/json" } `
>>   -Body '{ "url": "http://localhost:9003" }'
#all servers responds
```sh
 curl http://localhost:9090/api/v1
 curl http://localhost:9090/api/v1/health

```
#BLOCK THE IP IN ADMIN FIREWALL:
```sh
 Block an IP using the firewall:
 Invoke-RestMethod -Uri "http://localhost:9090/api/v1/firewall/block" `
  -Method Post `
  -Headers @{ "Content-Type" = "application/json" } `
  -Body '{ "ip": "192.168.1.100" }'
 View firewall rules:
 curl http://localhost:9090/api/v1/firewall


```
```sh
npm install react-icons

```

```
## Tips
- Pull before making changes.
- Use clear commit messages.
- Resolve conflicts carefully.

🚀 Happy coding!
