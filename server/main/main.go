package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
	migrations "university-management-golang/db"
	"university-management-golang/db/connection"
	um "university-management-golang/protoclient/university_management"
	"university-management-golang/server/internal/handlers"
)

const port = "2345"

//db
const (
	username = "postgres"
	password = "admin"
	host = "localhost"
	dbPort   = "5436"
	dbName = "postgres"
	schema = "public"
)

type Student struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	RollNo       string `db:"rollno"`
	DepartmentId int64  `db:"departmentid"`
}

func main() {
	err := migrations.MigrationsUp(username, password, host, dbPort, dbName, schema)
	if err != nil {
		log.Fatalf("Failed to migrate, err: %+v\n", err)
	}

	connectionmanager := &connection.DatabaseConnectionManagerImpl{
			&connection.DBConfig{
				host,dbPort,username,password,dbName,schema,
			},
			nil,
	}

	//insertSeedData(connectionmanager)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen to port: %s, err: %+v\n", port, err)
	}
	log.Printf("Starting to listen on port: %s\n", port)

	um.RegisterUniversityManagementServiceServer(grpcServer, handlers.NewUniversityManagementHandler(connectionmanager))
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to start GRPC Server: %+v\n", err)
	}
}

func insertSeedData(connectionManager connection.DatabaseConnectionManager) {
	connection, err := connectionManager.GetConnection()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	log.Println("Cleaning up department table")
	//_, err = connection.GetSession().DeleteFrom("departments").Exec()
	if err != nil {
		log.Fatalf("Could not delete from department table. Err: %+v", err)
	}

	log.Println("Cleaning up students table")
	//_, err = connection.GetSession().DeleteFrom("students").Exec()
	if err != nil {
		log.Fatalf("Could not delete from students table. Err: %+v", err)
	}
	log.Println("Inserting into department table")
	//_, err = connection.GetSession().InsertInto("departments").Columns("name").
	//	Values("Computer Science").Exec()

	log.Println("Inserting into students table")
	//studentData := &Student{
	//	Name:         "Rahul",
	//	RollNo:       "145008",
	//	DepartmentId: 1,
	//}
	//_, err = connection.GetSession().InsertInto("students").Columns("name", "rollno", "departmentid").Record(studentData).Exec()
	//
	//_, err = connection.GetSession().InsertInto("staff").Columns("name").Values("Staff1").Exec()
	//_, err = connection.GetSession().InsertInto("staff").Columns("name").Values("Staff2").Exec()
	//_, err = connection.GetSession().InsertInto("staff").Columns("name").Values("Staff3").Exec()
	//_, err = connection.GetSession().InsertInto("department_staff").Columns("department_id", "staff_id").Values(4, 2).Exec()
	//_, err = connection.GetSession().InsertInto("department_staff").Columns("department_id", "staff_id").Values(4, 3).Exec()
	dt := time.Now().UTC()

	_, err = connection.GetSession().InsertInto("student_attendance").Columns("student_id", "login_time", "date").Values(4,dt.Format("03:04:05"), dt.Format("2006-01-02") ).Exec()


	if err != nil {
		log.Fatalf("Could not insert into department table. Err: %+v", err)
	}

	defer connectionManager.CloseConnection()
}
