package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:crypto/md5"

main :: proc() {
    prefix := "ckczppom" // [1] 117946 // [2] 3938038
    cicle := 1

    for {
        current := fmt.aprintf("%s%d", prefix, cicle)
        b := md5.hash(current)
        if b[0] == 0 && b[1] == 0 && b[2] == 0 {
            fmt.println(cicle)
            os.exit(0)            
        }   
        cicle += 1
        if cicle > 10_000_000 { 
            fmt.println("Not found :(")
            os.exit(1) 
        }
    }
}