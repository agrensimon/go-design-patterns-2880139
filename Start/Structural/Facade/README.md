# Structural Patter: Facade

## Purpose:

- Provide a simple, front-facing interface to a more complex system, library, or API

## Scenarios

- Improve usability of a more complex API.
- Serve as a starting point for refactoring.
- Reduce tight coupling between different parts of a system.

## Use

A subsystem made up of 3 classes. Two different clients want to use this subsystem, but we might want to switch out this subsystem at some point. We implement a facade interface in between. 
