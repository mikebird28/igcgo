package igcgo

import(
    "testing"
    "os"
    _ "fmt"
)

func TestIGC(t *testing.T){
    path := "./20141223_2888_23.igc"
    fp,err := os.Open(path)
    if err != nil{
    }else{
        //igc,_ := ParseIGC(fp)
        ParseIGC(fp)
        //fmt.Println(igc.Json())
    }

}


