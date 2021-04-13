module github.com/Yossi-Dabush/simple-go-servic

go 1.16

require (
	dataBaseConn.com/db v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/websocket v1.4.2
	github.com/lib/pq v1.10.0 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	gorm.io/driver/postgres v1.0.8 // indirect
	gorm.io/gorm v1.21.6 // indirect

)

replace dataBaseConn.com/db => ../db

