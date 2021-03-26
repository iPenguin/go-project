/***
 * This code was borrowed with modifications from https://golangdocs.com/golang-unit-testing
 ***/
package main

import (
    "fmt"
    "strconv"
    "strings"
)

const (
    kilometersToMiles = 0.621371
    milesToKilometres = 1.60934
    )

func Convert (from, to string) (string,error) {
    var result float64
    switch {
    case strings.HasSuffix(from,"mi"):
        miles,err := strconv.ParseFloat(from[:len(from)-2],64)
        if err != nil {
            return "",err
            }

        switch to {
            case "km":
                result = miles * milesToKilometres
            case "m":
                result = miles * milesToKilometres * 1000
            case "mi":
                result = miles
        }
    }
    return strconv.FormatFloat(result, 'f', -1, 64) + "km", nil
}

func main () {

    data, _ := Convert( "50mi", "km" )
    fmt.Println("Converts to: ", data )

}
