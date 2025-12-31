# Architecture Documentation

## Overview

This project follows a **Domain-Driven Design (DDD)** approach with **Event-Driven Architecture (EDA)** principles. The codebase is organized into four main layers:

1.  **Domain**: Contains the core business logic, entities, value objects, and domain events. It is independent of external frameworks.
2.  **Application**: Orchestrates the domain logic to fulfill use cases. It handles commands, queries, and event publishing.
3.  **Infrastructure**: Provides concrete implementations for interfaces defined in the domain and application layers (e.g., repositories, event bus, database connections).
4.  **Worker**: (Optional) Handles background processes or event consumers.

## Directory Structure

```
internal/
├── domain/         # Core business logic (Entities, Value Objects, Events, Repository Interfaces)
├── application/    # Use Cases, Command/Query Handlers, Event Publishers
├── infrastructure/ # Concrete implementations (DB, Event Bus, External Services)
└── worker/         # Background workers
```

## Event-Driven Flow

The system uses an **Event Bus** to decouple components.

1.  **Command**: A request initiates a command (e.g., `RegisterUserCommand`).
2.  **UseCase**: The UseCase handles the command.
    *   It interacts with the **Domain** to perform business logic (e.g., `user.RegisterUser`).
    *   The Domain entity records **Domain Events** (e.g., `UserRegistered`).
    *   The UseCase persists the entity via a **Repository**.
    *   The UseCase publishes the recorded events using the **EventPublisher**.
3.  **Event Bus**: The `EventPublisher` (implemented by `InMemoryBus` in infrastructure) dispatches the event to registered **Handlers**.
4.  **Handler**: Event Handlers (e.g., `SendWelcomeEmailHandler`) react to the event (e.g., sending an email).

### Current Limitation: Synchronous Bus

> [!NOTE]
> The current `InMemoryBus` implementation is **synchronous**. Event handlers are executed in the same goroutine as the publisher. This ensures immediate consistency but means that slow handlers will block the main request. For production systems requiring high throughput or resilience, an asynchronous bus (using goroutines or a message queue like RabbitMQ/Kafka) is recommended.

## Key Components

### Domain Layer (`internal/domain`)
-   **User**: The aggregate root.
-   **Events**: `UserRegistered` (defined in `events.go`).

### Application Layer (`internal/application`)
-   **UseCase**: `UserUseCase` coordinates the registration process.
-   **Event Publisher Interface**: Defines how events are published.
-   **Event Handler Interface**: Defines how events are consumed.

### Infrastructure Layer (`internal/infrastructure`)
-   **InMemoryBus**: A simple loop-based event bus.
-   **Persistence**: Database repositories (e.g., Postgres).
