package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:strconv"

main :: proc() {
    txt, err := os.read_entire_file_from_filename("adventofcode.com_2015_day_6_input.txt")
    if err == false {
        fmt.println("Ошибка открытия файла!")
        os.exit(1)
    }
    test_strings := strings.split(cast(string)txt, "\n")

    lights := [1000 * 1000]int{}

    turn_on :: proc(x1, y1, x2, y2 : int, l : ^[1000_000]int) {
        for i in y1..=y2 {
            for j in x1..=x2 { l[i * 1000 + j] += 1 }
        }
    }
    turn_off :: proc(x1, y1, x2, y2 : int, l : ^[1000_000]int) {
        for i in y1..=y2 {
            for j in x1..=x2 { 
                if l[i * 1000 + j] > 0 { l[i * 1000 + j] -= 1 }
            }
        }
    }
    toggle :: proc(x1, y1, x2, y2 : int, l : ^[1000_000]int) {
        for i in y1..=y2 {
            for j in x1..=x2 { 
                l[i * 1000 + j] += 2 
            }
        }
    }
    for s in test_strings {
        ss, st, fn, op := strings.split(s, " "), "", "", ""
        if ss[0] == "" { break }
        if ss[0] == "toggle" {
            op = "tg"
            st, fn = ss[1], ss[3]
        } else {
            op = ss[1]
            st, fn = ss[2], ss[4]
        }
        sstr := strings.split(st, ",")
        x1, y1 := strconv.atoi(sstr[0]), strconv.atoi(sstr[1])
        fstr := strings.split(fn, ",")
        x2, y2 := strconv.atoi(fstr[0]), strconv.atoi(fstr[1])
        switch op {
            case "on":
                turn_on(x1, y1, x2, y2, &lights)
            case "off":
                turn_off(x1, y1, x2, y2, &lights)
            case "tg":
                toggle(x1, y1, x2, y2, &lights)
        }
        //fmt.println(x1, y1, x2, y2, op)
    }

    count := 0
    for i in 0..<1000_000 {
        if lights[i] >= 1 { count += lights[i] }
    }

    fmt.println(count) // [1] 377891 // [2] 14110788
}