package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func get_content() {
    // json data
    url := "http://localhost:8983/solr/dataset_ranks/select?q=CollegeName:a"
    // url += "&limit=1" // limit data for testing
    res, err := http.Get(url)
    if err != nil {
        panic(err.Error())
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }
    var data interface{} // TopTracks
    err = json.Unmarshal(body, &data)
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("Results: %v\n", data)
    os.Exit(0)
}

func main() {
    get_content()
}