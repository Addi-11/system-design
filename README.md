# Sytem Design Excercises

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
1. Implement Distributed Transactions using 2PC.
1. Ingest data in Neo4j and try paginating it.
1. Ingest data in MongDB and write aggregation pipeline.
1. Implement Message Broadcast across servers using Star Topology leveraging Redis PubSub.

## Week-3
1. Implement a load-balancer
1. Implement a simple blogging application where you shard by user id; and try to provide a unique ID to each blog. The idea is to understand the need to ID generation when database is sharded.
1. [Build a simple atomically incrementing integer ID](./19-atomic-int-ID/)
1. Implement the "Amazon's Way" of central ID generation service
1. Implement ths sturcutre of MongoDB Object ID
1. [Benchmark the impact of UUID on relational database as Primary Key](./18-benchmark-primarykey/)
1. [Benchmark MySQL's UPSERT using `ON DUPLICATE KEY UPDATE` and `REPLACE INTO`](./17-benchmark-mysql-upsert/)
1. Implement Flickr's Odd-Even based ID generation
1. Implement Snowflake on
    1. API, and
    1. Database as stored procedure
1. Implement benchmark pagination approaches
    1. Limit Offset
    1. ID Limit
    1. Implement multiple Logical Shards and Physical Servers
    1. Create multiple databases - insta1, insta2, etc.
    1. Create posts, users, profile table with same schem across all
    1. Alter the schema of one of the table to understand how tedious alter tables would be
    1. create a new database server (on some other port)
    1. Dump data from one shard and load it on another
        1. use mysqldump utility for this
        1. Implement this by iterating over rows to understand how slow and complex the process would be
1. Implement Distributed Transactions
