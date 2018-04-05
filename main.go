package main

import (
	"net/http"
	"fmt"
	"os"
	"time"
)


//ssl config
var SSL_config = map[string]string {
	"enable_HTTPS" : "NO", // <= Type YES, when you want to enable HTTPS
	"SSL_Cert" : "", //Cert file path
	"SSL_Key" : "", //Key file path
}



var http_server_port string = ":80"


func main(){
	if(!has_dir("logs")){
		os.Mkdir("logs",0644)
		fmt.Println("Created logs dir")
	}

	fmt.Println("HTTP Server running...")

	http.HandleFunc("/", http_server)

	var http_listen interface{}


	if(SSL_config["enable_HTTPS"] == "YES"){
		http_listen = http.ListenAndServeTLS(http_server_port,SSL_config["SSL_Cert"],SSL_config["SSL_Key"], nil)
	}else{
		http_listen = http.ListenAndServe(http_server_port ,nil)
	}

	if(http_listen != nil){
		fmt.Println(http_listen)
	}
}

func http_server(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	request_time := get_time("full")
	request_log := request_time + "  " +r.Method + "Request from > " + r.RemoteAddr + " Request to > " + r.URL.Scheme + r.Host + r.URL.Path

	insert_log ,_:=os.OpenFile("logs/request_log.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	buf := []byte(request_log + "\n")
	insert_log.Write(buf)
	insert_log.Close()
	fmt.Println(insert_log)


	fmt.Fprintf(w,"Hello,World")
}

//this function will check dir
func has_dir(Dirname string)(bool){
	finfo, err := os.Stat(Dirname)
	if(err != nil){
		return false
	}else {
		if (finfo.IsDir()) {
			return true
		} else {
			return false
		}
	}
}

func get_time(how string)(NowTime string){
	/* how :
	full format : Y-m-d H:m:s
	date : y-m-d
	 */
	if(how == "full"){
		return time.Now().Format("2006-01-02 15:04:05")
	}else{
		return time.Now().Format("2006-01-02")
	}

}
