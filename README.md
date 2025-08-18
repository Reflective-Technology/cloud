# Cloud

This is a Go package for managing cloud modes, supporting AWS, GCP, Azure, On Premise, and an unspecified mode.

## Project Overview

This package allows you to switch between cloud modes (`aws`, `gcp`, `azure`, `on-premise`, `unspecified`) via environment variables or code, storing and retrieving the current cloud state in a thread-safe manner. It’s intended for Go applications that need to adjust behaviors according to their running environment.

## Installation & Usage

Install using the go get command:

```bash
go get github.com/Reflective-Technology/cloud
```

Import in your code:

```go
import "github.com/Reflective-Technology/cloud"
```

## Features

- Automatically sets the current cloud mode based on the `CLOUD_MODE` environment variable
- Supports AWS, GCP, Azure, On Premise, and unspecified modes
- Programmatically set mode: `cloud.SetMode("aws")`, `cloud.SetMode("gcp")`, `cloud.SetMode("azure")`, `cloud.SetMode("on-premise")`
- Get current mode: `cloud.Mode()`
- Panics on unknown mode input to ensure correctness
- Thread-safe implementation for concurrent environments

## Example

```go
package main

import (
    "fmt"
    "github.com/Reflective-Technology/cloud"
)

func main() {
    cloud.SetMode(cloud.AWS)
    fmt.Println("Current cloud mode:", cloud.Mode()) // Output: aws

    cloud.SetMode(cloud.GCP)
    fmt.Println("Current cloud mode:", cloud.Mode()) // Output: gcp

    cloud.SetMode(cloud.AZURE)
    fmt.Println("Current cloud mode:", cloud.Mode()) // Output: azure

    cloud.SetMode(cloud.ON_PREMISE)
    fmt.Println("Current cloud mode:", cloud.Mode()) // Output: on-premise

    cloud.SetMode("")
    fmt.Println("Current cloud mode:", cloud.Mode()) // Output: unspecified
}
```

## Contributing

1. Fork this project and create your own branch
2. Submit a Pull Request after making changes
3. Please describe your changes in detail in the PR
4. This project is licensed under MIT License – contributions are welcome!

---

For further documentation or support, please visit the [Reflective-Technology organization page](https://github.com/Reflective-Technology).