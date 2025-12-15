//8080 portunu dinlyeip gelen http response ile mesaj yazdıran go uygulaması 

package main

import (
	"fmt"
	"net/http"
	"os"
)
//repsonse gleince mesajı verecek fonksiyon
func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Merhaba! Ben single-stage build demo uygulamasıyım.\nHostname: %s\n", hostname)
}


//ana fonksiyon
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server 8080 portunda dinleniyor...")
	http.ListenAndServe(":8080", nil)
}
