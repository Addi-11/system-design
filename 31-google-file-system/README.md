## Implementing Google File system

create 1 master server
create 3 chunk servers (primary) - and 1 replica each
Replication factor 2

File 1 = A + B
File 2 = C + D

A B D
A C D
B C

gfs master functions:
- has metadata on which chunk servers does the data exsist
- assign unique id to chunk
- send heartbeat to chunk to monitor health.
    - maintains replication factor of 2. (Wont do this now. - but keep code extensible)

- identify chunk servers (primary + replica) with data, and send their address.

gfs client functions:
- turn's user request to chunk and offset
- goes to chunk ip address to recieve data, 
    - if data not here goes to next chunk

- send's user the data


READ
1) client -> gfs_client -> gfs master -> chunk IP addresses
2) GET DATA: gfs_client -> chunk

WRITE
1) client -> gfs_client -> gfs master -> chunk IP addresses
2) SEND DATA: gfs_client -> chunk(keep primary) -> chunk replica
                            - primary commits order




