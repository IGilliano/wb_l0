#### WB Tech: level # 0

Service allows:
- Subscribe to NATS-streaming server to recieve orders
- Save received order in database and cache
- Receive order by ID from cache

####Usage:

# Launch STAN server:
go run ./cmd/stan_pub

# Launch service: 
go run ./cmd/l0

# launch STAN server:
go run ./cmd/stan_pub

# Interface:
http://localhost:8080/docs/index.html#
