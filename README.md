DistribuTrace — Distributed Tracing & AI Ops EngineDistribuTrace is an end-to-end, cloud-native distributed tracing and AI-powered observability platform designed to help engineers monitor, analyze, and troubleshoot microservices. By combining a custom lightweight tracing SDK, a high-performance span collector, automated anomaly detection, and an intelligent root-cause analysis engine powered by Large Language Models (LLMs), DistribuTrace turns complex system telemetry into actionable incident investigation files.🚀 Key Architectural Highlights (Planned & Active)Custom Go Tracer SDK: Minimalist tracing SDK featuring localized context propagation, parent-child span correlation, and structured context-carrying headers (HTTP/gRPC).High-Performance Collector: A highly optimized Go-based ingester designed to parse, batch, and flush span payload streams into persistent storage.AI-Powered Root Cause Analysis (RCA): An intelligent diagnostic agent using LangChain, Gemini AI, and ChromaDB vector retrieval to cross-reference real-time trace anomalies with historical system incidents and suggest direct remediation code.Trace Visualization Dashboard: A React-based Single Page Application (SPA) visualizing interactive span trace trees, service dependency map DAGs, and latencies.🗺️ System ArchitectureThe blueprint below outlines the flow of span telemetry from target microservices through the telemetry pipeline to the front-end dashboard and AI diagnostic agents.       ┌────────────────────────────────────────────────────────┐
       │                Instrumented Microservices              │
       └──────────────────────────┬─────────────────────────────┘
                                  │ (Telemetry Context Propagation)
                                  ▼
       ┌────────────────────────────────────────────────────────┐
       │                 Custom Go Tracer SDK                   │
       └──────────────────────────┬─────────────────────────────┘
                                  │ (Span Batches / JSON or Protobuf)
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
🛠️ Technology StackLayerTechnologiesRole / PurposeTelemetry & SDKGoHigh-concurrency client tracerCollector PipelineGo, DockerUltra-fast ingestion & validation of trace eventsStorage EnginePostgreSQL, ChromaDBStructured traces & high-dimensional context embeddingsAnalytical EnginePython, FastAPIBackend APIs & statistical anomaly calculationAI Ops AgentLangChain, Gemini, RAGAutomated incident triage, historical lookup, and RCAUser InterfaceReact, TailwindCSSWaterfall trace visualizer & service health maps🚧 Project Status & RoadmapDistribuTrace is currently in its early foundational stage. We are practicing transparent, milestone-driven development. Here is the current progress of the monorepo:Completed 🟢[x] Monorepo Structure Setup: Standardized layouts for Go services, Python runtimes, and shared infrastructure configs.[x] Go Collector Engine: Basic HTTP span receiver implementing thread-safe concurrent JSON telemetry parsing.[x] Docker Containerization: Dockerfile configuration for localized execution of the collector service.In Progress 🟡[ ] PostgreSQL Schema Design: Structuring normalization schemas for spans, traces, and service maps to handle high-write loads.[ ] Docker Compose Infrastructure Orchestration: Setting up unified networking, automated volumes, and cluster orchestration for the local database and collector instances.Planned 🔵[ ] Tracer SDK Development: Implement custom trace generation APIs for manual span-instrumentation in Go codebases.[ ] Collector Performance Pipeline: Add batch queueing, workers, rate-limiting, and error-handling pipelines to the Go Collector.[ ] Statistical Anomaly Detection: Python module evaluating trace latencies dynamically to alert on sudden anomalies.[ ] AI Root-Cause Engine: Integrate LangChain, ChromaDB, and Gemini AI to fetch relevant historical trace issues and auto-generate incident mitigation playbooks.[ ] React Dashboard UI: Build interactive waterfall charts representing hierarchical span relationships.📦 Local Development Setup (Current Collector Service)You can launch and test the collector service in its current state using Docker.PrerequisitesDocker DesktopGo 1.21+ (Optional, for local Go builds)1. Run the Collector ContainerBuild and spin up the Go Collector directly:# Build the Docker image
docker build -t distributrace-collector -f services/collector/Dockerfile .

# Start the collector on port 8080
docker run -p 8080:8080 distributrace-collector
2. Dispatch a Test Telemetry PayloadConfirm that the ingestion engine is working by dispatching a mock span JSON payload:curl -X POST http://localhost:8080/v1/traces \
  -H "Content-Type: application/json" \
  -d '{
    "trace_id": "8a3f89e1b2c3d4e5",
    "span_id": "c4d5e6f7",
    "parent_span_id": "",
    "service_name": "gateway-service",
    "operation_name": "GET /api/v1/checkout",
    "start_time_unix_nano": 1718010000000000000,
    "end_time_unix_nano": 1718010000250000000,
    "status": "OK"
  }'
🏷️ Repository TagsTo maximize discovery, make sure your GitHub repository settings include these topics:distributed-tracing • observability • golang • python • fastapi • react • postgresql • docker • microservices • langchain • chromadb • gemini-ai • aiops • monitoring • devops • cloud-native🤝 ContributingWe welcome early-stage collaboration! Since this project is in active bootstrap mode, please review our Issues Tracker or reach out directly to check what tasks are available before writing any code.📄 LicenseThis project is licensed under the MIT License.