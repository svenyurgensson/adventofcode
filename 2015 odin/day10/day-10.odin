package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:strconv"
import "core:runtime"
import "core:text/match"
import "core:slice"

main :: proc() {
    using fmt

    initial := []int{3, 1, 1, 3, 3, 2, 2, 1, 1, 3}
    //initial := []int{1,1,1,2,2,1}

    build_look_and_say :: proc(ary: ^[]int) -> []int {
        tmp_arry := [dynamic]int{}
        size := len(ary)
        pos := size - 1
        curr_arr_pos := 0
        for {
            curr_elm := ary[pos]
            inject_at(&tmp_arry, curr_arr_pos, curr_elm)
            cnt := 1
            for {
                pos -= 1
                if pos < 0 {
                    curr_arr_pos += 1
                    inject_at(&tmp_arry, curr_arr_pos, cnt)
                    break 
                }
                if ary[pos] == curr_elm { 
                    cnt += 1 
                } else {
                    curr_arr_pos += 1
                    inject_at(&tmp_arry, curr_arr_pos, cnt)
                    curr_arr_pos += 1
                    break
                }
            } 
            if pos < 0 { break }
        }

        slice.reverse(tmp_arry[:])
        return tmp_arry[:]
    }    

    for i in 0..<50 {
        initial = build_look_and_say(&initial)
    }

    //println(initial)
    println(len(initial))
}