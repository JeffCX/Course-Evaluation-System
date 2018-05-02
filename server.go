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
var  course_infos []string 
var course_infos_register []string

func check_session( users string ,password string ) bool {
	
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



type classes_struct struct {
    M map[string]string
    N map[string]string
   
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

func studentHandler(w http. ResponseWriter, r *http.Request) {
	if(check_session(Username,Password)) {
		m:=map[string]string{}
		for i:=0;i<len(course_infos);i++ {
			if i%2==0 {
				m[course_infos[i]]=course_infos[i+1]
			}
		}

		n:=map[string]string{}
		for i:=0;i<len(course_infos_register);i++ {
			if i%2==0 {
				n[course_infos_register[i]]=course_infos_register[i+1]
			}
		}


		haha:=classes_struct{M:m,N:n}
	    t, _ := template.ParseFiles("cart.html")
	    t.Execute(w, haha)
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
		db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
		//db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
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
		    session["username"] = Username

		    course_display,_:=db.Query("SELECT Course_Name,Professor_netID FROM Course;")
		    var name_list []string
		    var name_list_item string
		    var name_list_p string 
		    for course_display.Next(){
		    	course_display.Scan(&name_list_item,&name_list_p)
		    	name_list=append(name_list,name_list_item,name_list_p)
		    }
		    fmt.Println(name_list)
			course_infos = name_list

			course_display_selected,_:=db.Query("SELECT Course_Name,filled from StudentCourse;")
			var name_list_selected []string
			var name_list_item_selected string
			var filled string 
			for course_display_selected.Next(){
				course_display_selected.Scan(&name_list_item_selected,&filled)
		    	name_list_selected=append(name_list_selected,name_list_item_selected,filled)
			}
			course_infos_register = name_list_selected
			fmt.Println(course_infos_register)
			http.Redirect(e,r,"/student/",http.StatusSeeOther)
			control=false


				
		}
		db.Close()

		if (control){
		fmt.Println("gan")
		dbs,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
		//dbs, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
		row,_:=dbs.Query("SELECT Professor_netID, password FROM Professor where Professor_netID=?;",r.Form["username"][0])
		var Pusername string 
		var Ppassword string 
		fmt.Println(Pusername,Ppassword)
		for row.Next(){
			row.Scan(&Pusername,&Ppassword)
		}
		

		if CheckPasswordHash(r.Form["password"][0],Ppassword){
			Username = r.Form["username"][0]
		    Password = r.Form["password"][0]
		     session[Username] = Username
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
		db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
		//db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
		rows,_:=db.Query("SELECT Student_netID FROM Student where Student_netID=?",r.Form["username_reg"][0])
		row,_:=db.Query("SELECT Professor_netID FROM Professor where Professor_netID=?",r.Form["username_reg"][0])
	
		
		if rows.Next() == false && row.Next() == false {
			for rows.Next() {
          		rows.Scan(&Student_netID)
        	} 
        	pass,_:=HashPassword(r.Form["password_reg"][0]) 
        	if r.Form["Permission_code"][0]=="" {
        			stmt,_:=db.Prepare("INSERT into Student VALUES(?,?,?)")
        			stmt.Exec(r.Form["username_reg"][0],pass,r.Form["email_reg"][0])
				    Username = r.Form["username_reg"][0]
				    Password = r.Form["password_reg"][0]
				    session[Username+Password] =Username+Password
				    session["username"] = Username
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
		db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
		//db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
	
		Course_Rate:=convert_to_string(r.Form["Course_Rate"])
		
	
		Student_Comment_Advice:=convert_to_string(r.Form["Student_Comment_Advice"])
		
		
		Student_Comment_Improve:=convert_to_string(r.Form["Student_Comment_Improve"])
		Instructor_Rate:=convert_to_string(r.Form["Instructor_Rate"])
		Instructor_Inspiration :=convert_to_string(r.Form["Instructor_Inspiration"])
		Course_Comment:=convert_to_string(r.Form["Course_Comment"])

		
		Instructor_Clarity :=convert_to_string(r.Form["Instructor_Clarity"])
		Instructor_Feedback:=convert_to_string(r.Form["Instructor_Feedback"])
		Course_Content:=convert_to_string(r.Form["Course_Content"])
		Course_Application:=convert_to_string(r.Form["Course_Application"])
		Instructor_Comment:=convert_to_string(r.Form["Instructor_Comment"])
	

		Course_Organize:=convert_to_string(r.Form["Course_Organize"])
		Instructor_Goal :=convert_to_string(r.Form["Instructor_Goal"])
		Student_netID :="xc1008"
		Course_Name :="CS1004"
		fmt.Println()
		stmt,_:=db.Prepare("INSERT into Course_Response VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
		stmt.Exec(Student_netID,Course_Rate ,Student_Comment_Advice ,Student_Comment_Improve ,Instructor_Rate ,Instructor_Goal ,Instructor_Inspiration ,Course_Comment ,Instructor_Clarity ,Instructor_Feedback ,Course_Content,Course_Organize ,Course_Application ,Instructor_Comment ,Course_Name )

		db.Close()	
		http.Redirect(w,r,"/submit/",http.StatusSeeOther)
	
	}

	// insert into dabase
	//redirect to submit page
}
//error handling 


func get_data(w http.ResponseWriter,r *http.Request){
	db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
	//db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")

	stmt,_:=db.Query("select * from Course_Response;")
	var Student_netID string
	var Course_Rate string 
	var Student_Comment_Advice string 
	var Student_Comment_Improve string 
	var Instructor_Rate string 
	var Instructor_Inspiration string 
	var Course_Comment string 
	var Instructor_Clarity string 
	var Instructor_Feedback string 
	var Course_Content string 
	var Course_Application string 
	var Instructor_Comment string 
	var Course_Organize string 
	var Instructor_Goal string 
	var Course_Name string 
	for stmt.Next(){
		stmt.Scan(&Student_netID,&Course_Rate ,&Student_Comment_Advice ,&Student_Comment_Improve ,&Instructor_Rate ,&Instructor_Goal ,&Instructor_Inspiration ,&Course_Comment ,&Instructor_Clarity ,&Instructor_Feedback ,&Course_Content,&Course_Organize ,&Course_Application ,&Instructor_Comment ,&Course_Name )
	}
	data:=[]string {Student_netID,",",Course_Rate ,",",Student_Comment_Advice ,",",Student_Comment_Improve ,",",Instructor_Rate ,",",Instructor_Goal ,",",Instructor_Inspiration ,",",Course_Comment ,",",Instructor_Clarity ,",",Instructor_Feedback ,",",Course_Content,",",Course_Organize ,",",Course_Application ,",",Instructor_Comment ,",",Course_Name }
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

func createClass(w http.ResponseWriter, r * http.Request){
	if r.Method == "POST" {
		r.ParseForm()
		db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
		//db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
		CourseName:=(r.Form["CourseName"])[0]
		CourseID:=convert_to_string(r.Form["CourseID"])
		semseter:=convert_to_string(r.Form["semseter"])
		p_id :=  convert_to_string(r.Form["InstructorID"])
		fmt.Println(p_id)
		fmt.Println(CourseName,CourseID,semseter,p_id)
 		stmt,_:=db.Prepare("INSERT into Course VALUES (?,?,?,?);")
		stmt.Exec(CourseID,p_id,CourseName,semseter)
		db.Close()	
		http.Redirect(w,r,"/professor/",http.StatusSeeOther)
	}else{http.Redirect(w,r,"/index/",http.StatusSeeOther)}
}



func StudentRegister(w http.ResponseWriter, r * http.Request){
	if r.Method == "POST" {
		r.ParseForm()
		//db, _:= sql.Open("mysql", "sql9228084:WIKHkznFfd@tcp(sql9.freemysqlhosting.net:3306)/sql9228084")
		fmt.Println("Student register")
		CourseName:=(r.Form["class_list"])
		for i:=0;i<len(CourseName);i++ {
			db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
			stmt,_:=db.Prepare("INSERT into StudentCourse VALUES (?,?,?);")
			id:=session["username"]
			stmt.Exec(CourseName[i],id,"0")
			db.Close()			
		}	
		http.Redirect(w,r,"/student/",http.StatusSeeOther)	
	}else{http.Redirect(w,r,"/index/",http.StatusSeeOther)}
}

func EvaluteHandler(w http.ResponseWriter,r *http.Request) {
	if check_session(Username,Password){
	db,_:=sql.Open("mysql","root:heizhenzhu@/Course_Evaluation")
	rows,_:=db.Query("SELECT Course_Name,filled FROM StudentCourse where Student_netID=?",session["username"])
	var name_list_selected [] string
	var name_list_item_selected string
	var filled string 
	for rows.Next(){
		rows.Scan(&name_list_item_selected,&filled)
    	name_list_selected =append(name_list_selected,name_list_item_selected,filled)
	}
	
	m:=map[string]string{}
	for i:=0;i<len(course_infos);i++ {
		if i%2==0 {
			m[name_list_selected[i]]=name_list_selected[i+1]
		}
	}



	
    t, _ := template.ParseFiles("student.html")
    t.Execute(w, m)
	}else {
    		session["status"]="LoginFirst"
		http.Redirect(w,r,"/index/",http.StatusSeeOther)
    }
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
	http.HandleFunc("/create/",createClass)
	http.HandleFunc("/register_class/",StudentRegister)
	http.HandleFunc("/evaluate/",EvaluteHandler)
	http.ListenAndServe(":8000", nil) 
}