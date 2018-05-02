CREATE DATABASE IF NOT EXISTS Course_Evaluation;
Use Course_Evaluation;
CREATE TABLE Student(
	Student_netID varchar(10) NOT NULL primary key,
	password varchar(100) NOT NULL,
	email varchar(100) NOT NULL
);

CREATE TABLE Professor(
	Professor_netID varchar(10) NOT NULL primary key, 
	password varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	Permission_Code varchar(10)
);

CREATE TABLE Course(
	CourseID varchar(10) NOT NULL primary key,
	Professor_netID varchar(10) NOT NULL, 
	
	Course_Name varchar(100) NOT NULL,
	Semester varchar(100) NOT NULL
);

CREATE TABLE StudentCourse(
	Course_Name varchar(100),
	Student_netID varchar(10),
	filled varchar(1)
);

CREATE TABLE Course_Response(
	Student_netID varchar(10) NOT NULL,
	Course_Rate varchar(10),
	Student_Comment_Advice varchar(1000),
	Student_Comment_Improve varchar(1000),
	Instructor_Rate varchar(2),
	Instructor_Goal varchar(10),  
	Instructor_Inspiration varchar(10),  
	Course_Comment varchar(1000),
	Instructor_Clarity varchar(10),
	Instructor_Feedback varchar(10), 
	Course_Content varchar(10),
	Course_Organize varchar(10),
	Course_Application varchar(10),
	Instructor_Comment varchar(1000),
	Course_Name varchar(100) NOT NULL 

);