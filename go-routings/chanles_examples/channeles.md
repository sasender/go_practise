## What a Channel really is:
A channel is a pipe through which goroutines send data to each other safely.

Channels allow goroutines to communicate and synchronize safely by passing data instead of sharing memory, reducing race conditions and simplifying concurrent code.

## 2 Types of Channels: Unbuffered and Buffered
***Unbuffered:**
    Unbuffered Channels: Unbuffered channels are the default type of channel in Go. They provide a synchronous way of communication, meaning that both the sending and receiving Goroutines must be ready to perform the operation. The sender blocks until the receiver is ready, and vice versa.

    ## example:
    Unbuffered → direct hand-to-hand

**Buffered Channels**
     Buffered channels allow a specified number of elements to be stored in the channel. They provide asynchronous communication, meaning the sender can continue execution without waiting for the receiver, as long as the buffer is not full. Similarly, the receiver can continue execution without waiting for the sender, as long as the buffer is not empty.

     Example:
     Buffered → mailbox 📬

     