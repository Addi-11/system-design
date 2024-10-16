## Counting Primes: Multi-Threading Program
#### APPROACH 1: Without any Threading
```
checking till 100000000 found 5761455 prime numbers. Took 3m58.674784s
```

#### APPROACH 2: Unfair Threading

Partition the entire input into consecutive batches, where each thread handles one batch.
```
Thread 0 [3, 10000003) completed in 18.7782664s
Thread 1 [10000003, 20000003) completed in 30.4307852s
Thread 2 [20000003, 30000003) completed in 39.6527454s
Thread 3 [30000003, 40000003) completed in 42.1852373s
Thread 4 [40000003, 50000003) completed in 45.5728377s
Thread 5 [50000003, 60000003) completed in 50.2821488s
Thread 6 [60000003, 70000003) completed in 52.0371237s
Thread 7 [70000003, 80000003) completed in 52.4440178s
Thread 8 [80000003, 90000003) completed in 55.500766s
Thread 9 [90000003, 100000003) completed in 56.8292743s
checking till 100000000 found 5761455 prime numbers. Took 56.8292743s
```
#### APPROACH 3: Fair Threading

We need each thread to do equal amount of work.
Rather than chunking the entire data, every thread could pick a number and check if its prime or not.

```
Thread 3 completed in 55.2213104s
Thread 2 completed in 55.1357679s
Thread 5 completed in 55.2328191s
Thread 8 completed in 55.2318185s
Thread 6 completed in 55.2318185s
Thread 1 completed in 55.2328191s
Thread 0 completed in 55.2308207s
Thread 7 completed in 55.2318185s
Thread 9 completed in 55.2328191s
Thread 4 completed in 55.1561414s
checking till 100000000 found 5761455 prime numbers. Took 55.2346881s
```