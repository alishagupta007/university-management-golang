package handlers

import (
	"context"
	"encoding/json"
	"log"
	"university-management-golang/db/connection"
	um "university-management-golang/protoclient/university_management"
)

type universityManagementServer struct {
	um.UniversityManagementServiceServer
	connectionManager connection.DatabaseConnectionManager
}

func (u *universityManagementServer) GetDepartment(ctx context.Context, request *um.GetDepartmentRequest) (*um.GetDepartmentResponse, error) {
	connection, err := u.connectionManager.GetConnection()

	//defer u.connectionManager.CloseConnection()

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	var department um.Department
	connection.GetSession().Select("id", "name").From("department").Where("id = ?", request.GetId()).LoadOne(&department)

	_, err = json.Marshal(&department)
	 if err != nil {
		 log.Fatalf("Error while marshaling %+v", err)
	 }

	return &um.GetDepartmentResponse{Department: &um.Department{
		Id:   department.Id,
		Name: department.Name,
	}}, nil
}

func NewUniversityManagementHandler(connectionmanager connection.DatabaseConnectionManager) um.UniversityManagementServiceServer {
	return &universityManagementServer {
		connectionManager: connectionmanager,
	}
}

func (u *universityManagementServer) GetStudents(ctx context.Context, request *um.GetStudentsRequest) (*um.GetStudentsResponse, error) {
	connection, err := u.connectionManager.GetConnection()

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	var students []um.Student

	connection.GetSession().Select("id", "name", "rollno").From("students").Where("departmentid = ?", request.GetDepartmentId()).Load(&students)

	_, err = json.Marshal(&students)
	if err != nil {
		log.Fatalf("Error while marshaling %+v", err)
	}

	var finalStudents []*um.Student
	for _,val := range  students{
		var temp = um.Student{
			Id: val.Id,
			Name: val.Name,
			RollNo: val.RollNo,
			DepartmentId: val.DepartmentId,
		}
		finalStudents = append(finalStudents,&temp)
	}
	return &um.GetStudentsResponse{Student: finalStudents}, nil
}

func (u *universityManagementServer) GetStaffForStudent(ctx context.Context, request *um.GetStaffForStudentRequest) (*um.GetStaffForStudentResponse, error)  {
	connection, err := u.connectionManager.GetConnection()

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	var staff []um.Staff

	connection.GetSession().Select("department_staff.staff_id as id", "staff.name").From("students").Join("departments", "departments.id=students.departmentid").Join("department_staff", "departments.id=department_staff.department_id").Join("staff", "staff.id = department_staff.staff_id").Where("rollno = ?", request.GetRollNo()).Load(&staff)

	_, err = json.Marshal(&staff)
	if err != nil {
		log.Fatalf("Error while marshaling %+v", err)
	}

	var staffMembers []*um.Staff
	for _,val := range  staff{
		var temp = um.Staff{
			Id: val.Id,
			Name: val.Name,
		}
		staffMembers = append(staffMembers,&temp)
	}
	return &um.GetStaffForStudentResponse{Staff: staffMembers},err
}
