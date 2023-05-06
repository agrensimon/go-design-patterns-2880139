# Go Design Patterns
This is the repository for the LinkedIn Learning course Go Design Patterns. The full course is available from [LinkedIn Learning][lil-course-url].

![Go Design Patterns][lil-thumbnail-url]

Go, a multi-paradigm programming language, features design patterns that allow developers to address common problems efficiently. In this course, senior developer advocate Joe Marini covers creation, structural, and behavioral design patterns. Joe begins with an overview of design patterns and design pattern categories. Then he gives you overviews and examples of several creational patterns, including builder pattern, factory pattern, and singleton pattern. He does the same for structural patterns, covering adapter patterns and facade patterns. Joe concludes with behavioral patterns such as an observer pattern and an iterator pattern.

### Instructor

Joe Marini

Senior Director of Product and Engineering


Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/joe-marini).

[lil-course-url]: https://www.linkedin.com/learning/go-design-patterns
[lil-thumbnail-url]: https://cdn.lynda.com/course/2880139/2880139-1627493767900-16x9.jpg

# Notes from the Course

## Go Language Features That Affect Design Patterns

- No support for traditional OOP classes like in C# or Java.
    - No static members or constructors - affects patterns like Singleton.
- No support for class-based inheritance or method overloading.
    - Affects patterns like Visitor.
- No support for exceptions, error handling is via return values.
    - Affects patterns like Builder.
- No support for abstract classes.
    - Affects patterns like Abstract Factory and Bridge.

## Design Pattern Categories

### Creational

How objects are created during the lifetime of a program.

- Abstract Factory
- [Builder](./Start/Creational/Builder/README.md)
- [Factory](./Start/Creational/Factory/README.md)
- Lazy Initialization
- Multiton
- Object Pool
- Prototype
- [Singleton](./Start/Creational/Singleton/README.md)

### Structural

Used to organize the various parts of a program into larger, more complex structures and to provide new functionality based on that organization.

- [Adapter](./Start/Structural/Adapter/README.md)
- Bridge
- Composite
- Decorator
- Facade
- Flyweight
- Proxy

### Behavioral

Identify common patterns of communication and interaction between objects within a program, and are intended to increase flexibility and efficiency in carrying out those communications and interactions.

- Chain of Responsibility
- Command
- Interpreter
- Iterator
- Mediator
- Memento
- Observer
- Strategy
- Visitor
