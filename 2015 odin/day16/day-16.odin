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

    test_strings := strings.split_lines(cast(string)txt)

    target_aunt := map[string]int{ // Sue 183: trees: 9, cars: 3, goldfish: 5
        "children:" = 3,
        "cats:" = 7,
        "samoyeds:" = 2,
        "pomeranians:" = 3,
        "akitas:" = 0,
        "vizslas:" = 0,
        "goldfish:" = 5,
        "trees:" = 3,
        "cars:" = 2,
        "perfumes:" = 1
    }

    check_valid :: proc(dict: ^map[string]int, key: string, val: int) -> bool {
        if key == "cats:" || key == "trees:" {
            return dict[key] < val
        }
        if key == "pomeranians:" || key == "goldfish:" {
            return dict[key] > val
        }
        return dict[key] == val
    }

    for i in test_strings {
        if i == "" { break }
        
        ss := strings.split(i, " ")  

        curr_key := ss[2]        
        curr_val := strconv.atoi(ss[3])
        target_val := target_aunt[curr_key]
        if !check_valid(&target_aunt, curr_key, curr_val) { continue }
        
        curr_key = ss[4]        
        curr_val = strconv.atoi(ss[5])
        target_val = target_aunt[curr_key]
        if !check_valid(&target_aunt, curr_key, curr_val) { continue }
        
        curr_key = ss[6]        
        curr_val = strconv.atoi(ss[7])
        target_val = target_aunt[curr_key]
        if !check_valid(&target_aunt, curr_key, curr_val) { continue }

        println("Aunt ", ss)  // [2] 323           
        //println("Aunt ", sue_num) // [1] 213          
    }
}