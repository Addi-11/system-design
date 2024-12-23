# Sytem Design Excercises
To run any excercise, go the folder and run the following, Additional details added in the excercise folder: 
```
go mod init example.com/main
go mod tidy
go run .
```
## Week-1
1. [Implement a simple connection pool using Bounded Blocking Queue](./05-connection-pool/)
2. [Implement Database Sharding and Routing (from API server)](./07-sharding/)
3. [Setup Read-replica from a MySQL locally](./12-mysql-read-replica/)
4. [Implement fair multi-threaded program](./08-multi-thread-program/)
5. [Implement server-sent events](./09-basic-server-sent-events/)
5. [Implement server-sent events using Message Broker](./10-broker-server-sent-event/)
5. Implement server-sent events on React Components on a web-page.
6. Setup RabbitMQ and Kafka locally.Write producer and consumer for them.
    1. [Setup RabbitMQ](./13A-rabbitmq/)
    2. [Setup Kafka](./13B-kafka/)
7. Implement real-time chat using socket IO: Slack-Realtime Text Chat [Reference](https://github.com/socketio/socket.io-chat-platform)
8. [Mock EC2 creation & implement Short Polling and Long Polling](./16-long-short-polling/)
 
## Week-2
1. [Implement Airline Check-in System](./14-airline-checkin-system/)
1. [Deadlock simulation](./34-deadlock/)
1. [Implement a toy KV store on top of MySQL](./15A-kvstore-mysql/)
1. [Implement simple sharding with a hash or range based routing strategy in above KV store](./15B-kvstore-hash-mysql/)
1. Implement flag driven consistent reads.
1. [Implement Distributed Transactions using 2PC.](./21-zomato-two-phase-commit/)
1. Ingest data in Neo4j and try paginating it.
1. Ingest data in MongDB and write aggregation pipeline.
1. Implement Message Broadcast across servers using Star Topology leveraging Redis PubSub.

## Week-3
1. [Implement a load-balancer](./20-load-balancer/)
1. [Implement a simple blogging application where you shard by user id; and try to provide a unique ID to each blog. The idea is to understand the need to ID generation when database is sharded.](./24-blog/)
1. [Build a simple atomically incrementing integer ID](./19-atomic-int-ID/)
1. [Implement the "Amazon's Way" of central ID generation service](./22-ID-generation-amazon/)
1. [Implement ths sturcutre of MongoDB Object ID](./35-MongoDB-objectID/)
1. [Benchmark the impact of UUID on relational database as Primary Key](./18-benchmark-primarykey/)
1. [Benchmark MySQL's UPSERT using `ON DUPLICATE KEY UPDATE` and `REPLACE INTO`](./17-benchmark-mysql-upsert/)
1. [Implement Flickr's Odd-Even based ID generation](./31-ID-flickr-odd-even/)
1. Implement Snowflake on
    1. API, and
    1. Database as stored procedure
1. [Benchmark Pagination approaches.](./23-benchmark-pagination/)
1. [Implement Zomato Ordering Service using Distributed Transactions using 2PC](./21-zomato-two-phase-commit/)

## Week-4

1. [Implement a Toy CDN](./25A-toy-cdn/)
1. [Mimick CDN Failover - on Toy CDN](./25B-toy-cdn-wid-failover/)
1. Implement pre-signed URL based upload on S3
1. Configure CDN to serve Popular Searches JSON response
1. Implement JWT based auhthentication
1. Build GitHub like OG image and server it via CDN
    1. Key learning: generating images in backend server and putting it behind a CDN
1. Measure the impact of denormalization
    1. Define a user collection in MongoDB with blogs as its attribute
    1. Store blogs object in the user document demonting all blogs that a person wrote.
    1. Store the entire object intead of reference.
    1. Now benchmark and find out how slow the response times gets as we increase the number of elements in the blogs array
1. [Implement Lazy Loading of images on frontend](./26-lazyloading/)
1. [Implement 5 approaches to count post per hashtag & STOP THE WORLD](./32-count-posthashtag/)
1. Populate on_msg_event while using websocket.
    1. Try to identify when the connection breaks and use that opportunity to write event to Kafka
1. Configure Redis in cluster mode and figure out how data is distrubuted
1. Implement newly unread message indicator on database
    1. Compute on the fly
    1. Creates messages table with 1 million rows
    1. Add one indexes for each column part of the where clause that is queried and measure the time taken
    1. Compute with mentioned composite indexes, and measure the performance
    1. Re-arrange the columns and mesure the performance impact

## Week-5
1. [Implement Consistent Hashing](./27A-consistent-hashing/)
1. [Implement consistent hashing as a load balancer algorithm](./20A-load-balancer-consistenthash/)
1. Solve skewness problem in consistent hashing with Virtual Nodes
1. Implement a simple in-memory single-node cache like Redis as discussed in the session
1. Implement the word dictionary on local machine
    1. [using CSV format](./28-word-dictionary-csv/)
    1. using Bitcask format
1. Partial data write problem by writing 100mb file and killing the process in between
1. [Implement Checksum based WAL or bitcast](./33-checksum/)
1. Implement Checksum with DB recovery, as discussed in session.
1. Implement Bitcask
    1. Basic KV operations
    1. Merge and compaction
1. Benchmark sequential IO vs random IO

## Week-6
1. Implement B+ Trees
1. [LSM Tree Based Key-Value Store.](./36-LSM-KVStore/)
1. [Implement Bloom Filters and measure: FPR vs Size Vs Num of Hash Func](./29-bloom-filters/)
1. Implement Deletable Bloom Filters
1. Setup HLS Streaming following Akamai’s Documentation
1. [Video HLS Streaming Server in Go](./30-hls-video-stream/)
1. Implement a TCP server that accepts 1GB file
1. Transfer the file via one POST request
1. Stream the file from client to server from scratch
1. Implement GFS


## Week-7
1. Implement recent search as discussed during the session
1. Capture search logs and make them queryable
    1. From an HTTP request, extract all possible meta info
    1. Ingest them in ES
    1. Plot different graphs, segmenations, and gain insights using Kibana
1. Implement Full Text Search on your phone contacts
    1. implement fuzzy searching
    1. implement spell correction
    1. implement synonymic query expansion
    1. add support for phonetic search
1. Cache API responses on Akamai for very short duration
    1. Option 1: Set TTL on Akamai Console
    1. Option 2: Drive TTL using response headers from the origin
1. Stream some dummy logs from local machine to S3
    1. Query them using Athena
1. Implement Task Scheduler as discussed in the session
    1. Fixed Time Execution and Cron Schedule
    1. Implement Job Puller
    1. Make Jobs Puller Fault Tolerant
    1. For your machine, compute Unit Tech Economics for Job Puller
    1. Define a format that allows user to specify any task
    1. Build capability to run it - Docker Images a simple solution but overskill for simple tasks
    1. Induce failures in your scheduler and set up alerts if you breach SLA
1. Implement Team Relabance feature in Task scheduler
    1. Do it for Fault Tolerance
    1. Do it if you want to auto scale
1. Implement Brokers in all 3 flavours
    1. SQS like broker using MySQL as backend
    1. Kafka like broker using MySQL as backend
    1. SQS like broker using Bitcask as backend
1. Create an account on Razorpay and build simple payment system using their API
    1. use their “Test Mode”
    1. use Webhooks to receive Payment Notifications

## Week-8
1. Implement GeoHash
1. Implement Zoom-in and Zoom-out using Trie based approach
1. Evaluate the difference between EVAL and EVAL_RO command
    1. Fire EVAL command on Replica and observe the output
    1. Fire EVAL command on Replica using Redis SDK and observe
1. Write Lua script and mimic matching algorithm
    1. use factors like - vehicle type, rating, etc.
1. Implement cursor based pagination on MySQL
1. Benchmark ingestion throughput when using
    1. Auto-inc as Primary key
    1. UUID as primary key

### Excercises that can be extended:
1. [Zomato Delivery System](./21-zomato-two-phase-commit/)
2. [Airline Checkin System](./14-airline-checkin-system/)
3. [Load Balancer](./20-load-balancer/)
4. [TODO: React loading using Server Sent Events](./10-broker-server-sent-event/)
5. GFS