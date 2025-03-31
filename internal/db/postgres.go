package db

import (
	"database/sql"
	"employee-service/proto"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ✅ Initialize Database Connection
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return DB.Ping()
}

// ✅ Create Employee
func CreateEmployee(employee *proto.Employee) error {
	query := `
		INSERT INTO employees (
			full_name, date_of_birth, gender, nationality, aadhar_ssn, 
			contact_number, email_id, address_permanent, address_current, 
			emergency_contact, employee_id, joining_date, department, designation, 
			employment_type, work_shift, job_description, highest_qualification, 
			specialization, teaching_experience, certifications, research_publications, 
			professional_memberships, roles_responsibilities, reporting_manager, 
			subjects_assigned, class_grade_assigned, administrative_duties, 
			committee_memberships, salary_structure, bank_account_details, pan_number, 
			provident_fund, health_insurance, educational_certificates, 
			experience_certificates, government_id_proofs, address_proofs, 
			joining_letter, passport_photos, performance_reviews, student_feedback_ratings, 
			training_programs, promotions_salary_revisions, attendance_records, 
			leave_balance, leave_history, leave_approval_workflow, employee_category
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
			$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
			$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, 
			$31, $32, $33, $34, $35, $36, $37, $38, $39, $40, 
			$41, $42, $43, $44, $45, $46, $47, $48, $49
		)
	`

	_, err := DB.Exec(query,
		employee.FullName, employee.DateOfBirth, employee.Gender, employee.Nationality, employee.AadharSsn,
		employee.ContactNumber, employee.EmailId, employee.AddressPermanent, employee.AddressCurrent,
		employee.EmergencyContact, employee.EmployeeId, employee.JoiningDate, employee.Department,
		employee.Designation, employee.EmploymentType, employee.WorkShift, employee.JobDescription,
		employee.HighestQualification, employee.Specialization, employee.TeachingExperience,
		employee.Certifications, employee.ResearchPublications, employee.ProfessionalMemberships,
		employee.RolesResponsibilities, employee.ReportingManager, employee.SubjectsAssigned,
		employee.ClassGradeAssigned, employee.AdministrativeDuties, employee.CommitteeMemberships,
		employee.SalaryStructure, employee.BankAccountDetails, employee.PanNumber, employee.ProvidentFund,
		employee.HealthInsurance, employee.EducationalCertificates, employee.ExperienceCertificates,
		employee.GovernmentIdProofs, employee.AddressProofs, employee.JoiningLetter, employee.PassportPhotos,
		employee.PerformanceReviews, employee.StudentFeedbackRatings, employee.TrainingPrograms,
		employee.PromotionsSalaryRevisions, employee.AttendanceRecords, employee.LeaveBalance,
		employee.LeaveHistory, employee.LeaveApprovalWorkflow, employee.EmployeeCategory,
	)

	return err
}

// ✅ Get All Employees
func GetAllEmployees() ([]*proto.Employee, error) {
	query := `SELECT * FROM employees`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not fetch employees: %v", err)
	}
	defer rows.Close()

	var employees []*proto.Employee
	for rows.Next() {
		var employee proto.Employee
		var leaveBalance sql.NullString

		err := rows.Scan(
			&employee.FullName, &employee.DateOfBirth, &employee.Gender, &employee.Nationality, &employee.AadharSsn,
			&employee.ContactNumber, &employee.EmailId, &employee.AddressPermanent, &employee.AddressCurrent,
			&employee.EmergencyContact, &employee.EmployeeId, &employee.JoiningDate, &employee.Department,
			&employee.Designation, &employee.EmploymentType, &employee.WorkShift, &employee.JobDescription,
			&employee.HighestQualification, &employee.Specialization, &employee.TeachingExperience,
			&employee.Certifications, &employee.ResearchPublications, &employee.ProfessionalMemberships,
			&employee.RolesResponsibilities, &employee.ReportingManager, &employee.SubjectsAssigned,
			&employee.ClassGradeAssigned, &employee.AdministrativeDuties, &employee.CommitteeMemberships,
			&employee.SalaryStructure, &employee.BankAccountDetails, &employee.PanNumber, &employee.ProvidentFund,
			&employee.HealthInsurance, &employee.EducationalCertificates, &employee.ExperienceCertificates,
			&employee.GovernmentIdProofs, &employee.AddressProofs, &employee.JoiningLetter, &employee.PassportPhotos,
			&employee.PerformanceReviews, &employee.StudentFeedbackRatings, &employee.TrainingPrograms,
			&employee.PromotionsSalaryRevisions, &employee.AttendanceRecords, &leaveBalance,
			&employee.LeaveHistory, &employee.LeaveApprovalWorkflow, &employee.EmployeeCategory,
		)
		if err != nil {
			continue
		}

		employees = append(employees, &employee)
	}

	return employees, nil
}

