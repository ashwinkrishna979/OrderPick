package database

import (
	"log"

	"github.com/gocql/gocql"
)

type DBConnection struct {
	Cluster *gocql.ClusterConfig
	Session *gocql.Session
}

func SetupDBConnection() (DBConnection, error) {
	var connection DBConnection
	connection.Cluster = gocql.NewCluster("127.0.0.1")
	connection.Cluster.Consistency = gocql.Quorum
	connection.Cluster.Keyspace = "order_pickup"
	session, err := connection.Cluster.CreateSession()

	if err != nil {
		log.Fatal("Failed to connect to Cassandra: ", err)
	}
	connection.Session = session
	log.Println("Connected to Cassandra")

	// Create user table if it doesn't exist
	if err := createUserTable(connection.Session); err != nil {
		log.Fatal("Failed to create user table:", err)
	}
	return connection, nil
}

func createUserTable(session *gocql.Session) error {
	query := `
    CREATE TABLE IF NOT EXISTS user (
        user_id UUID PRIMARY KEY,
        email TEXT,
        phone TEXT,
        password TEXT,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
        first_name TEXT,
        last_name TEXT,
        avatar TEXT,
        token_ TEXT,
        refresh_token TEXT
    )`
	return session.Query(query).Exec()
}