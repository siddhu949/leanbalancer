cat > README.md << 'EOF'
# 🌀 LeanBalancer - Memory Efficient Load Balancer in Go

LeanBalancer is a blazing-fast, memory-optimized load balancer built in **Go**, supporting both **forward and reverse proxying**, with integrated **health checks**, **firewall rules**, **admin APIs**, and **Prometheus metrics** — all packed into a clean and modular architecture.

---

## ✅ Features

- 🔁 Round-Robin Load Balancing  
- 🔄 Reverse & Forward Proxy Support  
- 🔥 In-Memory Firewall with Rate Limiting  
- 🩺 Automatic Health Checks for Backends  
- 📊 Prometheus Metrics Integration  
- 🧱 Admin API Panel using Fiber  
- 🚀 High-Performance with fasthttp  
- ⚙️ Graceful Shutdown Support  

---

## 🧰 Technologies Used

| Component     | Purpose                      |
|---------------|------------------------------|
| **Go**        | Main language                |
| **fasthttp**  | Lightweight HTTP server/client |
| **Fiber**     | Admin REST API framework     |
| **Zap**       | Structured logging           |
| **Prometheus**| Metrics and Monitoring       |

---

## 📦 Project Structure

leanbalancer/
├── cmd/
│   └── main.go                 # Entry point of LeanBalancer  
├── internal/
│   ├── proxy/                  # Reverse & Forward Proxy Handlers  
│   ├── firewall/               # Rate-limiting firewall  
│   ├── health/                 # Backend health checker  
│   ├── metrics/                # Prometheus metric registration  
│   ├── algorithm/              # Round-robin load balancer logic  
│   ├── logger/                 # zap logger setup  
│   ├── pool/                   # HTTP client memory pool  
│   └── admin/                  # Admin API routes  
├── api/v1/                     # REST API routing  
├── utils/                      # Helper functions  
├── examples/backend.go         # Sample backend server  
└── README.md

---

## 🚀 Getting Started

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/siddhu949/leanbalancer.git
cd leanbalancer
```
2️⃣ Install Dependencies
```bash
go mod tidy
```
3️⃣ Start Sample Backends
```bash

BACKEND_ID=1 BACKEND_PORT=9001 go run examples/backend.go
BACKEND_ID=2 BACKEND_PORT=9002 go run examples/backend.go
BACKEND_ID=3 BACKEND_PORT=9003 go run examples/backend.go
```
4️⃣ Start LeanBalancer
```bash
go run cmd/main.go
```
🌐 Available Routes
Route	Method	Description
```
/	GET	Load balancer health/status
/reverse	GET	Reverse proxy to backends
/forward	GET	Forward proxy to external site
/health	GET	LB health check
/metrics	GET	Prometheus metrics
/admin/...	GET	Admin API (via Fiber)
```
🧩 Module-wise Explanation
🔁 Proxy Module
Handles reverse and forward proxy logic.

ReverseProxyHandler sends request to backend.

ForwardProxyHandler makes external web requests.

go
```
proxy.ReverseProxyHandler(ctx)
proxy.ForwardProxyHandler(ctx)
```
🔄 Algorithm Module
Implements round-robin backend selection:
```
go

healthy := healthChecker.GetHealthyBackends()
backend := healthy[index % len(healthy)]
```
🔥 Firewall Module
Prevents abuse via IP-based rate limiting using sync.Map.

🩺 Health Checker
Periodically checks all backends using /health endpoint.
Only healthy servers are used for routing.

📊 Metrics Module
Exposes backend call counts using Prometheus:

```
backend_requests_total{backend_id="1", path="/reverse"} 24
```
🧱 Admin API
Fiber server on port 9090, modular routes defined in:

api/v1/

internal/admin/

🚀 HTTP Client Pool
Uses Go's sync.Pool for fasthttp client reuse:
```

var clientPool = sync.Pool{
  New: func() interface{} {
    return &fasthttp.Client{}
  },
}
```
📜 Logger (Zap)
Structured logs with levels like Info, Warn, Error.

```
log.Info("Started", zap.String("service", "LeanBalancer"))
```
📊 Prometheus Setup
Sample prometheus.yml:

yaml
```
scrape_configs:
  - job_name: 'leanbalancer'
    static_configs:
      - targets: ['localhost:8080']
```
Access Prometheus at:
```
👉 http://localhost:9090
```

🧪 Sample Backend Setup
Run this for mock servers:

bash
```
BACKEND_ID=1 BACKEND_PORT=9001 go run examples/backend.go
Responds to:

/

/reverse

/health
```
🧼 Graceful Shutdown
Handles OS signals, cleans up:
```

os.Signal, context.WithTimeout, server.Shutdown()
No dropped connections during exit.
```
🔮 Coming Soon
🤖 ML-based Backend Scoring

🔐 Secured Admin Dashboard

⚙️ Docker & Docker Compose Support

🌐 TLS (HTTPS) Reverse Proxy

📉 Real-time Rate Monitor

🧾 License

MIT License © 2025 @siddhu949


🤝 Contribute
Pull requests and suggestions are welcome!
Feel free to fork, improve, or open issues.
```
grafana:
grafana>windows>program files>downloads>grafana>bin>grafana.exe
prometheus:
download zip
bin .\prometheus.exe --config.file=prometheus.yml
apache bench marks:zip
.\ab.exe -n 100 -c 10 http://localhost:8080/reverse
