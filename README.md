# Sytem Design Excercises
To run any excercise, go the folder and run: 
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
1. Hit deadlock in database by cn top of MySQL.
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
1. Implement ths sturcutre of MongoDB Object ID
1. [Benchmark the impact of UUID on relational database as Primary Key](./18-benchmark-primarykey/)
1. [Benchmark MySQL's UPSERT using `ON DUPLICATE KEY UPDATE` and `REPLACE INTO`](./17-benchmark-mysql-upsert/)
1. Implement Flickr's Odd-Even based ID generation
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
1. Implement 5 approaches to count post per hashtag
    1. Naive (count++) for every event
    1. Naive batching (batch on server and then write to database)
    1. Efficient batching with minimizing stop-the-world usng deep-copy
    1. Efficient batching with minimizing stop-the-world using two-maps
    1. Kafka adapter pattern to re-ingest the post hashtags partitioned by hashtag
        1. Measure the number of writes on the database in each of the above approaches
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
    1. [using CSV format](./28-word-dictionary/)
    1. using Bitcask format
1. Partial data write problem by writing 100mb file and killing the process in between
1. Implement Checksum based
    1. Identify if data in WAL or Bitcask is corrupt using Checksum
    1. Implement database recovery as discussed in the session
1. Implement Bitcask
    1. Basic KV operations
    1. Merge and compaction
1. Benchmark sequential IO vs random IO

## Week-6
1. Implement LSM Trees
1. Implement B+ Trees
1. LSM Tree Based Key-Value Store. [Reference](http://daslab.seas.harvard.edu/classes/cs265/project.html)
1. Implement Bloom Filters and measure
1. False Positive Rate vs Size of the filter
1. False Positive Rate vs Number of Hash Function
1. Implement Deletable Bloom Filters
1. Setup HLS Streaming following Akamai’s Documentation
1. Implement HLS Streaming Server in Golang
1. Video Streaming Server in Go. [Reference](https://medium.com/bootdotdev/create-a-golang-video-streaming-server-using-hls-a-tutorial-f8c7d4545a0f)
1. Implement a TCP server that accepts 1GB file
1. Transfer the file via one POST request
1. Stream the file from client to server from scratch
1. Implement GFS

### Excercises that can be extended:
1. [Zomato Delivery System](./21-zomato-two-phase-commit/)
2. [Airline Checkin System](./14-airline-checkin-system/)
3. [Load Balancer](./20-load-balancer/)
4. [TODO: React loading using Server Sent Events](./10-broker-server-sent-event/)