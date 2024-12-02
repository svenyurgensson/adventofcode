package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:strconv"
import "core:runtime"

main :: proc() {
    using fmt

    txt, err := os.read_entire_file_from_filename("input.txt")
    if err == false {
        println("Ошибка открытия файла!")
        os.exit(1)
    }

    arry := strings.split_lines(cast(string)txt)
    repl := map[string]([dynamic]string){}
    target := arry[len(arry) - 2] // because of last new line
    keys := [dynamic]string{}
    comb := [dynamic]string{}

    for s in arry {
        if s == "" { break }
        ss := strings.split(s, " ")
        if ss[0] in repl {
            z := repl[ss[0]]
            append(&z, ss[2])
            repl[ss[0]] = z
        } else {
            z : [dynamic]string
            append(&z, ss[2])
            append(&keys, ss[0])
            repl[ss[0]] = z
        }
    }

    println(target)
    
    for x in strings.split_multi_iterate(&target, keys[:]) {
        //println(x)
    }
}