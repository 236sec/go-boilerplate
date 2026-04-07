---
description: "Use when writing, reviewing, or refactoring Go code to enforce SOLID principles, avoid code smells, and maintain idiomatic Go conventions."
applyTo: "**/*.go"
---

# Go Best Practices & SOLID Principles

## 1. SOLID Principles in Go

- **Single Responsibility Principle (SRP)**: Keep structs and functions focused on a single purpose. Prefer small, cohesive files.
- **Open/Closed Principle (OCP)**: Allow behavior extension through interfaces and struct embedding, not by modifying existing code.
- **Liskov Substitution Principle (LSP)**: Ensure implementations of an interface don't panic or break contracts unexpectedly. Return explicit domain errors.
- **Interface Segregation Principle (ISP)**: Give clients exactly what they need. Prefer small interfaces (e.g., 1-3 methods like standard library `io.Reader`). Define interfaces where they are _used_, rather than where they are implemented.
- **Dependency Inversion Principle (DIP)**: High-level modules (use-cases) must not depend on low-level modules (repos). Both should depend on abstractions (interfaces). Follow the project's dependency injection pattern using [src/di/](src/di) and `sync.OnceValue` instead of `init()` hooks.

## 2. Avoiding Code Smells

- **Return Early (Guard Clauses)**: Avoid deep nesting. Handle errors or invalid conditions first, and keep the "happy path" aligned to the left margin.
- **No Package-Level State**: Avoid global variables (`var` at the package level) and `init()` functions to maintain testability and avoid side effects. Pass dependencies explicitly.
- **Magic Values**: Avoid magic numbers or strings. Define them as constants or custom types.
- **Large Functions**: Break functions longer than ~50 lines into smaller, testable helpers.
- **Improper Struct Initialization**: Prefer explicitly named initialization parameters (e.g., `User{Name: "X"}`) or a dedicated `New()` constructor function when internal maps/slices need initialization.

## 3. Error Handling and Idiomatic Go

- **Contextual Errors**: Never use `_` to suppress errors. Wrap errors with context using `fmt.Errorf("action failed: %w", err)` so the caller can trace the root cause.
- **No Panics**: Never use `panic` for standard control flow or validation. Only panic for developer errors that signify a truly unrecoverable state (e.g., malformed statically compiled regex).
- **Receiver Types**: Be consistent. Use pointer receivers `func (s *Struct)` if you need to mutate the state or if the struct is large. Don't mix value and pointer receivers for the same struct.
- **Concurrency**: Prevent Goroutine leaks. Always ensure that every goroutine you launch has a guaranteed mechanism to exit (using context cancellation or closed channels).
