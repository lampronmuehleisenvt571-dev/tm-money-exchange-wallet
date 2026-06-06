# Getting Started - TM Money Exchange Wallet

## 5-Minute Quick Start

### Prerequisites
- Docker & Docker Compose
- Git
- Minimum 4GB RAM, 10GB disk space

### Step 1: Clone Repository

```bash
git clone https://github.com/lampronmuehleisenvt571-dev/tm-money-exchange-wallet.git
cd tm-money-exchange-wallet
```

### Step 2: Setup Environment

```bash
cp .env.example .env
```

Edit `.env` with your local configuration (optional for development).

### Step 3: Start Development Environment

```bash
make dev-up
```

This starts:
- PostgreSQL database
- Redis cache
- MongoDB document store
- Elasticsearch search engine
- Kafka message broker
- RabbitMQ message queue
- Grafana monitoring
- Prometheus metrics
- Jaeger tracing

### Step 4: Verify Services

```bash
# Check running containers
docker-compose ps

# View application logs
make dev-logs
```

### Step 5: Access Dashboards

- **Grafana:** http://localhost:3000 (admin/admin)
- **Prometheus:** http://localhost:9090
- **Jaeger:** http://localhost:6831
- **Kibana:** http://localhost:5601

---

## Local Development Workflow

### Running Tests

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run specific test
go test -v ./src/services/wallet-service/...
```

### Code Quality

```bash
# Run linters
make lint

# Format code
make format
```

### Database Operations

```bash
# Run migrations
make db-migrate

# Seed test data
make db-seed

# Reset database (CAREFUL!)
make db-reset
```

### Building & Running

```bash
# Build binary
go build -o bin/tm-money ./src/main.go

# Run application
make dev-run

# Or with Docker
docker build -t tm-money:latest .
docker run -p 8080:8080 tm-money:latest
```

---

## Project Structure Navigation

### Source Code
```
src/
├── services/            # 21 microservices
├── shared/              # Shared libraries & utilities
├── config/              # Configuration management
└── main.go             # Application entry point
```

### Configuration
```
deployment/
├── docker/              # Docker files
├── kubernetes/          # K8s manifests
├── terraform/           # Infrastructure as Code
└── helm/               # Helm charts
```

### Documentation
```
docs/
├── TDD/                 # 33-section technical docs
├── ARCHITECTURE/        # System design docs
├── API/                # API specifications
└── GUIDES/             # How-to guides
```

---

## Useful Commands

```bash
# View all available commands
make help

# Start development environment
make dev-up

# Stop development environment
make dev-down

# Run application
make dev-run

# Run tests
make test

# Code quality checks
make lint
make format

# Database
make db-migrate
make db-seed

# Cleanup
make clean
```

---

## Troubleshooting

### Port Already in Use

```bash
# Find process using port 8080
lsof -i :8080

# Kill process
kill -9 <PID>
```

### Docker Compose Issues

```bash
# Remove all containers and volumes
make clean-docker

# Restart from scratch
make dev-up
```

### Database Connection Issues

```bash
# Check PostgreSQL logs
docker-compose logs postgres

# Reset database
make db-reset
```

---

## Next Steps

1. **Read Documentation:** Start with [Technical Design Document](../TDD/00-COMPLETE-TABLE-OF-CONTENTS.md)
2. **Explore APIs:** Check [API Documentation](../API/)
3. **Review Code:** Look at example services in `src/services/`
4. **Write Code:** Create your first feature following project conventions
5. **Submit PR:** Follow contribution guidelines in CONTRIBUTING.md

---

## Support

- **Documentation:** [docs/](../)
- **Issues:** [GitHub Issues](https://github.com/lampronmuehleisenvt571-dev/tm-money-exchange-wallet/issues)
- **Email:** architecture@tmmoney.io

---

**Happy coding! 🚀**
