# Structural Pattern: Adapter

## Purpose

- Allows the interface of an existing subsystem of API to be used as another interface without modifying the code of the existing API.

## Scenarios

- Enables incompatible objects to work together without having to make changes to either one. Such scenarios are common when working with legacy APIs or code that is not maintained by yourself.

## Use

A client class expects to use a particular interface for a given purpose. The target class that provides the API does not provide the same API. An adapter class can be used to allow the client to use its preferred interface, and adapt it to the target interface. Neither the client or the target needs to change their code.
