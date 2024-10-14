### Naive Way
1. Establishing the TCP connection
2. Firing the query
3. Getting the response
4. Closing the connection

Query execution may itself take less time, but time on network is more
- 3-way handshake [SYN, SYN-ACK, ACK]
- 2-way tear down [FIN -> ACK] (sender and reciever, 4-step process)

**OPTIMIZATION**: We will reuse the connections => **Connection Pooling**

#### Reference Links:
- https://go.dev/tour/concurrency/9
- https://go.dev/doc/tutorial/database-access