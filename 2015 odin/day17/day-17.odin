package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:slice"

containers := []int{
    50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 40
}

main :: proc() {
    using fmt
    slice.reverse_sort(containers)
    target :: 150
    combinations := [dynamic]int{}

    calc :: proc(remaining: int, conts: ^[]int, used: int, combin: ^[dynamic]int) {
        if remaining == 0 {
            append(combin, used)
        }            
        else if remaining > 0 && len(conts) > 0 {
            idx := len(conts)
            rest := conts[1 : idx]
            calc(remaining - conts[0], &rest, used+1, combin)
            calc(remaining, &rest, used, combin)
        }
    }  
    calc(target, &containers, 0, &combinations)

    m := slice.min(combinations[:])
    second := 0

    for c in combinations {
        if c == m do second += 1
    }

    println(len(combinations))
    println(second)
}

// 654
// 57