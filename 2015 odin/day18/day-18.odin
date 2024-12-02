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
    lights := [100][100]int{}
    for row, y in arry {
        if row == "" { break }
        for s, x in row {
            if s == '#' { lights[y][x] = 1 }
            else { lights[y][x] = 0 }
        }
    }

    count_neighb :: proc(x: int, y: int, arr: ^[100][100]int) -> int {
        neighb := 0
        if x > 0 && arr[y][x - 1] == 1 { neighb += 1}  // CL
        if x < 99 && arr[y][x + 1] == 1 { neighb += 1} // CR
        if y > 0 {
            if x > 0 && arr[y - 1][x - 1] == 1 { neighb += 1}  // TL
            if x < 99 && arr[y - 1][x + 1] == 1 { neighb += 1} // TR
            if arr[y - 1][x] == 1 { neighb += 1 } // TC
        } 
        if y < 99 {
            if x > 0 && arr[y + 1][x - 1] == 1 { neighb += 1}  // BL
            if x < 99 && arr[y + 1][x + 1] == 1 { neighb += 1} // BR
            if arr[y + 1][x] == 1 { neighb += 1 } // BC
        } 

        return neighb
    }

    turn_on_corners :: proc(arr: ^[100][100]int) {
        arr[0][0] = 1
        arr[99][0] = 1
        arr[0][99] = 1
        arr[99][99] = 1
    }

    next_gen :: proc(arr: ^[100][100]int) {
        new_arr := [100][100]int{}
        for y in 0..<100 {
            for x in 0..<100 {
                cnt := count_neighb(x, y, arr)
                if arr[y][x] == 1 {
                    if cnt != 2 && cnt != 3 { new_arr[y][x] = 0 }
                    else { new_arr[y][x] = 1 }                    
                }
                if arr[y][x] == 0 {
                    if cnt == 3 { new_arr[y][x] = 1 }
                    else { new_arr[y][x] = 0 }                    
                }
                turn_on_corners(&new_arr)
            }
        }
        for y in 0..<100 {
            for x in 0..<100 {
                arr[y][x] = new_arr[y][x]
            }
        }
    }

    turn_on_corners(&lights) // [2]

    for i in 0..<100 {
        next_gen(&lights)
    }

    count := 0
    for y in 0..<100 {
        for x in 0..<100 {
            if lights[y][x] == 1 { count += 1 }
        }
    }

    println("Count:", count)   // [1] 768  [2] 781
}