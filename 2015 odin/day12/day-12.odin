package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:strconv"
import "core:runtime"
import "core:text/match"
import "core:encoding/json"
import "core:reflect"

main :: proc() {
    using fmt

    bindata, err := os.read_entire_file_from_filename("input.txt")
    if err == false {
        println("Ошибка открытия файла!")
        os.exit(1)
    }

    text := cast(string)bindata
    sum := 0
 
    matcher := match.matcher_init(text, "-?%d+" , 0)
    for {
        str, idx, ok := match.matcher_match_iter(&matcher)
        if !ok do break 
        sum += strconv.atoi(str)
    }
    println("sum:", sum)

    walk_nodes :: proc(node: json.Value, sum: int) -> int {
        n_sum := sum
     
        object, is_object := node.(json.Object) 
        if is_object {
            for k in object {
                #partial switch _ in object[k] {
                    case json.String:
                        if object[k].(json.String) == "red" { return 0 }
                    case:
                }
            }
        }

        #partial switch _ in node {
            case json.Object:
                for j in node.(json.Object) {
                    n_sum += walk_nodes(node.(json.Object)[j], 0)
                }              
            case json.Float:
                return cast(int)node.(json.Float)
            case json.Array:
                for j in node.(json.Array) {
                    n_sum += walk_nodes(j, 0)
                }                   
            case:
        }

        return n_sum
    }

    println()
    test := "{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}"
    zbindata := transmute([]u8)test
    
    data, err2 := json.parse(bindata) 
    sum = 0
    res := walk_nodes(data, sum)
    println("---")
    println("Sum without 'red':", res)
}