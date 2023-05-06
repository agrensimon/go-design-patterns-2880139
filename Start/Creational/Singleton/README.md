# Creational Pattern: Singleton

## Purpose

- Restrict the instantiation of a single instance and provides global access.
- Allows for lazy initialization (isn't instantiated until needed).

## Scenarios

Situations where you want to ensure there is only one instance of a class: logging, configuration, telemetry/analytics, debugging.

You don't want the client of the class to be responsible for knowing if the instance has been created. The application code should be able to use the features when needed.

## Use

Instead of calling a constructor, call a function to get the instance.

In Go, we use structs to implement the singleton pattern.

Something to consider is concurrency and goroutine safety.