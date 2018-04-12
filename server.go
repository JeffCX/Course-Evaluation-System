package main

import (
	"fmt"
	"net/http"
	"html/template"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	
)



var session = make(map[string]string)

var Username string =""
var Password string =""

func check_session( users string ,password string ) bool {
	fmt.Println(users)
	fmt.Println(password)
	fmt.Println(session[users+password])
	if (!(session[users+password]=="")){
		return true
	} else {
		return false
	}
}



type NewsAggPage struct {
    Title string
    News string
}



func check_code(code string) bool {
	lst:=[]string{"1","2","3","4","5","6","7","8","9","10"}
	for i:=0;i<len(lst);i++ {
		if code==lst[i] {
			return true
		}
	}
	return false
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if (session["status"]=="Invalid"){
		p := NewsAggPage{Title: "Invalid credientals", News: "some news"}
	    t, _ := template.ParseFiles("index.html")
	    t.Execute(w, p)
	    session["status"]="temp"
	}else if (session["status"]=="WrongCode"){
		p := NewsAggPage{Title: "Incorrect Permission Code", News: "some news"}
	    t, _ := template.ParseFiles("index.html")
	    t.Execute(w, p)
	    session["status"]="temp"
	}else if(session["status"]=="RepeatName"){
		p := NewsAggPage{Title: "Users name already taken", News: "some news"}
	    t, _ := template.ParseFiles("index.html")
	    t.Execute(w, p)
	    session["status"]="temp"

	}else if (session["status"]=="logOut"){
		p := NewsAggPage{Title: "You are logout", News: "some news"}
	    t, _ := template.ParseFiles("index.html")
	    t.Execute(w, p)
	    session["status"]="temp"
	}else if (session["status"]=="LoginFirst"){
		p := NewsAggPage{Title: "You need to login first!", News: "some news"}
	    t, _ := template.ParseFiles("index.html")
	    t.Execute(w, p)
	}else{
	    p := NewsAggPage{Title: "", News: "some news"}
	    t, _ := template.ParseFiles("index.html")
	    t.Execute(w, p)
	}
}



func studentHandler(w http.ResponseWriter, r *http.Request) {
	if(check_session(Username,Password)) {
		p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	    t, _ := template.ParseFiles("student.html")
	    t.Execute(w, p)
	}else {
		session["status"]="LoginFirst"
		http.Redirect(w,r,"/index/",http.StatusSeeOther)
	}
	
	

	// Print secret message

}

func professorHandler(w http.ResponseWriter, r *http.Request){

    if(check_session(Username,Password)) {
		p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	    t, _ := template.ParseFiles("professor_try.html")
	    t.Execute(w, p)
	}else {
		session["status"]="LoginFirst"
		http.Redirect(w,r,"/index/",http.StatusSeeOther)
	}
	
}

func submitHandler(w http.ResponseWriter, r *http.Request){
	
		p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
    t, _ := template.ParseFiles("submit.html")
    t.Execute(w, p)

	
}

//helper functions 
func HashPassword(password string) (string , error) {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes),err
}

func CheckPasswordHash(password,hash string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}




func login(e http.ResponseWriter, r *http.Request){
	//get data,if method is not post, redirect to login page

	//select data from db, if null, redirect to login page, show invalid credientals

	//else: if password is right, redirect to studentpage, else, redirect to login page, show invalid credientals

	//use js to check the data and make sure it is not empty
    control:=true
	if r.Method=="POST" {
		r.ParseForm()
		//db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
		db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
		rows,_:=db.Query("SELECT Student_netID, password FROM Student where Student_netID=?;",r.Form["username"][0])
		var username string 
		var password string 
		for rows.Next(){
			rows.Scan(&username,&password)
		}

		if CheckPasswordHash(r.Form["password"][0],password){
		    Username = r.Form["username"][0]
		    Password = r.Form["password"][0]
		    session[Username+Password] =Username+Password
			http.Redirect(e,r,"/student/",http.StatusSeeOther)
			control=false
				
		}
		db.Close()

		if (control){fmt.Println("gan")
		dbs, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
		row,_:=dbs.Query("SELECT Professor_netID, password FROM Professor where Professor_netID=?;",r.Form["username"][0])
		var Pusername string 
		var Ppassword string 
		for row.Next(){
			row.Scan(&Pusername,&Ppassword)
		}
		

		if CheckPasswordHash(r.Form["password"][0],Ppassword){
			Username = r.Form["username"][0]
		    Password = r.Form["password"][0]
		    session[Username+Password] =Username+Password
		    
		    http.Redirect(e,r,"/professor/",http.StatusSeeOther)
				
		}else {
			fmt.Println("hello")
			session["status"]="Invalid"
			http.Redirect(e,r,"/index/",http.StatusSeeOther)
		    
				
			//"wrong password"
		}
		
		dbs.Close()}
		

}
}	


