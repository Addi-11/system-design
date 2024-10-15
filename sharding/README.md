## Simulating Data Sharding and Routing from API Server
1. Create a DB.<br>
Run: `mysql -u root -p < <path to set_db.sql>`
2. Simulate Multiple Shards: Create 2 connections from our machine to the DB, simulating 2 different shards.
3. API Server Routing: our API server is aware of th DB topolgy, `r.GET("/user/:userID", getUser)`.
The `getUser` function determines the appropiate shard based on the userID
<br>
4. **Alternative**: Instead of handling routing at the application layer, we could be using a Proxy Database. A proxy DB handles the routing and scaling behind the scenes, and automatically distribute traffic across different shards.

#### Reference
- [Sample Implementation of Gin API](../golang-prerequisites/gin/main.go)
- [Go-Gin developer Guide](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)