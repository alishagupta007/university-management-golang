package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"university-management-golang/protoclient/university_management"
)

const (
	host = "localhost"
	port = "2345"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error occured %+v", err)
	}
	client := university_management.NewUniversityManagementServiceClient(conn)
	var departmentID int32 = 1
	var studentRollNo string = "145008"
	departmentResponse, err := client.GetDepartment(context.TODO(), &university_management.GetDepartmentRequest{Id: departmentID})

	studentsResponse, err := client.GetStudents(context.TODO(), &university_management.GetStudentsRequest{DepartmentId: departmentID})

	staffResponse, err := client.GetStaffForStudent(context.TODO(), &university_management.GetStaffForStudentRequest{RollNo: studentRollNo})
	if err != nil {
		log.Fatalf("Error occured while fetching department for id %d,err: %+v", departmentID, err)
	}
	log.Println(departmentResponse)
	log.Println(studentsResponse)
	log.Println(staffResponse)
}
