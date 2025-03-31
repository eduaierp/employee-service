

# Employee Service API - Setup & Documentation

This guide will walk you through setting up the **Employee Service API** project locally on both **Windows** and **Mac**, installing dependencies, setting up PostgreSQL, granting user permissions, and running the application.

---

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Setting Up PostgreSQL](#setting-up-postgresql)
3. [Setup Project Locally](#setup-project-locally)
   - [Windows](#windows-setup)
   - [Mac](#mac-setup)
4. [Install Dependencies](#install-dependencies)
5. [Grant User Permissions to the Table](#grant-user-permissions-to-the-table)
6. [Postman Request Bodies](#postman-request-bodies)
7. [Running the Application](#running-the-application)

---

## Prerequisites
1. **Go (Golang)**: The project is written in Go. Make sure Go is installed. You can install it from the [Go official website](https://golang.org/dl/).
2. **PostgreSQL**: The project uses PostgreSQL as the database.
3. **Postman**: A tool for testing the API (Optional for testing, but recommended).

---

## Setting Up PostgreSQL

### 1. Install PostgreSQL

**On Windows**:
1. Download PostgreSQL from [here](https://www.postgresql.org/download/windows/).
2. Run the installer and follow the installation steps.
3. After installation, open **pgAdmin** (PostgreSQL GUI) and connect to your database server.

**On Mac**:
1. Install PostgreSQL using Homebrew:
   ```bash
   brew install postgresql
   ```
2. Start PostgreSQL service:
   ```bash
   brew services start postgresql
   ```
3. Verify PostgreSQL is running:
   ```bash
   psql --version
   ```

---

## Setup Project Locally

### Windows Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/employee-service.git
   cd employee-service
   ```

2. **Install Go** (if not installed):
   Download and install Go from [here](https://golang.org/dl/), then verify:
   ```bash
   go version
   ```

3. **Run the Application**:
   - Set up PostgreSQL credentials and database URL in your `config.json` or `env` file.
   - To run the server, open a terminal/command prompt and run:
     ```bash
     go run main.go
     ```

### Mac Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/employee-service.git
   cd employee-service
   ```

2. **Install Go** (if not installed):
   Install Go via Homebrew:
   ```bash
   brew install go
   ```

3. **Run the Application**:
   - Set up PostgreSQL credentials and database URL in your `config.json` or `env` file.
   - To run the server, open a terminal and run:
     ```bash
     go run main.go
     ```

---

## Install Dependencies

1. Inside the project directory, run the following command to install dependencies:

   ```bash
   go mod tidy
   ```

2. This will download and install the required Go packages as listed in the `go.mod` file.

---

## Setting Up PostgreSQL Database

1. **Create a PostgreSQL Database**:
   If you haven't created a database already, you can create one using the following command:

   ```bash
   createdb employee_service
   ```

2. **Create Tables**:
   You can use the SQL scripts provided to create tables in PostgreSQL. Create the necessary table schema as per your `CREATE TABLE` query for `employees`.

   Example:
   ```sql
   CREATE TABLE employees (
       full_name VARCHAR(100),
       date_of_birth DATE,
       gender VARCHAR(10),
       nationality VARCHAR(50),
       aadhar_ssn VARCHAR(20),
       contact_number VARCHAR(20),
       email_id VARCHAR(100),
       address_permanent VARCHAR(255),
       address_current VARCHAR(255),
       emergency_contact VARCHAR(20),
       employee_id VARCHAR(50) PRIMARY KEY,
       joining_date DATE,
       department VARCHAR(50),
       designation VARCHAR(50),
       employment_type VARCHAR(20),
       work_shift VARCHAR(20),
       job_description TEXT,
       highest_qualification VARCHAR(100),
       specialization VARCHAR(100),
       teaching_experience INT,
       certifications TEXT[],
       research_publications TEXT[],
       professional_memberships TEXT[],
       roles_responsibilities TEXT[],
       reporting_manager VARCHAR(100),
       subjects_assigned TEXT[],
       class_grade_assigned TEXT[],
       salary_structure DECIMAL,
       bank_account_details VARCHAR(50),
       pan_number VARCHAR(50),
       provident_fund VARCHAR(50),
       health_insurance VARCHAR(50),
       educational_certificates TEXT[],
       leave_balance INT,
       leave_history TEXT[]
   );
   ```

3. **Create the User and Grant Permissions**:
   You need to create a PostgreSQL user for the application and grant them permission to the table. Run the following commands in your PostgreSQL CLI:

   ```sql
   CREATE USER your_user WITH PASSWORD 'your_password';
   GRANT ALL PRIVILEGES ON DATABASE employee_service TO your_user;
   ```

   Then, to give your user access to the tables:

   ```sql
   GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO your_user;
   ```

---

## Grant User Permissions to the Table

You can grant user permissions by running the following commands on PostgreSQL:

1. **Granting Select, Insert, Update, Delete Permissions**:
   ```sql
   GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE employees TO your_user;
   ```

2. **Granting Permissions for Sequence**:
   If you're using sequences (e.g., for auto-increment fields):
   ```sql
   GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO your_user;
   ```

---

## Postman Request Bodies

### 1. **Create Employee**

#### POST: `/employee`

```json
{
  "full_name": "John Doe",
  "date_of_birth": "1985-06-15",
  "gender": "Male",
  "nationality": "American",
  "aadhar_ssn": "123456789",
  "contact_number": "9876543210",
  "email_id": "john.doe@example.com",
  "address_permanent": "1234 Main St, Springfield",
  "address_current": "5678 Oak St, Springfield",
  "emergency_contact": "9876543211",
  "employee_id": "E12345",
  "joining_date": "2020-01-01",
  "department": "Engineering",
  "designation": "Software Engineer",
  "employment_type": "Full-Time",
  "work_shift": "Day",
  "job_description": "Develop software solutions.",
  "highest_qualification": "Bachelors in Computer Science",
  "specialization": "Software Development",
  "teaching_experience": "5",
  "certifications": ["AWS Certified", "Google Cloud Certified"],
  "research_publications": ["Paper 1", "Paper 2"],
  "professional_memberships": ["IEEE", "ACM"],
  "roles_responsibilities": ["Develop applications", "Mentor juniors"],
  "reporting_manager": "Jane Smith",
  "subjects_assigned": ["Software Engineering", "Data Structures"],
  "class_grade_assigned": ["Grade 10", "Grade 12"],
  "salary_structure": "50000",
  "bank_account_details": "1234567890123456",
  "pan_number": "ABCDE1234F",
  "provident_fund": "PF1234",
  "health_insurance": "HealthIns123",
  "educational_certificates": ["BSc Certificate", "MSc Certificate"],
  "leave_balance": "20",
  "leave_history": ["Sick Leave", "Vacation"]
}
```

---

### 2. **Get Employee by ID**

#### POST: `/employee/{id}`

```json
{
  "employee_id": "E12345"
}
```

---

### 3. **Get All Employees**

#### POST: `/employees`

```json
{}
```

---

### 4. **Update Employee**

#### PUT: `/employee/{id}`

```json
{
  "employee_id": "E12345",
  "full_name": "John Doe Updated",
  "date_of_birth": "1985-06-15",
  "gender": "Male",
  "nationality": "American",
  "aadhar_ssn": "123456789",
  "contact_number": "9876543210",
  "email_id": "john.doe@example.com",
  "address_permanent": "1234 Main St, Springfield",
  "address_current": "5678 Oak St, Springfield",
  "emergency_contact": "9876543211",
  "joining_date": "2020-01-01",
  "department": "Engineering",
  "designation": "Software Engineer",
  "employment_type": "Full-Time",
  "work_shift": "Day",
  "job_description": "Develop software solutions.",
  "highest_qualification": "Bachelors in Computer Science",
  "specialization": "Software Development",
  "teaching_experience": "5",
  "certifications": ["AWS Certified", "Google Cloud Certified"],
  "research_publications": ["Paper 1", "Paper 2"],
  "professional_memberships": ["IEEE", "ACM"],
  "roles_responsibilities": ["Develop applications", "Mentor juniors"],
  "reporting_manager": "Jane Smith",
  "subjects_assigned": ["Software Engineering", "Data Structures"],
  "class_grade_assigned": ["Grade 10", "Grade 12"],
  "salary_structure": "50000

",
  "bank_account_details": "1234567890123456",
  "pan_number": "ABCDE1234F",
  "provident_fund": "PF1234",
  "health_insurance": "HealthIns123",
  "educational_certificates": ["BSc Certificate", "MSc Certificate"],
  "leave_balance": "20",
  "leave_history": ["Sick Leave", "Vacation"]
}
```

---

### 5. **Delete Employee**

#### DELETE: `/employee/{id}`

```json
{
  "employee_id": "E12345"
}
```

---

