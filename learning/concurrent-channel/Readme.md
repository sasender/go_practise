Concurrency:
*** Concurrency can be defined as the process of a system managing and executing multiple tasks or instructions seemingly at the same time even if they are not truly running simultaneously. ***

### What are Channels ?

*** Channels can be defined as the connection between multiple multiple goroutines which allows the transfer between these goroutines. ***

You can simply send values into channels from one goroutine and receive those values into another goroutine.

Channels excel in facilitating communication between goroutines.

They provide a safe and efficient way to pass data and synchronize goroutines’ execution.

Also the shared values are usually referenced (copied) into the these goroutines involved.