// ✅ Get Employee by ID
func GetEmployeeByID(employeeID string) (*proto.Employee, error) {
	query := `SELECT * FROM employees WHERE employee_id = $1`
	row := DB.QueryRow(query, employeeID)

	var employee proto.Employee
	err := row.Scan(&employee)
	if err != nil {
		return nil, fmt.Errorf("could not fetch employee: %v", err)
	}

	return &employee, nil
}

// ✅ Update Employee
func UpdateEmployee(employeeID string, updatedEmployee *proto.Employee) error {
	query := `
		UPDATE employees SET 
			full_name = $1, date_of_birth = $2, gender = $3, nationality = $4, 
			aadhar_ssn = $5, contact_number = $6, email_id = $7, address_permanent = $8, 
			address_current = $9, emergency_contact = $10, joining_date = $11, 
			department = $12, designation = $13, employment_type = $14, work_shift = $15, 
			job_description = $16, highest_qualification = $17, specialization = $18, 
			teaching_experience = $19, certifications = $20, research_publications = $21, 
			professional_memberships = $22, roles_responsibilities = $23, 
			reporting_manager = $24, subjects_assigned = $25, class_grade_assigned = $26, 
			administrative_duties = $27, committee_memberships = $28, salary_structure = $29, 
			bank_account_details = $30, pan_number = $31, provident_fund = $32, 
			health_insurance = $33, leave_balance = $34 
		WHERE employee_id = $35
	`

	_, err := DB.Exec(query,
		updatedEmployee.FullName, updatedEmployee.DateOfBirth, updatedEmployee.Gender, updatedEmployee.Nationality,
		updatedEmployee.AadharSsn, updatedEmployee.ContactNumber, updatedEmployee.EmailId, updatedEmployee.AddressPermanent,
		updatedEmployee.AddressCurrent, updatedEmployee.EmergencyContact, updatedEmployee.JoiningDate,
		updatedEmployee.Department, updatedEmployee.Designation, updatedEmployee.EmploymentType, updatedEmployee.WorkShift,
		updatedEmployee.JobDescription, updatedEmployee.HighestQualification, updatedEmployee.Specialization,
		updatedEmployee.TeachingExperience, updatedEmployee.Certifications, updatedEmployee.ResearchPublications,
		updatedEmployee.ProfessionalMemberships, updatedEmployee.RolesResponsibilities, updatedEmployee.ReportingManager,
		updatedEmployee.SubjectsAssigned, updatedEmployee.ClassGradeAssigned, updatedEmployee.AdministrativeDuties,
		updatedEmployee.CommitteeMemberships, updatedEmployee.SalaryStructure, updatedEmployee.BankAccountDetails,
		updatedEmployee.PanNumber, updatedEmployee.ProvidentFund, updatedEmployee.HealthInsurance, updatedEmployee.LeaveBalance,
		employeeID,
	)

	return err
}

// ✅ Delete Employee
func DeleteEmployee(employeeID string) error {
	query := `DELETE FROM employees WHERE employee_id = $1`
	_, err := DB.Exec(query, employeeID)
	return err
}
