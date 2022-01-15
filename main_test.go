package main

import (
	"fmt"
	"strings"
	"math"
	"testing"
	"io/ioutil"
    	"github.com/CapralDavid/Architecture4/engine"
)




func UpdateFileToDefault() {
    ioutil.WriteFile("./testFile.txt", []byte("reverse kurwa"), 0644)
}


func ConcatStringNtimes(k int) {
    for i := 0; i < k; i++ {
        input,_  := ioutil.ReadFile("./testFile.txt")
        parts := strings.Fields(string(input))

        replacedString := fmt.Sprintf("%s%s", parts[1], parts[1])
        newCmd := fmt.Sprintf("%s %s", parts[0], replacedString)
        ioutil.WriteFile("./testFile.txt", []byte(newCmd), 0644)
    }
}




func BenchmarkParser(b *testing.B) {

	for i := 0; i < 12; i++ {
		UpdateFileToDefault()
	    ConcatStringNtimes(i)
		cmd, _ := ioutil.ReadFile("./testFile.txt")

		b.Run(fmt.Sprintf("len=%d", int(math.Pow(2.0, float64(i)))), func(b *testing.B) {
		    for i := 0; i < b.N; i++ {
		        engine.Parse(string(cmd))
		    }
		})

        UpdateFileToDefault()
	}
    

}
