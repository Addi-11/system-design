## Implement 5 approaches to count post per hashtag

### Problem Statement
Extract and manage HashTags from all the uploaded photos. Assuming 5 million photos uploaded every hour is 5 million. <br>

More Details: [Designing Hashtag Service](https://github.com/arpitbbhayani/system-design-questions/blob/master/hashtag-service.md)

### Excercise:
1. Naive (count++) for every event
1. Naive batching (batch on server and then write to database)
1. Efficient batching with minimizing stop-the-world usng deep-copy
1. Efficient batching with minimizing stop-the-world using two-maps
1. Kafka adapter pattern to re-ingest the post hashtags partitioned by hashtag

### Implementation
1. Create a DB hashtags with following schema:
```
+------------+-------------+------+-----+
| Field      | Type        | Null | Key |
+------------+-------------+------+-----+
| hashtag_id | varchar(50) | NO   | PRI |
| count      | int         | YES  |     |
+------------+-------------+------+-----+
```
2. Run `insertDBData()` function from [insert_db_data](insert_db_data.go) to create a `setup_db.sql` file.
3. Run the following command to insert data into the table.
```
mysql -u root -p < setup_db.sql
```
4. Run `main.go`

### Results
For 6000 unique hastags and 10000 photo uploads, we have:
```
Naive Counting, Time Taken: 10.6695487s
Naive Batching, Time Taken: 9.0829599s
Efficient Batching (Deep-Copy), Time Taken: 547.9µs
Efficient Batching (Two-Maps), Time Taken: 521.4µs
Kafka Adapter, Time Taken: 4.4738473s
```

### Notes

#### 1. **Naive Counting**  
   - Count for each hashtag is updated directly in the database every time a new post is processed. Each post triggers an immediate database write operation.  
   - Leads to a high frequency of database writes, causing significant latency as each write is blocking. The system stops processing new posts until the write operation completes, introducing delays.

#### 2. **Naive Batching**  
   - Posts are batched in memory, and once a batch is complete, the entire batch is written to the database at once. Reduces the number of database writes, but the system still waits for the batch write to complete before processing new posts.  
   - **Stop-the-World**: The system "pauses" while waiting for the batch to be written to the database. New posts can't be processed until the database update is finished.

#### 3. **Efficient Batching with Minimizing Stop-the-World (Deep-Copy)**  
   - Batches are accumulated in memory, and a separate goroutine is used to handle the database write in parallel. Before the write, a deep copy of the batch is made, allowing the main batch to continue accepting new posts.  
   - **Stop-the-World**: The effect is minimized because the batch write happens in parallel, and the system can continue processing new posts without waiting for the database update. However, there may still be some delay due to concurrent access to shared resources.

#### 4. **Efficient Batching with Minimizing Stop-the-World (Two-Maps)**  
   - Uses two maps to handle batch updates. One map is used for accumulating counts, while another map is used for processing the database write in parallel. The system switches between maps to ensure uninterrupted processing of new posts.  
   - **Stop-the-World**: Minimized by using two maps, ensuring that while one map is being written to the database, the other map continues accepting new posts. This reduces latency and improves concurrency.

#### 5. **Kafka Adapter Pattern**  
   - Instead of directly writing to the database during post processing, a Kafka producer is used to publish hashtag updates as events to a Kafka topic. A separate Kafka consumer processes these events and updates the database asynchronously.  
   - **Stop-the-World**: Eliminated because posts are immediately processed and pushed to Kafka, with database updates happening asynchronously in the background. This **decouples** the post processing from the database write operations, ensuring no delays in post handling.
