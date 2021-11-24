/*
Go: JSON Encoding

Create a JSON-encoding system for objects.

Write a struct called Manager with fields FullName and Position (both string), 
then Age and YearsinCompany (both int32) - in that exact order. 

Also write a function that encodes objects of this struct to an encoded JSON object. 
It returns an io.Reader with this JSON object. 
Fields in the JSON object should be renamed as full_name,
position, age and years_in_company. Empty fields should be omitted.

*/

package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)


/*
 * Complete the 'EncodeManager' function below and the struct Manager.
 *
 * The function is expected to return an io.Reader and an error.
 * The function accepts *Manager manager as parameter.
 */

type Manager struct {
    FullName       string `json:"full_name,omitempty"`
    Position       string `json:"position,omitempty"`
    Age            int32 `json:"age,omitempty"`
    YearsInCompany int32 `json:"years_in_company,omitempty"`
}

func EncodeManager(manager *Manager) (io.Reader, error) {
    b, err := json.Marshal(&manager)
    r := bytes.NewReader(b)
    return r, err   
    
}
func main() {