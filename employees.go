package main

import (
	"context"
	"fmt"
)

// type Employee struct {
// 	Id                int32
// 	Empno             string //员工号
// 	Name              string //员工姓名
// 	Gender            bool   //性别代码
// 	BirthDate         string //出生日期
// 	IdNumber          string //身份证号
// 	Nationality       string //民族
// 	Birthplace        string //出生地
// 	PhoneNumber       string //电话号码
// 	WorkDate          string //工作时间
// 	Note              string //备注
// 	Address           string //地址
// 	Education         string //教育程度
// 	Email             string //邮箱
// 	ProfessionalTitle string //专业职称
// }

// DB ===================================================================================
// GetPet should authorize the user in ctx and return a pet or error
func (db *DB) getEmployee(ctx context.Context, id int32) (*Employee, error) {
	var p Employee
	err := db.DB.First(&p, id).Error
	fmt.Println("+++++++++++")
	//fmt.Printf(id)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (db *DB) updateEmployee(ctx context.Context, args *employeeInput) (*Employee, error) {
	// get the pet to be updated from the db
	var p Employee
	err := db.DB.First(&p, args.Id).Error
	if err != nil {
		return nil, err
	}

	updated := Employee{
		Name:              args.Name,
		Empno:             args.Empno,
		Gender:            args.Gender,
		BirthDate:         args.BirthDate,
		IdNumber:          args.IdNumber,
		Nationality:       args.Nationality,
		Birthplace:        args.Birthplace,
		PhoneNumber:       args.PhoneNumber,
		WorkDate:          args.WorkDate,
		Note:              args.Note,
		Address:           args.Address,
		Education:         args.Education,
		Email:             args.Email,
		ProfessionalTitle: args.ProfessionalTitle}

	err = db.DB.Model(&p).Updates(updated).Error
	if err != nil {
		return nil, err
	}

	err = db.DB.First(&p, args.Id).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}
