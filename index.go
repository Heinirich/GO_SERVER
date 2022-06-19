package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"html"
)


type User struct{
	
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(html.EscapeString(r.URL.Path));
	if(r.URL.Path == "/"){
		if r.Method == "GET"{
			t,_ := template.ParseFiles("files/index.hgf")
			t.Execute(w,nil)
		}else{
			r.ParseForm()
			fmt.Println(r.Form)
			fmt.Println("Path ",r.URL.Path)
			fmt.Println("scheme",r.URL.Scheme)
			fmt.Println(r.Form["url_long"])
			if len(r.Form["name"]) == 0 {
				fmt.Fprint(w,"All Fields Are Needed")
			}
			for k,v := range r.Form{
				fmt.Println("Key ",k)
				fmt.Println("Value",v)
			}
			fmt.Fprintf(w, "Hello astaxie!") // send data to client side
		}
	}
}

func contactUs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t,_ := template.ParseFiles("files/contact.hgf")
		t.Execute(w,nil)
	}
	if r.Method == "POST"{
		r.ParseForm()
		for k,v := range r.Form{
			fmt.Printf("KEY %v",k)
			fmt.Printf("VALUE %v",v)
		}
	}
}


func main(){
	http.HandleFunc("/",welcomePage)
	http.HandleFunc("/contact",contactUs)

	
	err := http.ListenAndServe(":9000", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}