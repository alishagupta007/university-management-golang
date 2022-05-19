package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
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

	var staffMembers []*um.Staff
	for _,val := range  staff{
		staffMembers = append(staffMembers,&val)
	}
	return &um.GetStaffForStudentResponse{Staff: staffMembers},err
}

func (u *universityManagementServer) Login(ctx context.Context, request *um.LoginRequest) (*um.LoginResponse, error) {
	defer recoverFunc()

	connection, err := u.connectionManager.GetConnection()

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	_,err = connection.GetSession().InsertBySql(
		`INSERT INTO student_attendance(student_id, login_time, date)
				VALUES(?,CURRENT_TIME, CURRENT_DATE) 
				ON CONFLICT ON CONSTRAINT student_date_unique
				DO UPDATE SET login_time=EXCLUDED.login_time
				WHERE student_attendance.login_time IS NULL`,
				request.GetId(),
		).Exec()

	if err != nil {
		panic("runtime error: Error trying to login, incorrect student id passed")
	}

	return &um.LoginResponse{
		Message: "Login successful",
	}, err
}

func recoverFunc() {
	if r := recover(); r!= nil {
		log.Println("Recovered from failure")
	}
}

func (u *universityManagementServer) Logout(ctx context.Context, request *um.LogoutRequest) (*um.LogoutResponse, error) {
	defer recoverFunc()

	connection, err := u.connectionManager.GetConnection()

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	_, err =  connection.GetSession().InsertBySql(
		`INSERT INTO student_attendance(student_id, logout_time, date)
				VALUES(?,CURRENT_TIME, CURRENT_DATE) 
				ON CONFLICT ON CONSTRAINT student_date_unique
				DO UPDATE  SET logout_time=EXCLUDED.logout_time
               `,request.GetId(),
	).Exec()

	if err != nil {
		panic("runtime error: Error trying to logout, incorrect student id passed")
	}

	return &um.LogoutResponse{
		Message: "Logout Successful",
	}, err
}

func (u *universityManagementServer) Notify(ctx context.Context, request *um.GetNotifyRequest) (*um.GetNotifyResponse, error) {
	channel := make(chan string)

	go u.waitForStudentToLogin(request.GetId(), channel)

	message := <-channel

	return &um.GetNotifyResponse{Message: message}, nil
}

func (u *universityManagementServer) waitForStudentToLogin(ids []int32, channel chan string) {
	connection, err := u.connectionManager.GetConnection()

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	for {
		time.Sleep(1 * time.Second)
        var userLogin struct{
        	StudentId string `db:"student_id,omitempty"`
		}

		if len(ids) != 0 {
			connection.GetSession().Select("student_id").From("student_attendance").Where("student_id IN (?)", ids).LoadOne(&userLogin)
		}else {
			connection.GetSession().Select("student_id").From("student_attendance").LoadOne(&userLogin)
		}

		if userLogin.StudentId != "" {
			channel <- fmt.Sprintf("User %s has logged in", userLogin.StudentId)
			return
		}
	}
}

