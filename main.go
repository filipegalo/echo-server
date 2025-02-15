package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
)

func handler(w http.ResponseWriter, r *http.Request) {
	podName := os.Getenv("HOSTNAME")
	nodeName := os.Getenv("NODE_NAME")
	podIP := os.Getenv("POD_IP")

	if podName != "" && nodeName != "" && podIP != "" {
		fmt.Fprintf(w, "Request served by %s (%s) on node %s\n\n", podName, podIP, nodeName)
	} else if podName != "" && podIP != "" {
		fmt.Fprintf(w, "Request served by %s (%s)\n\n", podName, podIP)
	} else if podName != "" && nodeName != "" {
		fmt.Fprintf(w, "Request served by %s on node %s\n\n", podName, nodeName)
	} else if podName != "" {
		fmt.Fprintf(w, "Request served by %s\n\n", podName)
	} else if nodeName != "" {
		fmt.Fprintf(w, "Request served on node %s\n\n", nodeName)
	}
	fmt.Fprintf(w, "HTTP/1.1 %s %s\n", r.Method, r.URL.Path)

	var headerKeys []string
	for name := range r.Header {
		headerKeys = append(headerKeys, name)
	}
	sort.Strings(headerKeys)

	for _, name := range headerKeys {
		for _, value := range r.Header[name] {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	fmt.Printf("Version: %s, Build Time: %s, Commit: %s\n", Version, BuildTime, CommitSHA)
	http.ListenAndServe(":8080", nil)
}
