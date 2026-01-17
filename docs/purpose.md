# The Purpose of this Repository 
![why](/docs/img/why.jpg)

The majority of the repositories in Spudmash Media serve as Quickstart templates for your next software build.

This particular repository intends to solves the Client Side issue of type-safety of objects produced by Business Logic code without piling on layers of frameworks or remote server abstractions to achieve this.

This particular repository intends to demonstrate:
- a production ready REST API written in GO
- project layout following Go's recommended project layout and some items from [project-layout](https://github.com/golang-standards/project-layout) repo.
- concurrent workload pattern using sync.Waitgroup + channels to manage batching (see [users.go#L78](../internal/users/service.go#L78))
