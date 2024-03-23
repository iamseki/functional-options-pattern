# Functional Options Pattern in Golang

This code demonstrates the following aspects of the Functional Options pattern:
1. We define a `DBConfigOptions` struct to hold some db connections options details
2. We define an `Option` type as a function that takes a pointer to `DBConfigOptions` and returns another function for configuration(closure like).
3. The `InitFakeDB` function provides a handy an API-friendly way to handle options:
    ```go
      // With default values
      InitFakeDB("test:test@test.com")

      // Set some options in a "functional" way
      InitFakeDB("test:test@test.com",
        WithPort(12),
        WithPoolSize(10))
    ```

# Why? :thinking:

Imagine you're writing a function that can handle optional settings. You don't want people to accidentally pass an empty struct or None to it. You also can encapsulate default logic implementations in a better way.

The functional options pattern is like a magic trick to avoid that mess! It lets you provide configurations in a clean and readable way avoiding things like:
```go
        // if I do want to connect with default values
        InitFakeDB("test:test@test.com", nil)
        // Another way
        InitFakeDB("test:test@test.com", package.Options{})
```

> This is a basic example, but it showcases the core principles of using functional options for configuring objects in Golang. You can extend this pattern to handle more complex configurations and logic.