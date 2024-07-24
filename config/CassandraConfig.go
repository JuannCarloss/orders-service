package config

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	Session *gocql.Session
)

func InitCluster() (*gocql.Session, error) {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "orderms"
	Session, err := cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	fmt.Println("success")

	return Session, nil

}
