## Bloom Filters

Bloom filter is a probabilistic data structure that's used to determine if an element is in a set.

Implement Bloom Filters and measure
1. False Positive Rate vs Size of the filter
1. False Positive Rate vs Number of Hash Function


#### False positive Rate vs Size of Bloom Filter - (Single Hash Function)
```
False Positive Rate: 1.000000, Filter Size: 16
False Positive Rate: 1.000000, Filter Size: 32
False Positive Rate: 1.000000, Filter Size: 64
False Positive Rate: 1.000000, Filter Size: 128
False Positive Rate: 1.000000, Filter Size: 256
False Positive Rate: 0.983000, Filter Size: 512
False Positive Rate: 0.857000, Filter Size: 1024
False Positive Rate: 0.629000, Filter Size: 2048
False Positive Rate: 0.403000, Filter Size: 4096
False Positive Rate: 0.234000, Filter Size: 8192
False Positive Rate: 0.131000, Filter Size: 16384
False Positive Rate: 0.071000, Filter Size: 32768
False Positive Rate: 0.032000, Filter Size: 65536
False Positive Rate: 0.015000, Filter Size: 131072
False Positive Rate: 0.006000, Filter Size: 262144
```

#### False positive Rate vs Size of Bloom Filter - (10 Hash Function)
```
False Positive Rate: 1.000000, Filter Size: 16
False Positive Rate: 1.000000, Filter Size: 32
False Positive Rate: 1.000000, Filter Size: 64
False Positive Rate: 1.000000, Filter Size: 128
False Positive Rate: 1.000000, Filter Size: 256
False Positive Rate: 1.000000, Filter Size: 512
False Positive Rate: 1.000000, Filter Size: 1024
False Positive Rate: 1.000000, Filter Size: 2048
False Positive Rate: 0.929000, Filter Size: 4096
False Positive Rate: 0.418000, Filter Size: 8192
False Positive Rate: 0.033000, Filter Size: 16384
False Positive Rate: 0.000000, Filter Size: 32768
False Positive Rate: 0.000000, Filter Size: 65536
False Positive Rate: 0.000000, Filter Size: 131072
False Positive Rate: 0.000000, Filter Size: 262144
```

