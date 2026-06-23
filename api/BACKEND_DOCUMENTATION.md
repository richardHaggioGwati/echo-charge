# Technical Documentation

> The sole purpose of this file is to document everything.
>
> $\color{red}{\text{This file is subject to changes!}}$
>
> Out of a desire to keep functions concise and code clean, the comments that may describe what a function is doing may not be enough to completely illustrate what a file is doing and how that is accomplished. So for any future viewers of this project and myself, this README exist in orders to Hopefully answer the questions "what's this file/function even doing...."

---

## Table of Contents

[Overview](#overview)

[File Structure](#file-structure)

[Server Entry Point](#server-entry-point-cmd)

---

## Overview

I will try to explain the thought process behind every file/function I think could go over someone head including mine in a day or two. I will however not explain every line as that would get long and boring. Complex functions or parts of complex functions will be my main priority and any tricks deemed not to be straight forward. I will also reference resources that may contain a better explanation.

---

## File Structure

```text
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
└── BACKEND_DOCUMENTATION.md
```

## Files or Functions

### Server Entry Point (cmd)

This file servers as the server entry point for the whole backend/api

```go
srv := &http.Server{
Addr:         addr,
Handler:      http.HandlerFunc(tmpHandler),
ReadTimeout:  10 * time.Second,
WriteTimeout: 10 * time.Second,
IdleTimeout:  60 * time.Second,
}
```

Creates a new server configuration contain the hardcoded address which will probable be read in via an environment variable in the future. Handler points to a function that will handle all incoming requests.ReaTimeout is the maximum amount of time allowed to read a request. Write timeout works basically on the same principal allowing limited time to write. I am sure you can take a smart guess at what Idle timeout is.

#### Starting the server in the background

The server is started with a goroutine which is a lightweight thread managed by Go. It allows functions to run concurrently. You can read more [go by example](https://gobyexample.com/goroutines). Without it the code would block forever.

```go
go func() {
    slog.Info("EchoCharge Server Starting..", "addr", addr)
    if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
      slog.Error("server error", "err", err)
      os.Exit(1)
 }  
}()
```

<code>slog.Info</code> Is a message printed to the console to indicate that the server has started running. We listen for an error when attempting to run the server with <code>if err := srv.ListenAndServe();</code>. If our error struct is not empty, we check for a specific error <code>http.ErrServerClosed</code> which indicates that our server is gracefully shutting down. If we have any other error then it will be printed on the console by <code>slog.Error("server error", "err", err)</code>. If that does happen we are to terminate instantly with with failure status (1).

#### Creating channels

```go
quit := make(chan os.Signal, 1)
````

The channel is created for the purpose of waiting for a shutdown signal if the server receives one. The channel is only capable of holding one OS signal.

```go
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
````

Register or listen for particular signals that can be outputted. The signals are SIGINT or SIGTERM into the channel. Which are produced <code>Ctrl+C</code> or <code>kill pid</code> on bash.

This line <code><-quit</code> blocks until the signal arrives.

#### Graceful Shutdown

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	5*time.Second,
)
````

Creates a timeout context with a deadline, allowing the server at the most 5 seconds to shutdown cleanly.

<code>context.Background()</code> Creates a root context with no cancellation or deadline. <code>defer cancel()</code>, releases timer resources when function exits.

#### Shutdown Server

```go
if err := srv.Shutdown(ctx); err != nil {
    slog.Error("forced shutdown", "err", err)
}
slog.Info("stopped")
````

Gracefully shuts down the server. Internally the server stops accepting new connections, allows active requests to finish and waits until either all request are finished or context expires.

Forced shutdown occurs when request take too long or timeout expires.
