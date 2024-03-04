package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// HEdb (MySQL), which is used to stroe HEApp users and records data.
	mysqldbAddr = "140.113.151.61:3306"
	sqldbDriver = "mysql"
	sqldbUser   = "rtes913"
	sqldbPass   = "MYSQLrtes913"
	SqldbName   = "HEdb"

	// HEWebdb (MongoDB), which is used to store HEWeb users' information
	mongodbURI        = "mongodb://140.113.151.61:27017"
	MongodbDatabase   = "HEWEBdb"
	mongodbCollection = "admin"
)

func ConnectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongodbURI)
	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB %v, %v", MongodbDatabase, err)
	}

	// Check the MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB %v, %v", MongodbDatabase, err)
	}

	log.Printf("Connected to MongoDB %v", MongodbDatabase)
	return client
}

func ConnectMysqlDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", sqldbUser, sqldbPass, mysqldbAddr, SqldbName)
	mysqldb, err := sql.Open(sqldbDriver, dsn)
	if err != nil {
		log.Fatalf("Error Connecting to MySQL %v, %v", SqldbName, err)
	}
	err = mysqldb.Ping()
	if err != nil {
		log.Fatalf("Error Pinging MySQL %v, %v", SqldbName, err)
	}

	log.Printf("Connected to MySQL %v", SqldbName)
	return mysqldb
}
