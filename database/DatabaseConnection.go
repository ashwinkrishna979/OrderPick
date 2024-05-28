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

	if err := createItemTable(connection.Session); err != nil {
		log.Fatal("Failed to create item table:", err)
	}
	return connection, nil
}

func createUserTable(session *gocql.Session) error {
	createTableQuery := `
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

	if err := session.Query(createTableQuery).Exec(); err != nil {
		return err
	}
	createEmailIndexQuery := `CREATE INDEX IF NOT EXISTS ON user(email)`
	if err := session.Query(createEmailIndexQuery).Exec(); err != nil {
		return err
	}
	createPhoneIndexQuery := `CREATE INDEX IF NOT EXISTS ON user(phone)`
	if err := session.Query(createPhoneIndexQuery).Exec(); err != nil {
		return err
	}
	return nil
}

func createItemTable(session *gocql.Session) error {
	createTableQuery := `
	    CREATE TABLE IF NOT EXISTS item (
        item_id UUID PRIMARY KEY,
        name TEXT
    )
	`
	if err := session.Query(createTableQuery).Exec(); err != nil {
		return err
	}
	return nil
}
