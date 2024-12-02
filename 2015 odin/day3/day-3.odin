package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:math"

main :: proc() {
    txt, err := os.read_entire_file_from_filename("adventofcode.com_2015_day_3_input.txt")
    if err == false {
        fmt.println("Ошибка открытия файла!")
        os.exit(1)
    }
    x, y := 0, 0
    xr, yr := 0, 0
    visited := map[string]int { "0_0" = 1 }
    defer delete(visited)

    key : string
    for ch, idx in txt {
        even := idx % 2 == 0
        if even {
            switch ch {
                case '^':
                    y += 1
                case 'v':
                    y -= 1
                case '<':
                    x -= 1
                case '>':
                    x += 1
                case:
                    continue 
            }
            key = fmt.aprintf("%d_%d", x, y)
        } else {
            switch ch {
                case '^':
                    yr += 1
                case 'v':
                    yr -= 1
                case '<':
                    xr -= 1
                case '>':
                    xr += 1
                case:
                    continue 
            }
            key = fmt.aprintf("%d_%d", xr, yr)
        }

        ok := key in visited 
        if ok {
            visited[key] += 1
        } else {
            visited[key] = 1
        }
    }
    total_visited := len(visited)
    fmt.println(total_visited)
}