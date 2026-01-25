## What is a Go Module?
A container for your Go project that manages code + dependencies

## What is go.mod? ##

go.mod is the identity card of your project.

## Why do we need Go Modules? ##
With modules ✅:

Exact versions are locked

Builds are repeatable

Easy to share code

Works anywhere

## When do we use modules?
If your code imports anything outside standard library, you need modules.

## What is a Package then? ##

*** A package is:**

A folder that groups related code

### Go module → manages code & libraries

### Terraform module → manages infrastructure templates

### Terraform module ≈ Go package

*** “Terraform modules are reusable infrastructure units, whereas Go modules manage application code and dependencies. Conceptually, Terraform modules are closer to Go packages.” ***

*** All .go files inside the SAME folder must have the SAME package name**


