### Naive Way
1. Establishing the TCP connection
2. Firing the query
3. Getting the response
4. Closing the connection

Query execution may itself take less time, but time on network is more
- 3-way handshake [SYN, SYN-ACK, ACK]
- 2-way tear down [FIN -> ACK] (sender and reciever, 4-step process)

Moreoever there is a harware limit to number of new connections that could be made to the DB. After a number we will get the following issue.

Non-pool benchmarks:
```
Non-pool time for 10 threads: 39.4193ms
Non-pool time for 100 threads: 241.0037ms
Error pinging the database: Error 1040: Too many connections
```

**OPTIMIZATION**: We will reuse the connections => **Connection Pooling**


Pool Benchmarks, for Pool Size = 10
```
Pool time for 10 threads: 26.6947ms
Pool time for 100 threads: 190.0564ms
Pool time for 200 threads: 380.6691ms
Pool time for 300 threads: 512.1296ms
Pool time for 500 threads: 870.5625ms
Pool time for 1000 threads: 1.7288569s
```
Pool benchmarks, for Pool Size = 100:
```
Pool time for 10 threads: 72.4968ms
Pool time for 100 threads: 38.1217ms
Pool time for 200 threads: 122.6984ms
Pool time for 300 threads: 103.8956ms
Pool time for 500 threads: 225.8707ms
Pool time for 1000 threads: 278.3791ms
```

#### Reference Links:
- [Accessing mysql DB using GO](https://go.dev/doc/tutorial/database-access)
- [Concurrency in GO](https://go.dev/tour/concurrency/9)
- [Thread Safe Queue Implementation](../thread-safe-queue/)
- [Blocking Queue Naive Implementation](../blocking-queue/)
- [Blocking Queue Implementation using channels](../blocking-queue-channel/)