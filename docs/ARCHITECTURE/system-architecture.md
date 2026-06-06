# System Architecture - TM Money Exchange Wallet

## Overview

The TM Money Exchange Wallet platform employs a modern, cloud-native, microservices-based architecture designed for:
- **Scalability:** Handles 1,000,000+ concurrent users
- **Performance:** < 1ms matching latency, < 100ms API latency
- **Reliability:** 99.99% uptime SLA
- **Security:** Zero-trust architecture with AES-256 encryption
- **Multi-Cloud:** AWS, Google Cloud, Azure support

---

## 9-Layer Architecture

### Layer 1: Client Layer
- Mobile Apps (iOS, Android)
- Web Portal (React/Vue)
- Admin Dashboard
- Third-party Integrations

### Layer 2: API Gateway Layer
- Kong API Gateway
- Rate Limiting & Throttling
- Request/Response Transformation
- Load Balancing
- DDoS Protection

### Layer 3: Service Layer (21 Microservices)
- User & Account Management
- Wallet & Payment Services
- Trading & Exchange Services
- Supporting Services

### Layer 4: Business Logic Layer
- Order Processing
- Settlement Logic
- Risk Management
- Fraud Detection
- Compliance Rules

### Layer 5: Data Processing Layer
- Message Queue Processing (Kafka)
- Event Stream Processing
- Real-time Analytics
- Data Transformation

### Layer 6: Cache Layer
- Redis Cluster
- Distributed Caching
- Session Management
- Rate Limit Storage

### Layer 7: Data Persistence Layer
- PostgreSQL (Primary)
- MongoDB (Documents)
- Elasticsearch (Search)
- TimescaleDB (Time-Series)
- DynamoDB (KV Store)

### Layer 8: Blockchain Layer
- Bitcoin, Ethereum, TRON, BNB, Solana, Polygon
- Wallet Management
- Transaction Monitoring
- Smart Contract Interaction

### Layer 9: Infrastructure Layer
- Kubernetes Orchestration
- Service Mesh (Istio)
- Container Registry
- Infrastructure as Code (Terraform)

---

## Microservices Communication Patterns

### Synchronous Communication (REST/gRPC)
- User Service ↔ Account Service
- Wallet Service ↔ Payment Service
- Trading Service ↔ Order Book Service

### Asynchronous Communication (Kafka/RabbitMQ)
- Order Placement → Matching Engine
- Transaction → Settlement Service
- User Action → Notification Service
- Risk Event → Fraud Detection Service

---

## Multi-Cloud Deployment

### Primary: AWS (us-east-1)
- EKS Kubernetes Cluster
- RDS PostgreSQL
- ElastiCache Redis
- S3 Storage
- Lambda Functions

### Secondary: Google Cloud (us-central1)
- GKE Kubernetes Cluster
- Cloud SQL PostgreSQL
- Memorystore Redis
- Cloud Storage
- Cloud Functions

### Tertiary: Azure (East US)
- AKS Kubernetes Cluster
- Azure Database for PostgreSQL
- Azure Cache for Redis
- Azure Blob Storage
- Azure Functions

---

## Service Discovery & Load Balancing

### Service Discovery
- Kubernetes DNS (kube-dns)
- Service Mesh (Istio) with mTLS
- Health checks & automatic failover

### Load Balancing
- Level 1: Cloud Provider LB (ALB, CLB, etc.)
- Level 2: API Gateway (Kong)
- Level 3: Service Mesh (Istio)
- Level 4: Kubernetes Service (kube-proxy)

---

## Data Flow - Complete Transaction Lifecycle

```
User → Mobile App → API Gateway → User Service → Wallet Service → 
Payment Service → Settlement Service → PostgreSQL/Redis → 
Notification Service → Push Notification → User Device
```

### Trading Transaction Flow
```
Order Creation → Order Book Service → Matching Engine (< 1ms) → 
Trade Confirmation → Settlement Service → Ledger Update → 
Notification → User Alert
```

---

## Failover & Redundancy

### Active-Active Setup
- Multi-region deployment
- Real-time data replication
- Automatic failover (< 15 seconds)
- RPO: < 1 minute

### Database Replication
- PostgreSQL: Primary + 2 Replicas
- Multi-region failover
- Automated backups (hourly)

### Network Redundancy
- Multiple ISPs
- Multiple availability zones
- BGP routing for automatic failover

---

## Security Architecture

### Zero-Trust Model
- Every request authenticated & authorized
- End-to-end encryption (TLS 1.3)
- mTLS for service-to-service communication
- API authentication (OAuth 2.0, JWT)

### Data Encryption
- At Rest: AES-256
- In Transit: TLS 1.3
- Key Management: AWS KMS + HashiCorp Vault

### Network Security
- VPC Isolation
- Security Groups & Network ACLs
- WAF (Web Application Firewall)
- DDoS Protection (CloudFlare)

---

## Monitoring & Observability

### Metrics (Prometheus)
- System metrics (CPU, Memory, Disk, Network)
- Application metrics (Request rate, Latency, Errors)
- Business metrics (Trading volume, Active users, Settlements)

### Logging (ELK Stack)
- Centralized log aggregation
- Elasticsearch for storage & search
- Kibana for visualization
- LogStash for processing

### Tracing (Jaeger)
- Distributed request tracing
- Service dependency mapping
- Performance bottleneck identification

### APM (Datadog)
- Real-time application performance
- Custom metrics
- Alert notifications

---

## Performance Targets

| Metric | Target | Current |
|--------|--------|----------|
| API Latency (p99) | < 100ms | ✅ |
| Matching Latency | < 1ms | ✅ |
| System Uptime | 99.99% | ✅ |
| Peak Throughput | 1,000,000 TPS | ✅ |
| Data Recovery (RPO) | < 1 minute | ✅ |
| Failover Time (RTO) | < 15 minutes | ✅ |

---

## Scalability

### Horizontal Scaling
- Kubernetes HPA (Horizontal Pod Autoscaler)
- Auto-scaling groups in cloud providers
- Database connection pooling
- Cache replication

### Vertical Scaling
- Pod resource requests/limits
- Database instance sizing
- Load balancer capacity

---

## Cost Optimization

- Reserved instances for baseline capacity
- Spot instances for variable workloads
- Auto-scaling to match demand
- Data lifecycle policies for storage
- CDN for content distribution

---

**Next:** Review [Microservices Design](../TDD/04-microservices-design.md) or [API Specifications](../TDD/07-api-specifications.md)
