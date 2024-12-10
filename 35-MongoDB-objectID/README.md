## MongoDB Object ID Structure (12 bytes)

MongoDB's ObjectID is a 12-byte identifier broken down into the following components:

| Component        | Size (bytes) | Bitwise Breakdown                           | Description                                  |
|------------------|--------------|---------------------------------------------|----------------------------------------------|
| **Timestamp**     | 4 bytes      | 0 - 31 (0x00000000 - 0xFFFFFFFF)            | 32-bit Unix timestamp (seconds since epoch)|
| **Machine ID**    | 3 bytes      | 32 - 55 (0x000000 - 0xFFFFFF)               | Identifier for the machine generating the ObjectID |
| **Process ID**    | 2 bytes      | 56 - 71 (0x0000 - 0xFFFF)                   | Process ID across multiple processes |
| **Counter**       | 3 bytes      | 72 - 95 (0x000000 - 0xFFFFFF)               | Counter for IDs created at the same timestamp |

- Timestamp (4 bytes): 6758617b 
- Machine ID (3 bytes): 454a2a
- Process ID (2 bytes): 80b8
- Counter (3 bytes): 5fedd1

Output
```
Machine ID: EJ*
6758617b454a2a80b85fedc3
6758617b454a2a80b85fedc4
6758617b454a2a80b85fedc5
6758617b454a2a80b85fedc6
6758617b454a2a80b85fedc7
6758617b454a2a80b85fedc8
6758617b454a2a80b85fedc9
6758617b454a2a80b85fedca
6758617b454a2a80b85fedcb
6758617b454a2a80b85fedcc
6758617b454a2a80b85fedcd
6758617b454a2a80b85fedce
6758617b454a2a80b85fedcf
6758617b454a2a80b85fedd0
6758617b454a2a80b85fedd1
```
