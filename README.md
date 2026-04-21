# Peril - Learn Pub/Sub

A multiplayer strategy game built with Go and RabbitMQ to learn Pub/Sub messaging patterns. Completed as part of Boot.dev's [Learn Pub/Sub](https://learn.boot.dev/learn-pub-sub) course.

## Architecture

- **Server** (`cmd/server`) — Publishes pause/resume commands and consumes game logs from a durable queue.
- **Client** (`cmd/client`) — Players connect to RabbitMQ, spawn units, move armies, and engage in wars. Uses transient queues for pause and move subscriptions, and a shared durable queue for war events.
- **Internal packages:**
  - `internal/pubsub` — Reusable RabbitMQ helpers: `PublishJSON`, `PublishGob`, `SubscribeJSON`, `SubscribeGob`, `DeclareAndBind`
  - `internal/routing` — Exchange names, routing keys, and message types
  - `internal/gamelogic` — Game state, commands, and war resolution logic

## Prerequisites

- Go 1.22+
- Docker

## Running

```bash
# Start RabbitMQ
./rabbit.sh start

# Run the server
go run ./cmd/server

# Run a client (in a separate terminal)
go run ./cmd/client
```

## RabbitMQ Setup

The following exchanges need to exist (create via the management UI at http://localhost:15672, login `guest`/`guest`):

- `peril_direct` (type: direct)
- `peril_topic` (type: topic)
- `peril_dlx` (type: fanout)
