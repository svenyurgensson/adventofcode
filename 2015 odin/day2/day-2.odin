package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:slice"
import "core:mem"
import "core:strconv"
import "core:math"
import "core:sort"

main :: proc() {
    txt, err := os.read_entire_file_from_filename("adventofcode.com_2015_day_2_input.txt")
    if err == false {
        fmt.println("Ошибка открытия файла!")
        os.exit(1)
    }
    dimensions := strings.split(cast(string)txt, "\n")
    total, total_ribbon := 0, 0

    for d in dimensions {
        ss := strings.split(d, "x")
        if len(ss[0]) == 0 || len(ss[1]) == 0 || len(ss[2]) == 0  do break
        dl, dw, dh := strconv.atoi(ss[0]), strconv.atoi(ss[1]), strconv.atoi(ss[2])
        
        lw, wh, hl := dl * dw, dw * dh, dh * dl
        surf_area := 2 * lw + 2 * wh + 2 * hl
        smallest := math.min(math.min(lw, wh), hl)
        total += surf_area + smallest

        // second
        ary := []int{dl, dw, dh}
        sort.quick_sort(ary)
        total_ribbon += 2 * ary[0] + 2 * ary[1] + dl * dw * dh 
    }

    fmt.println(total) // first: 1588178
    fmt.println(total_ribbon) // second: 3783758
}