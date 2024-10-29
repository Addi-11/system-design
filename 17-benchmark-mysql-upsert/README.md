## Benchmark MySQL's UPSERT
- ON DUPLICATE KEY UPDATE
- REPLACE INTO

```
CREATE TABLE IF NOT EXISTS test_table (id INT PRIMARY KEY, data VARCHAR(20));
```

| `n`   | `ON DUPLICATE KEY UPDATE` Time Taken | `REPLACE INTO` Time Taken |
|-------|--------------------------------------|---------------------------|
| 1     | 2.525ms                              | 1.267ms                   |
| 10    | 2.8037ms                             | 17.238ms                  |
| 100   | 27.2609ms                            | 139.3891ms                |
| 500   | 105.1636ms                           | 706.9529ms                |
| 1000  | 217.6239ms                           | 1.2950787s                |
| 5000  | 1.235172s                            | 6.8721067s                |
| 10000 | 2.3808408s                           | 12.8781352s               |
| 50000 | 11.9876923s                          | 3m2.1582408s              |


### Reason for Performance Difference
- **`ON DUPLICATE KEY UPDATE`** only updates existing rows on conflict, which minimizes locking and reduces index operations.
- **`REPLACE INTO`** first deletes any existing row and then inserts a new one, causing additional I/O overhead, index adjustments, and potentially larger locks, which slow down performance, particularly as the data volume grows.

### Conclusion
- `ON DUPLICATE KEY UPDATE` is generally more efficient than `REPLACE INTO` for large-scale operations, especially when avoiding unnecessary deletions and insertions.
