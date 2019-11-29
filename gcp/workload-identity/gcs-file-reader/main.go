package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ts, err := google.DefaultTokenSource(ctx)
	if err != nil {
		fmt.Fprintf(w, "Can't read object!,%s", err)
		return
	}
	gcs := oauth2.NewClient(ctx, ts)
	resp, err := gcs.Get(fmt.Sprintf("https://storage.cloud.google.com/%s/%s", "workload-identity", "aa"))
	if err != nil {
		fmt.Fprintf(w, "Can't read object!,%s", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
