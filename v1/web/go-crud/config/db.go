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

var (
	// HEdb (MySQL), which is used to stroe HEApp users and records data.
	mysqldbAddr = "X"
	sqldbDriver = "X"
	sqldbUser   = "X"
	sqldbPass   = "X"
	SqldbName   = "X"

	// HEWebdb (MongoDB), which is used to store HEWeb users' information
	mongodbAddr     = "X"
	MongodbUser     = "X"
	MongodbPass     = "X"
	MongodbDatabase = "X"
)

func ConnectMongoDB() *mongo.Client {
	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s/", MongodbUser, MongodbPass, mongodbAddr)
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
