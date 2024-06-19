package dao

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type csqlDb struct {
	session *gocql.Session
}

var DB csqlDb

func InitDB() {
	//find host with ipconfig refer read me
	cluster := gocql.NewCluster("172.30.64.1") //172.30.64.1
	cluster.Keyspace = "dev"                   // replace with your keyspace name
	cluster.Port = 9042
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		log.Fatal("Connection fail", err)
	}
	DB.session = session
	fmt.Println("DB connected successfully !")
}
