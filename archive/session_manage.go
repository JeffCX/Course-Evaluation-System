package main
import (
    "net/http"
     "github.com/gorilla/sessions"
     "fmt"
    
)
//username + password hash, byte, create cookies put in database,if logout delete session 
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func MyHandler(w http.ResponseWriter, r *http.Request) {
    // Get a session. Get() always returns a session, even if empty.
    session, err := store.Get(r, "gan")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set some session values.
    session.Values["foo"] = "bar"
    session.Values[42] = 43
    fmt.Println(session.Values)
    // Save it before we write to the response/return from the handler.
    
    fmt.Println(session)
    
    session.Save(r, w)

    session = nil
    fmt.Println(session)
}

func main(){
    http.HandleFunc("/index/",MyHandler)
    http.ListenAndServe(":8000", nil) 
}