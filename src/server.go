package main
import (
    "fmt"
    "net"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    ifaces, _ := net.Interfaces()
    for _, iface := range ifaces {
	addrs, _ := iface.Addrs()
	for _, addr := range addrs {
	    if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	        if ipnet.IP.To4() != nil {
                    fmt.Fprintf(w, "%s\n", ipnet.IP.String())
	        }
            }
        }
    }
}
