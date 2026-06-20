# EchoCharge Core

> The service layer behind EchoCharge.
>
> A lightweight Go backend responsible for managing Bluetooth devices, battery information, notifications, and application state.

---

## Overview

The backend api is responsible for the business logic of the application. It will be kept essentially small in order to be production ready first before implementing every single feature.

It monitor Bluetooth audio devices and expose their state to Android applications built with Jetpack Compose.

---

## Version 1 Goals

The initial release focuses on one problem:

> **Connect earbuds → retrieve battery information → keep it visible.**

Version 1 includes:

* Device management (hopefully)
* Battery monitoring
* Notification support
* Event-driven updates
* In-memory state management

Version 1 intentionally excludes:

* Battery history
* Analytics
* Vendor-specific integrations

---

## Architecture

```text
Bluetooth Adapter
        │
        ▼
Device Service
        │
        ▼
Battery Service
        │
        ▼
State Store
        │
        ▼
Notification Service
        │
        ▼
Android UI / Widgets
```

Services communicate through events instead of direct dependencies.

---

## Project Structure

```text
.
├── cmd/
│   └── main.go/
│
├── internal/
│   ├── battery/
│   ├── bluetooth/
│   ├── device/
│   ├── events/
│   ├── notification/
│   ├── settings/
│   └── state/
│
├── configs/
├── docs/
├── scripts/
├── test/
│
├── go.mod
├── go.sum
└── README.md
```

---

## Packages

### battery

Contains battery models, services, and in-memory battery state. Will probably be stored in the device in order to ensure no user accounts needed.

Responsibilities:

* Process battery updates
* Store current battery levels
* Publish battery events

---

### bluetooth

Provides the abstraction layer for Bluetooth communication.

Responsibilities:

* Receive Bluetooth events
* Translate platform-specific information into domain models (hopefully)

---

### device

Manages Bluetooth device information and change between devices

Responsibilities:

* Track connected devices
* Maintain aliases
* Provide device state
* Change between devices

---

### events

Implements a lightweight publish/subscribe mechanism.

Events include:

* DeviceConnected
* DeviceDisconnected
* BatteryChanged

---

### notification

Handles notification logic.

Responsibilities:

* Persistent notifications
* Low battery alerts
* Notification state

---

### settings

Stores application configuration.

Examples:

* Notification enabled
* Refresh interval

---

### state

Maintains the current in-memory application state.

Responsibilities:

* Connected devices
* Battery levels
* Synchronization

---

## Principles

* **Simplicity** The core should remain small and easy to understand. As I am still learning go and I think keeping things simple is always best practice.

* **Transparency** Battery values are only reported when exposed by the device or operating system.

* **Event Driven** Services communicate through events rather than tight coupling.

* **Local First** No cloud services, user accounts, or external dependencies are required.

---

## Roadmap

### Phase 1

* [ ] Device service
* [ ] Battery service
* [ ] Event bus
* [ ] State store
* [ ] Notification service

### Phase 2

* [ ] Persistent storage
* [ ] Battery history
* [ ] Widget preferences

### Phase 3

* [ ] Multi-device support
* [ ] Capability detection
* [ ] Diagnostics

### Future

* [ ] Wear OS support
* [ ] Desktop client
* [ ] OpenAPI interface
* [ ] gRPC interface
* [ ] Vendor-specific integrations

---

## Main Goal

EchoCharge Core follows one idea:

> **Know your battery. At a glance.**

The backend exists to provide accurate and reliable device information while remaining lightweight, privacy-friendly, and easy to extend.

---

## License

Released under the MIT License.

🚀 Happy Coding. 😊
