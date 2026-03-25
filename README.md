# gouser

A Go API Template using [Chi Router](https://go-chi.io/#/README) that integrates with [randomuser.me](https://randomuser.me/documentation) API.

This repo demonstrates:
  - Go Chi Router REST API
  - generating a Strongly Typed proxy library for randomuser.me
  - Go routines for simple and long running loads using sync.Waitgroup and channels
  - Proxy routes
  - Go Unit Testing

# Quick Start - Generating A New Project From This Template
Assuming we called our new API "myservice" we use [npx degit](https://www.npmjs.com/package/degit) to create a fresh project by cloning this template:
```
  mkdir myservice
  cd myservice
  npx degit spudmashmedia/gouser#master
```

# Getting Started
- [Prequisite Before Starting](/docs/prerequisite.md)
- [Build Instructions](/docs/build.md)
- [Testing Instructions](/docs/testing.md)
  
# Documentation
  
- [The Purpose of this Repository](/docs/purpose.md)
- [References](/docs/references.md)

# License
This code is distributed under the terms and conditions of the [MIT License](/LICENSE)