func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var Student_netID string
		//db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
		db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
		rows,_:=db.Query("SELECT Student_netID FROM Student where Student_netID=?",r.Form["username_reg"][0])
		row,_:=db.Query("SELECT Professor_netID FROM Professor where Professor_netID=?",r.Form["username_reg"][0])
		fmt.Println("register")
		fmt.Println(r.Form)
		
		if rows.Next() == false && row.Next() == false {
			for rows.Next() {
          		rows.Scan(&Student_netID)
        	} 
        	pass,_:=HashPassword(r.Form["password_reg"][0]) 
        	if r.Form["Permission_code"][0]=="" {
        			stmt,_:=db.Prepare("insert into Student VALUES(?,?,?)")
        			stmt.Exec(r.Form["username_reg"][0],pass,r.Form["email_reg"][0])
				    Username = r.Form["username_reg"][0]
				    Password = r.Form["password_reg"][0]
				    session[Username+Password] =Username+Password
        			http.Redirect(w,r,"/student/",http.StatusSeeOther)//success sign up 
        		}else{
        			
        			 if check_code(r.Form["Permission_code"][0]){
        			 	  stmt,_:=db.Prepare("insert into  Professor VALUES(?,?,?,?)")
       						stmt.Exec(r.Form["username_reg"][0],pass,r.Form["email_reg"][0],r.Form["Permission_code"][0])
       						Username = r.Form["username_reg"][0]
						    Password = r.Form["password_reg"][0]
						    session[Username+Password] =Username+Password
						    
       						http.Redirect(w,r,"/professor/",http.StatusSeeOther)
        			 	}else{
        			 		session["status"]="WrongCode"
        			 		http.Redirect(w,r,"/index/",http.StatusSeeOther) 
        			 		//invalid permission code
        			 	}
        		}
			}else {
				session["status"]="RepeatName"
				http.Redirect(w,r,"/index/",http.StatusSeeOther)
			} 
				//already taken choose another one
	
	//get the data, connect to database

	//search data base, if in database,redirect the login,invalid credientals

	//else: insert into database 
}
}

//create class



func convert_to_string (lst []string) string{
	result :=""
	for i:=0;i<len(lst);i++{
		result+=lst[i]
	}
	return result;
}

func finish_eval(w http.ResponseWriter,r *http.Request){
	//some javascript check

	//if method is post,
	if r.Method == "POST" {
		r.ParseForm()
		//db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
		db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")

		
		CourseID:="123"
		IPresentation:=convert_to_string(r.Form["Instructor_Content[]"])
		IClarity:=convert_to_string(r.Form["Instructor_Clarity[]"])
		IHelpfulness:=convert_to_string(r.Form["Instructor_Availability[]"])
		IFeedback:=convert_to_string(r.Form["Instructor_Feedback[]"])
		IInsipration:=convert_to_string(r.Form["Instructor_Inspiration[]"])
		Instructor_rate := r.Form["Instructor_Rate"][0]
		Instructor_comment:=r.Form["Instructor_Comment"][0]
		CContent:=convert_to_string(r.Form["Course_Content[]"])
		CSkill:=convert_to_string(r.Form["Course_Application[]"])
		CTheory:=convert_to_string(r.Form["Course_Theory[]"])
		Course_rate:=r.Form["Course_Rate"][0]	
		Course_comment:=r.Form["Course_Comment"][0]
		Student_Comment_advice:=r.Form["Student_Comment_advice"][0]
		Student_Comment_improve:=r.Form["Student_Comment_improve"][0]	
		//two more 	
 		stmt,_:=db.Prepare("insert into Course_Response (CourseID,IPresentation,IClarity,IHelpfulness,IFeedback,IInsipration,Instructor_rate,Instructor_comment ,CContent ,CSkill,CTheory,Course_rate  ,Course_comment , Student_Comment_improve, Student_Comment_advice) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
		stmt.Exec(CourseID,IPresentation,IClarity,IHelpfulness,IFeedback,IInsipration,Instructor_rate,Instructor_comment,CContent,CSkill,CTheory,Course_rate,Course_comment,Student_Comment_advice,Student_Comment_improve)
		db.Close()	
		http.Redirect(w,r,"/submit/",http.StatusSeeOther)
	
	}

	// insert into dabase
	//redirect to submit page
}
//error handling 


func get_data(w http.ResponseWriter,r *http.Request){
	//db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
	db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")

	stmt,_:=db.Query("select * from Course_Response where CourseID=?","123")
	var CCourseID string 
	var IIIInsipration string 
	var IIClarity string 
	var IIHelpfulness string 
	var IIFeedback string 
	var IIInsipration string 
	var IInstructor_rate string 
	var IInstructor_comment string 
	var CCContent string 
	var CCSkill string 
	var CCTheory string 
	var CCourse_rate string 
	 var CCourse_comment string 
	var SStudent_Comment_improve string 
	var SStudent_Comment_advice string 
	for stmt.Next(){
		stmt.Scan(&CCourseID,&IIIInsipration,&IIClarity,&IIHelpfulness,&IIFeedback,&IIInsipration,&IInstructor_rate,&IInstructor_comment,&CCContent,&CCSkill,&CCTheory,&CCourse_rate,&CCourse_comment,&SStudent_Comment_improve,&SStudent_Comment_advice)
	}
	data:=[]string {CCourseID,",",IIIInsipration,",",IIClarity,",",IIHelpfulness,",",IIFeedback,",",IIInsipration,",",IInstructor_rate,",",IInstructor_comment,"," ,CCContent,",",CCSkill,",",CCTheory,",",CCourse_rate,",",CCourse_comment,",",SStudent_Comment_improve,",",SStudent_Comment_advice }
	fmt.Fprintln(w,data)
}

func logout(w http.ResponseWriter,r *http.Request){
	//db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
	Password=""
	Username=""
	session["status"]="logOut"
	http.Redirect(w,r,"/index/",http.StatusSeeOther)
}

func DefaultRedirect(w http.ResponseWriter, r * http.Request){
	http.Redirect(w,r,"/index/",http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/",DefaultRedirect)
	http.HandleFunc("/index/", indexHandler)
	http.HandleFunc("/student/", studentHandler)
	http.HandleFunc("/professor/",professorHandler)
	http.HandleFunc("/submit/",submitHandler)
	http.HandleFunc("/login/",login)
	http.HandleFunc("/register/",register)
	http.HandleFunc("/submit_Eval/",finish_eval)
	http.HandleFunc("/get_data/",get_data)
	http.HandleFunc("/logout/",logout)
	http.ListenAndServe(":8000", nil) 
}