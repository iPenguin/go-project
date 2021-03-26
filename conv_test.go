/***
 * This code was borrowed with modifications from https://golangdocs.com/golang-unit-testing
 ***/
package main

import "testing"

func TestConvert(t *testing.T) {
    str, err := Convert("50mi" , "km")
    if err != nil {
        t.Log("error should be nil",err)
            t.Fail()
    }
    if str != "80.467km" {
        t.Log("error should be 80.467km, but got",err)
        t.Fail()
    }
}
