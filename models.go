package main

type (
	Employee struct {
		Id                int32
		Empno             string //员工号
		Name              string //员工姓名
		Gender            bool   //性别代码
		BirthDate         string //出生日期
		IdNumber          string //身份证号
		Nationality       string //民族
		Birthplace        string //出生地 
		PhoneNumber       string //电话号码
		WorkDate          string //工作时间
		Note              string //备注
		Address           string //地址
		Education         string //教育程度
		Email             string //邮箱
		ProfessionalTitle string //专业职称
	}

	Photo struct {
		Id         int32
		EmployeeId string //员工ID
		Photo      string //图片
	}

	Positiond struct {
		Id            int32  //
		EmployeeId    string //职位ID
		Team          string //所属机构
		Name          string //职位名称
		Date          string //
		PositionLevel string //岗位
		Level         string //
		OnTheJob      string //
		JobTitle      string //工作明细
	}
)

func (employee *Employee) ID() *int32 {
	return &employee.Id
}

func (employee *Employee) EMPNO() *string {
	return &employee.Empno
}

func (employee *Employee) GENDER() *bool {
	return &employee.Gender
}

func (employee *Employee) NAME() *string {
	return &employee.Name
}

func (employee *Employee) BIRTHDATE() *string {
	return &employee.BirthDate
}

func (employee *Employee) IDNUMBER() *string {
	return &employee.IdNumber
}

func (employee *Employee) NATIONALITY() *string {
	return &employee.Nationality
}

func (employee *Employee) BIRTHPLACE() *string {
	return &employee.Birthplace
}

func (employee *Employee) PHONENUMBER() *string {
	return &employee.PhoneNumber
}

func (employee *Employee) WORKDATE() *string {
	return &employee.WorkDate
}

func (employee *Employee) NOTE() *string {
	return &employee.Note
}

func (employee *Employee) ADDRESS() *string {
	return &employee.Address
}

func (employee *Employee) EDUCATION() *string {
	return &employee.Education
}

func (employee *Employee) EMAIL() *string {
	return &employee.Email
}

func (employee *Employee) PROFESSIONALTITLE() *string {
	return &employee.ProfessionalTitle
}

func (photos *Photo) ID() *int32 {
	return &photos.Id
}

func (photos *Photo) EMPLOYEEID() *string {
	return &photos.EmployeeId
}

func (photos *Photo) PHOTO() *string {
	return &photos.Photo
}

func (positiond *Positiond) ID() *int32 {
	return &positiond.Id
}

func (positiond *Positiond) EMPLOYEEID() *string {
	return &positiond.EmployeeId
}

func (positiond *Positiond) TEAM() *string {
	return &positiond.Team
}

func (positiond *Positiond) NAME() *string {
	return &positiond.Name
}

func (positiond *Positiond) DATE() *string {
	return &positiond.Date
}

func (positiond *Positiond) POSITIONLEVEL() *string {
	return &positiond.PositionLevel
}

func (positiond *Positiond) LEVEL() *string {
	return &positiond.Level
}

func (positiond *Positiond) ONTHEJOB() *string {
	return &positiond.OnTheJob
}

func (positiond *Positiond) JOBTITLE() *string {
	return &positiond.JobTitle
}

// func (photo *Photo) EMPLOYEE() (*Employee, error) {
// 	var employee Employee
// 	err := db.First(&employee, "id=?", photo.EmployeeId).Error
// 	return &employee, err
// }

// func (positiond *Positiond) EMPLOYEE() (*Employee, error) {
// 	var employee Employee
// 	err := db.First(&employee, "id=?", positiond.EmployeeId).Error
// 	return &employee, err
// }
