## Load Balancer with Backend Server Management

Simple load balancer that distributes requests among multiple backend servers. The load balancer adds new backend servers when the load on existing servers exceeds a specified threshold.

#### Endpoints

- **Load Balancer**: 
  - `localhost:8000/`: Main entry point that forwards requests to backend servers.
  - `localhost:8000/health`: Health check endpoint to retrieve the current load of all backend servers.

- **Backend Servers**: 
  - `localhost:8001/b_server1`, `localhost:8002/b_server2`, etc.: Response endpoints for each backend server.

- Simulate requests using a web browser or a tool like `curl`:
   ```bash
   curl http://localhost:8000/
   ```
#### Outputs
```
Starting load balancer on port 8000
2024/10/31 02:03:20 Starting backend server 3 on port 8003
2024/10/31 02:03:20 Starting backend server 4 on port 8004
2024/10/31 02:03:20 Starting backend server 1 on port 8001
2024/10/31 02:03:20 Starting backend server 2 on port 8002
```
Health Enpoint
![alt text](../images/lb-health.png)
Response
![alt text](../images/lb-response.png)