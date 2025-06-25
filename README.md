
# Multi-Service Docker Application with NGINX Reverse Proxy

This project demonstrates a simple microservices architecture using Docker, Docker Compose, and NGINX as a reverse proxy. It includes:

- A **Go** application (Service 1)
- A **Python Flask** application (Service 2)
- An **NGINX** reverse proxy to route traffic between services
- A **Docker Compose** setup to orchestrate everything

---

## 📁 Project Structure

```
Assignment/
│
├── docker-compose.yml            # Compose file to run all services
├── nginx/
│   └── nginx.conf                # NGINX reverse proxy configuration
│
├── service_1/                    # Go service
│   └── main.go
│   └── Dockerfile
│
└── service_2/                    # Python Flask service
    └── app.py
    └── requirements.txt
    └── Dockerfile
```

---

## 🚀 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop) installed
- [Docker Compose](https://docs.docker.com/compose/) (v2+)

---

## 🏗️ How to Run the Project

1. **Navigate to the project folder:**
```bash
cd Assignment
```

2. **Run the entire stack using Docker Compose:**
```bash
docker-compose up --build
```

> This command builds the Docker images and starts all three services: Go app, Python app, and NGINX.

### ✅ After running the above command:

- Visit `http://localhost:8080/service1` → You will see the **Go application**.
- Visit `http://localhost:8080/service2` → You will see the **Python Flask application**.

---

## 🌐 Accessing the Services

- **Go Service (Service 1):**
  - `http://localhost:8080/service1/ping` → returns `"pong"`
  - `http://localhost:8080/service1/hello` → returns `"Hello from Go!"`

- **Python Flask Service (Service 2):**
  - `http://localhost:8080/service2/` → returns `"Hello from Python Service!"`

---

## 🐳 Docker Compose Explained

Here’s what happens when you run `docker-compose up`:

- `service1`: Runs a Go app on port `8001`
- `service2`: Runs a Flask app on port `8002`
- `nginx`: Listens on port `8080` and proxies requests to the appropriate backend

---

## 🔀 NGINX Reverse Proxy Routing

Inside `nginx/nginx.conf`, the reverse proxy handles routing like:

```nginx
location /service1/ {
    proxy_pass http://service1:8001/;
}

location /service2/ {
    proxy_pass http://service2:8002/;
}
```

---

## 📦 Python Dependencies

Python service requirements are defined in:

```
service_2/requirements.txt
```

Installed automatically during image build.

---

## 🧹 Clean Up

To stop and remove all running containers:

```bash
docker-compose down
```

---

## 🛠️ Troubleshooting

- If a service fails to start, ensure no other apps are using ports 8001, 8002, or 8080.
- Use `docker-compose logs` to view logs and debug issues.

---

## 📄 License

This project is for educational purposes only.

---

## 👨‍💻 Author

Sai Kiran
