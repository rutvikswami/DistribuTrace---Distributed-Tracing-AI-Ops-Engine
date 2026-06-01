# DistribuTrace — Distributed Tracing & AI Ops Engine

DistribuTrace is an end-to-end, cloud-native distributed tracing and AI-powered observability platform designed to help engineers monitor, analyze, and troubleshoot microservices.

By combining a custom lightweight tracing SDK, a high-performance span collector, automated anomaly detection, and an intelligent Root Cause Analysis (RCA) engine powered by Large Language Models (LLMs), DistribuTrace aims to transform complex telemetry data into actionable operational insights.

---

## 🚀 Key Features

### Custom Go Tracer SDK
- Lightweight distributed tracing SDK written in Go
- Context propagation across services
- Parent-child span correlation
- HTTP/gRPC trace context support

### High-Performance Collector
- Go-based telemetry collector
- Span validation and ingestion pipeline
- Batch processing and storage
- Optimized for high throughput

### AI-Powered Root Cause Analysis
- LangChain-powered RCA agent
- Gemini AI integration
- Historical incident retrieval using ChromaDB
- Automated remediation suggestions

### Real-Time Anomaly Detection
- Service latency monitoring
- Statistical anomaly detection
- Alert generation and severity classification

### Trace Visualization Dashboard
- Interactive waterfall traces
- Service dependency visualization
- Latency heatmaps
- Alert and RCA dashboards

---

# 🏗️ System Architecture

```text
       ┌────────────────────────────────────────────────────────┐
       │                Instrumented Microservices              │
       └──────────────────────────┬─────────────────────────────┘
                                  │
                                  ▼
       ┌────────────────────────────────────────────────────────┐
       │                 Custom Go Tracer SDK                   │
       └──────────────────────────┬─────────────────────────────┘
                                  │
                                  ▼
       ┌────────────────────────────────────────────────────────┐
       │               Collector Service (Go)                   │
       └──────────────────────────┬─────────────────────────────┘
                                  │
                                  ▼
       ┌────────────────────────────────────────────────────────┐
       │               PostgreSQL Trace Store                   │
       └──────┬───────────────────┬──────────────────────┬──────┘
              │                   │                      │
              ▼                   ▼                      ▼
       ┌──────────────┐    ┌──────────────┐       ┌──────────────┐
       │  Anomaly     │    │ FastAPI      │       │ LangChain    │
       │  Detector    │    │ Backend      │       │ RCA Agent    │
       │  (Python)    │    │ (Python)     │       │ (Python)     │
       └──────┬───────┘    └──────┬───────┘       └──────┬───────┘
              │                   │                      │
              ▼                   ▼                      ▼
       ┌──────────────┐    ┌──────────────┐       ┌──────────────┐
       │   Alerting   │    │  React Web   │       │ Gemini AI &  │
       │   System     │    │  Dashboard   │       │ ChromaDB RAG │
       └──────────────┘    └──────────────┘       └──────────────┘
```

---

# 🛠️ Technology Stack

| Layer | Technologies | Purpose |
|---------|------------|---------|
| Tracing SDK | Go | Span generation and context propagation |
| Collector | Go, Docker | High-performance trace ingestion |
| Database | PostgreSQL | Trace and alert storage |
| Vector Store | ChromaDB | Historical incident retrieval |
| Backend API | FastAPI, Python | REST APIs and analytics |
| AI Engine | LangChain, Gemini AI | Root Cause Analysis |
| Frontend | React, Tailwind CSS | Observability dashboard |
| Infrastructure | Docker, Docker Compose | Local orchestration and deployment |

---

# 📂 Monorepo Structure

```text
DistribuTrace/
│
├── tracer-sdk/
├── collector/
├── anomaly/
├── api/
├── frontend/
│
├── docker-compose.yml
├── .env
├── README.md
└── .gitignore
```

---

# 🚧 Project Status

DistribuTrace is currently under active development.

## Completed ✅

- [x] Monorepo setup
- [x] Go collector service
- [x] Dockerized collector service
- [x] GitHub repository setup

## In Progress 🟡

- [ ] Docker Compose infrastructure
- [ ] PostgreSQL integration
- [ ] Environment configuration management

## Planned 🔵

- [ ] Custom Go Tracer SDK
- [ ] Span collector pipeline
- [ ] Trace storage layer
- [ ] Anomaly detection engine
- [ ] LangChain RCA agent
- [ ] ChromaDB integration
- [ ] FastAPI REST layer
- [ ] React dashboard
- [ ] AWS deployment

---

# 🎯 Learning Goals

This project is being built from scratch to understand and implement:

- Distributed tracing internals
- Context propagation
- Span correlation
- Observability architecture
- AI-powered incident analysis
- Vector databases and RAG pipelines
- Cloud-native deployment workflows

---

# 📦 Local Development

## Clone Repository

```bash
git clone https://github.com/<your-username>/DistribuTrace.git
cd DistribuTrace
```

## Build Collector

```bash
cd collector

docker build -t distributrace-collector .
```

## Run Collector

```bash
docker run -p 8080:8080 distributrace-collector
```

## Verify

Open:

```text
http://localhost:8080
```

Expected output:

```text
DistribuTrace Collector Running
```

---

# 🏷️ Topics

```
distributed-tracing
observability
golang
python
fastapi
react
postgresql
docker
microservices
langchain
chromadb
gemini-ai
aiops
monitoring
devops
cloud-native
```

---

# 🤝 Contributing

Contributions, ideas, and feedback are welcome.

Please open an issue before starting major work so discussions can happen before implementation.

---

# 📄 License

This project is licensed under the MIT License.