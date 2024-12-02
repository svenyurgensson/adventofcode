package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:strconv"
import "core:text/match"

main :: proc() {
    txt, err := os.read_entire_file_from_filename("adventofcode.com_2015_day_8_input.txt") // ("tests.txt") //
    if err == false {
        fmt.println("Ошибка открытия файла!")
        os.exit(1)
    }

    count_total := 0
    count_in_mem := 0
    total_len := len(txt)
    i := 0
    for {
        ch := txt[i]
        count_total += 1
        i += 1
        switch ch {
            case '\n': {
                count_total -= 1 // correct \n
                count_in_mem -= 2 // correct " and "
            }
            case '\\': { // \
                count_in_mem += 1               
                if txt[i] == 'x' { // \xFF
                    i += 3
                    count_total += 3
                } else { // \" \\ 
                    i += 1 
                    count_total += 1
                }
            }
            case: { // all others
                count_in_mem += 1
            }
        }
        if i >= total_len { break }
    }

    fmt.println("--- first part ---")
    fmt.println("file length:", total_len)
    fmt.println("char in code:", count_total)
    fmt.println("char in mem:", count_in_mem)
    fmt.println("diff:", count_total - count_in_mem)
    //
    fmt.println("\n--- second part ---")

    count_total = 0
    count_in_mem = 0
    i = 0
    
    for {
        ch := txt[i]
        count_total += 1
        i += 1
        switch ch {
            case '\n': {
                count_total -= 1 // correct \n
                count_in_mem += 4 // correct " and "
            }
            case '\\': { // \                              
                if txt[i] == 'x' { // \xFF
                    i += 3
                    count_in_mem += 5 
                    count_total += 3
                } else { // \" \\ 
                    i += 1 
                    count_in_mem += 4
                    count_total += 1
                }
            }
            case: { // all others
                count_in_mem += 1
            }
        }
        if i >= total_len { break }
    }

    fmt.println("char in code:", count_total)
    fmt.println("char in mem:", count_in_mem)
    fmt.println("diff:", count_in_mem - count_total)
}

/*

""
"abc"
"aaa\"aaa"
"\x27"

 */