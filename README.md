# Barbatos

A comprehensive Go library providing common interfaces and utilities for application development. Barbatos offers standardized abstractions for caching, logging, pub/sub messaging, ORM operations, and application lifecycle management.

## Installation

```shell
go get github.com/zeroxsolutions/barbatos
```

## Requirements

- Go 1.18 or higher

## Features

### 🚀 Application Lifecycle Management
- **App Interface**: Standardized application startup and shutdown management
- Graceful shutdown handling
- Resource cleanup coordination

### 💾 Caching
- **Cache Interface**: Universal caching abstraction
- Support for key-value operations with expiration
- Pattern-based key operations
- Connection status monitoring

### 📝 Logging
- **Logger Interface**: Comprehensive logging abstraction
- Multiple log levels: Debug, Info, Warn, Error, Panic, Fatal
- Support for formatted logging and structured logging
- Key-value pair logging for better log analysis

### 📡 Pub/Sub Messaging
- **Publisher Interface**: Message publishing abstraction
- **Subscriber Interface**: Message subscription and consumption
- **Message Interface**: Standardized message format
- Topic-based messaging system
- Connection status monitoring

### 🗄️ ORM (Object-Relational Mapping)
- **MModel**: Base model for MySQL databases
- **PModel**: Base model for PostgreSQL databases
- UUID primary keys with automatic generation
- Automatic timestamps (created_at, updated_at)
- Soft deletion support
- GORM integration

## Package Overview

### `app`
Provides the `App` interface for managing application lifecycle:

```go
type App interface {
    Run() error
    Shutdown() error
}
```

### `cache`
Offers a universal caching interface with support for:
- Connection status checking
- Key pattern matching
- TTL (Time To Live) support
- Bulk operations

```go
type Cache interface {
    IsConnected(ctx context.Context) bool
    Keys(ctx context.Context, pattern string) ([]string, error)
    Get(ctx context.Context, key string) (string, error)
    Set(ctx context.Context, key string, value interface{}) error
    SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    Del(ctx context.Context, keys ...string) error
    DelWithPattern(ctx context.Context, pattern string) error
    Close() error
}
```

### `log`
Provides comprehensive logging capabilities:
- Multiple log levels with different output methods
- Formatted logging (using format strings)
- Structured logging (using key-value pairs)
- Panic and Fatal level support

```go
type Logger interface {
    Debug(args ...interface{})
    Debugf(template string, args ...interface{})
    Debugw(msg string, keysValues ...interface{})
    // ... similar methods for Info, Warn, Error, Panic, Fatal
}
```

### `pubsub`
Implements publish-subscribe messaging patterns:

**Publisher:**
```go
type Publisher interface {
    Publish(ctx context.Context, topic string, messages ...[]byte) error
    IsConnected(ctx context.Context) bool
    Close() error
}
```

**Subscriber:**
```go
type Subscriber interface {
    Subscribe(ctx context.Context, topics ...string) error
    Unsubscribe(ctx context.Context, topics ...string) error
    Receiver(ctx context.Context) (<-chan Message, error)
    IsConnected(ctx context.Context) bool
    Close() error
}
```

**Message:**
```go
type Message interface {
    Topic() string
    Data() []byte
}
```

### `orm`
Provides base models for database operations:

**MySQL Model (MModel):**
- String UUID primary key
- MySQL-optimized field types
- Automatic UUID generation via BeforeCreate hook

**PostgreSQL Model (PModel):**
- UUID type primary key
- PostgreSQL-optimized field types
- Native UUID generation

Both models include:
- `ID`: Primary key (string UUID for MySQL, uuid.UUID for PostgreSQL)
- `CreatedAt`: Creation timestamp
- `UpdatedAt`: Last modification timestamp
- `DeletedAt`: Soft deletion timestamp (GORM soft delete)

## Usage Examples

### Application Lifecycle
```go
package main

import (
    "github.com/zeroxsolutions/barbatos/app"
)

type MyApp struct {
    // your app fields
}

func (a *MyApp) Run() error {
    // Initialize and start your application
    return nil
}

func (a *MyApp) Shutdown() error {
    // Gracefully shutdown your application
    return nil
}

func main() {
    var myApp app.App = &MyApp{}
    
    if err := myApp.Run(); err != nil {
        panic(err)
    }
    
    // Handle shutdown...
    myApp.Shutdown()
}
```

### Using ORM Models
```go
package main

import (
    "github.com/zeroxsolutions/barbatos/orm"
    "gorm.io/gorm"
)

// User model embedding MModel for MySQL
type User struct {
    orm.MModel
    Name     string `json:"name" gorm:"column:name;type:varchar(255);not null"`
    Email    string `json:"email" gorm:"column:email;type:varchar(255);unique;not null"`
    IsActive bool   `json:"isActive" gorm:"column:is_active;type:tinyint(1);default:1"`
}

// PostgreSQL version using PModel
type UserPG struct {
    orm.PModel
    Name     string `json:"name" gorm:"column:name;type:varchar(255);not null"`
    Email    string `json:"email" gorm:"column:email;type:varchar(255);unique;not null"`
    IsActive bool   `json:"isActive" gorm:"column:is_active;type:boolean;default:true"`
}
```

### Implementing Cache Interface
```go
package main

import (
    "context"
    "time"
    "github.com/zeroxsolutions/barbatos/cache"
)

func useCacheExample(c cache.Cache) {
    ctx := context.Background()
    
    // Check connection
    if !c.IsConnected(ctx) {
        panic("Cache not connected")
    }
    
    // Set a value with expiration
    err := c.SetWithExpiration(ctx, "key1", "value1", 5*time.Minute)
    if err != nil {
        panic(err)
    }
    
    // Get the value
    value, err := c.Get(ctx, "key1")
    if err != nil {
        panic(err)
    }
    
    // Clean up
    defer c.Close()
}
```

## Architecture

Barbatos follows interface-driven design principles, providing abstract interfaces that can be implemented by various concrete implementations. This approach offers:

- **Flexibility**: Easy to swap implementations
- **Testability**: Simple to mock interfaces for testing
- **Consistency**: Standardized patterns across different components
- **Maintainability**: Clear separation of concerns

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
