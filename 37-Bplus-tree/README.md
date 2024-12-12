## Implementing B+ Tree based KV store
A **B+ Tree** is a self-balancing tree data structure commonly used in databases and file systems for efficient data storage and retrieval.

1. **Structure**: 
   - It consists of **internal nodes** (for indexing) and **leaf nodes** (which store actual data).
   - The tree is organized to maintain sorted keys and allows sequential traversal through the leaf nodes, making it suitable for range queries.

2. **Properties**:
   - All leaf nodes are at the same level, ensuring balanced access time.
   - Internal nodes only store keys for navigation, while leaf nodes store key-value pairs.
   - The number of children per node (order) is fixed, e.g., an order-4 tree allows up to 3 keys per node.

3. **Operations**:
   - **Insertion**: 
     - Keys are added in a sorted manner. If a node overflows (exceeds its order), it's split, and the middle key is propagated upward.
   - **Search**: 
     - Traverses down the tree, following keys in internal nodes, and finally searches in the appropriate leaf node.
   - **Splitting**: When a node becomes too full, it splits into two nodes, ensuring the tree remains balanced.

4. **Efficiency**:
   - Operations like search, insert, and delete are \(O(\log N)\), 
   - Sequential access of leaf nodes is efficient due to linked-list-like pointers.