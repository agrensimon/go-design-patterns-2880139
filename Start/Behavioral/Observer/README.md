# Behavioral Pattern: Observer

Sometimes referred to as publish-subscribe.

## Purpose

- Provides the ability for a subject to notify a set of "observers" about changes to the subject.
- Used to maintain loose coupling between elements of a system that interact with each other.

## Scenarios

- Systems where state changes of a subject cannot be predicted and the list of interested "listeners" cannot be known in advance or changes during runtime.

## Use

A subject object that has a list of interested observers. Also exposes an API for (un-)registering observers and for publishing updates.
Observer objects implement an API method which is to be called when the subject changes. 
