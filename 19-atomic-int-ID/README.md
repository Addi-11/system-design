## Atomically incrementing integer ID
Implement an atomically incrementing counter-based unique ID generator.
- Atomically increment counter and persist it on disk every time
- Persist the counter every n iterations to the disk
- Vary the n to understand how the throughput varies
- Implement and test recovery by resuming from the safest value


Thought process for generating unique IDs in a distributed system:

1. **Using the Clock**: 
   - A timestamp-based ID could work but may lead to conflicts in distributed systems where multiple machines generate IDs simultaneously.

2. **Adding `machine_id`**:
   - To make IDs unique across machines, combine `machine_id` with the timestamp: `machine_id + time`.
  
3. **Considering Multithreading**:
   - With multiple threads on each machine, add `thread_id` for further uniqueness: `machine_id + time + thread_id`.
  
4. **Replacing `thread_id` with a Counter**:
   - Using a counter instead of `thread_id` makes IDs sequentially unique within each machine and thread: `machine_id + time + counter`.

5. **Eliminating the Timestamp**:
   - If each ID has a unique counter per machine, the timestamp becomes redundant, so we can simplify to `machine_id + counter`.


#### Summary
1. **Atomic Counter Increment**: Each ID is created by combining `machine_id` with a globally shared, atomically incremented `counter`.
2. **Periodic Persistence**: The counter's value is saved to disk every `n` iterations to safeguard progress without saving every time, which balances performance and reliability.
3. **Fault Recovery**: On startup, the counter loads from the last saved value, ensuring continuity even after program restarts.

|n   |Time Taken |
|----|-----------|
|1000  | 408.5958ms|
|5000  | 2.1122575s|
|10000 |3.8377447s|
|100000|1m4.908891s|