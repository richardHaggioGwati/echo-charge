# EchoCharge

> **Simple. Reliable. Always visible.**
>
> A widget-first Android utility for monitoring the battery status of wireless earphones and Bluetooth audio devices.

---

## Overview

**EchoCharge** is a modern Android application built with **Jetpack Compose** and powered by a **Go core service layer**. It provides clear and timely battery information for Bluetooth earphones, headphones, and other audio devices through widgets, notifications, and a clean Material 3 interface.

Unlike manufacturer-specific companion apps, EchoCharge aims to provide a lightweight, privacy-friendly experience without requiring accounts or unnecessary permissions.

---

## Features

### 🔋 Battery Monitoring

* View battery percentages for supported Bluetooth devices
* Real-time updates when available
* Support for Bluetooth Low Energy (BLE) and system battery reporting
* Graceful fallback when battery information is unavailable

### 🎧 Device Management

* Discover paired Bluetooth audio devices
* View connection status
* Assign custom device aliases
* Attempt device renaming when supported

### 🧩 Widget-First Experience

* Home screen widgets for quick battery information
* Multiple widget sizes
* Designed for glanceable information

### 🔔 Persistent Notifications

* Ongoing battery notification
* Live updates while connected
* Quick access to device details

### ⚙️ Modern Android Design

* Jetpack Compose
* Material Design 3
* Dark and light themes
* Android 13+ support

## Development Goals

### Phase 1

* [ ] Device discovery
* [ ] Battery monitoring
* [ ] Material 3 UI
* [ ] Foreground service
* [ ] Persistent notification

### Phase 2

* [ ] Home screen widgets
* [ ] Custom device aliases
* [ ] Battery history
* [ ] Settings screen

### Phase 3

* [ ] Multi-device support
* [ ] Device capability detection
* [ ] Rename support
* [ ] Debug information

### Future Ideas

* [ ] Wear OS companion app
* [ ] Desktop application
* [ ] Battery analytics
* [ ] Vendor-specific integrations
* [ ] Export battery history
* [ ] OpenAPI / gRPC interface

---

## Philosophy

EchoCharge follows three principles:

### Transparency

Battery information depends on what devices expose. EchoCharge never fabricates or estimates values beyond what the system or device reports.

### Simplicity

The app is designed around quick access to information through widgets and notifications rather than complex dashboards.

### Privacy

* No account required
* No cloud synchronization
* No advertising
* Local-first architecture

---

## License

This project is released under the MIT License.

**Know your battery. At a glance.** 🎧🔋
