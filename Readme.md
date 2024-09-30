# multiLog_go

**multiLog_go** is the Go-based logging backend for the **multiLog** project, which features a user interface built using **Tauri** and **Rust**. The multiLog project allows for flexible and customizable logging solutions with multiple log outputs and a clean, easy-to-use UI for managing logs.

## Project Overview

The **multiLog** project integrates a backend written in Go for efficient logging and a frontend built with Tauri and Rust for a lightweight, desktop-friendly user experience. The Go library (**multiLog_go**) provides powerful logging capabilities that can be configured to log to different outputs, while the UI lets users interact with and visualize the logs in real-time.

## Features

- **Multi-log management**: Create and manage multiple logs simultaneously.
- **Configurable log levels**: Supports different log levels like info, warning, error, and critical.
- **Real-time UI**: Visualize and manage logs using the Rust-based Tauri UI.
- **Lightweight and efficient**: Designed to be highly performant with minimal resource usage.

## Installation

To install **multiLog_go**, run:

```bash
go get github.com/Maruqes/multiLog_go


```
    import multiLog "github.com/Maruqes/multiLog_go"

```
    package main

    import (
        multiLog "github.com/Maruqes/multiLog_go"
    )

    func main() {
    multiLog.Init_multiLog()
    
    // Creating multiple logs
    multiLog.CreateLog("log1")
    multiLog.CreateLog("log2")
    multiLog.CreateLog("log3")
    multiLog.CreateLog("log4")

    // Logging messages with different severity levels
    multiLog.Error("log1", "This is an error")
    multiLog.Warning("log1", "This is a warning")
    multiLog.Critical("log1", "This is a critical error")
    multiLog.Info("log1", "This is an info message")
    }
```



### Custom Log Levels

**multiLog_go** supports several log levels to help categorize messages:

```
    multiLog.Error(log, message) 
    multiLog.Warning(log, message)
    multiLog.Critical(log, message) 
    multiLog.Info(log, message)
```


This project is part of multiLog project in my profile