#### False Positive Rate vs Number of Hash Function
```
False Positive Rate: 1.000000, Num of hash Functions: 1
False Positive Rate: 0.071000, Num of hash Functions: 2
False Positive Rate: 0.024000, Num of hash Functions: 3
False Positive Rate: 0.012000, Num of hash Functions: 4
False Positive Rate: 0.009000, Num of hash Functions: 5
False Positive Rate: 0.003000, Num of hash Functions: 6
False Positive Rate: 0.003000, Num of hash Functions: 7
False Positive Rate: 0.003000, Num of hash Functions: 8
False Positive Rate: 0.000000, Num of hash Functions: 9
False Positive Rate: 0.000000, Num of hash Functions: 10
False Positive Rate: 0.000000, Num of hash Functions: 11
False Positive Rate: 0.000000, Num of hash Functions: 12
False Positive Rate: 0.000000, Num of hash Functions: 13
False Positive Rate: 0.000000, Num of hash Functions: 14
False Positive Rate: 0.001000, Num of hash Functions: 15
False Positive Rate: 0.001000, Num of hash Functions: 16
False Positive Rate: 0.003000, Num of hash Functions: 17
False Positive Rate: 0.003000, Num of hash Functions: 18
False Positive Rate: 0.004000, Num of hash Functions: 19
False Positive Rate: 0.005000, Num of hash Functions: 20
False Positive Rate: 0.003000, Num of hash Functions: 21
False Positive Rate: 0.003000, Num of hash Functions: 22
False Positive Rate: 0.003000, Num of hash Functions: 23
False Positive Rate: 0.001000, Num of hash Functions: 24
False Positive Rate: 0.003000, Num of hash Functions: 25
False Positive Rate: 0.004000, Num of hash Functions: 26
False Positive Rate: 0.003000, Num of hash Functions: 27
False Positive Rate: 0.006000, Num of hash Functions: 28
False Positive Rate: 0.008000, Num of hash Functions: 29
False Positive Rate: 0.011000, Num of hash Functions: 30
False Positive Rate: 0.016000, Num of hash Functions: 31
False Positive Rate: 0.017000, Num of hash Functions: 32
False Positive Rate: 0.020000, Num of hash Functions: 33
False Positive Rate: 0.018000, Num of hash Functions: 34
False Positive Rate: 0.024000, Num of hash Functions: 35
False Positive Rate: 0.032000, Num of hash Functions: 36
False Positive Rate: 0.037000, Num of hash Functions: 37
False Positive Rate: 0.043000, Num of hash Functions: 38
False Positive Rate: 0.046000, Num of hash Functions: 39
False Positive Rate: 0.054000, Num of hash Functions: 40
False Positive Rate: 0.063000, Num of hash Functions: 41
False Positive Rate: 0.066000, Num of hash Functions: 42
False Positive Rate: 0.084000, Num of hash Functions: 43
False Positive Rate: 0.092000, Num of hash Functions: 44
False Positive Rate: 0.106000, Num of hash Functions: 45
False Positive Rate: 0.127000, Num of hash Functions: 46
False Positive Rate: 0.143000, Num of hash Functions: 47
False Positive Rate: 0.150000, Num of hash Functions: 48
False Positive Rate: 0.165000, Num of hash Functions: 49
False Positive Rate: 0.189000, Num of hash Functions: 50
False Positive Rate: 0.209000, Num of hash Functions: 51
False Positive Rate: 0.227000, Num of hash Functions: 52
False Positive Rate: 0.239000, Num of hash Functions: 53
False Positive Rate: 0.259000, Num of hash Functions: 54
False Positive Rate: 0.265000, Num of hash Functions: 55
False Positive Rate: 0.283000, Num of hash Functions: 56
False Positive Rate: 0.301000, Num of hash Functions: 57
False Positive Rate: 0.320000, Num of hash Functions: 58
False Positive Rate: 0.331000, Num of hash Functions: 59
False Positive Rate: 0.357000, Num of hash Functions: 60
False Positive Rate: 0.365000, Num of hash Functions: 61
False Positive Rate: 0.394000, Num of hash Functions: 62
False Positive Rate: 0.412000, Num of hash Functions: 63
False Positive Rate: 0.434000, Num of hash Functions: 64
False Positive Rate: 0.448000, Num of hash Functions: 65
False Positive Rate: 0.467000, Num of hash Functions: 66
False Positive Rate: 0.495000, Num of hash Functions: 67
False Positive Rate: 0.513000, Num of hash Functions: 68
False Positive Rate: 0.521000, Num of hash Functions: 69
False Positive Rate: 0.540000, Num of hash Functions: 70
False Positive Rate: 0.551000, Num of hash Functions: 71
False Positive Rate: 0.563000, Num of hash Functions: 72
False Positive Rate: 0.580000, Num of hash Functions: 73
False Positive Rate: 0.596000, Num of hash Functions: 74
False Positive Rate: 0.619000, Num of hash Functions: 75
False Positive Rate: 0.624000, Num of hash Functions: 76
False Positive Rate: 0.635000, Num of hash Functions: 77
False Positive Rate: 0.645000, Num of hash Functions: 78
False Positive Rate: 0.650000, Num of hash Functions: 79
False Positive Rate: 0.665000, Num of hash Functions: 80
False Positive Rate: 0.682000, Num of hash Functions: 81
False Positive Rate: 0.698000, Num of hash Functions: 82
False Positive Rate: 0.709000, Num of hash Functions: 83
False Positive Rate: 0.722000, Num of hash Functions: 84
False Positive Rate: 0.741000, Num of hash Functions: 85
False Positive Rate: 0.753000, Num of hash Functions: 86
False Positive Rate: 0.766000, Num of hash Functions: 87
False Positive Rate: 0.782000, Num of hash Functions: 88
False Positive Rate: 0.782000, Num of hash Functions: 89
False Positive Rate: 0.793000, Num of hash Functions: 90
False Positive Rate: 0.806000, Num of hash Functions: 91
False Positive Rate: 0.817000, Num of hash Functions: 92
False Positive Rate: 0.830000, Num of hash Functions: 93
False Positive Rate: 0.838000, Num of hash Functions: 94
False Positive Rate: 0.855000, Num of hash Functions: 95
False Positive Rate: 0.867000, Num of hash Functions: 96
False Positive Rate: 0.871000, Num of hash Functions: 97
False Positive Rate: 0.872000, Num of hash Functions: 98
False Positive Rate: 0.875000, Num of hash Functions: 99
False Positive Rate: 0.882000, Num of hash Functions: 100
```