/*
Генерируем пермутации по городам, проходим по всем и расчитываем расстояние
пермутации по алгоритму Хипа http://ruslanledesma.com/2016/06/17/why-does-heap-work.html
 */

package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:strconv"
import "core:runtime"
import "core:text/match"

main :: proc() {
    using fmt

    isdigit :: proc(s: string) -> bool {
        matcher := match.matcher_init(s, "^%d+$")
	    _, ok := match.matcher_match(&matcher)
        return ok
    }

    txt, err := os.read_entire_file_from_filename("input.txt")
    if err == false {
        println("Ошибка открытия файла!")
        os.exit(1)
    }

    test_strings := strings.split_lines(cast(string)txt)
    count := len(test_strings)

    cities := map[string]int{}
    defer delete(cities)
    routes := map[string]int{} // [Bel_Dub=141, Dub_Bel=141, Bel_Lon=518, Lon_Dub=464, Lon_Bel=518, Dub_Lon=464]
    defer delete(routes)

    idx := 0
    for s in test_strings {
        ss := strings.split(s, " ")
        dist := strconv.atoi(ss[4])
        key := fmt.aprintf("%s_%s", ss[0][0:3], ss[2][0:3])
        routes[key] = dist
        key = fmt.aprintf("%s_%s", ss[2][0:3], ss[0][0:3])
        routes[key] = dist
        cities[ss[0][0:3]] = 1
        cities[ss[2][0:3]] = 1
    }
    cities_count := len(cities)
    cities_arr := make([dynamic]string) // ["Dub", "Lon", "Bel"]
    cities_idx := make([dynamic]int)    // [0, 1, 2]
    idx = 0
    for s in cities {
        append_elem(&cities_arr, s)
        append_elem(&cities_idx, idx)
        idx += 1
    }

    calc_total_dist :: proc(cities_indexes : ^[dynamic]int, cities : ^[dynamic]string, routes : ^map[string]int) -> int {
        ln, dist := len(cities_indexes), 0
        for idx in 0..<(ln - 1) {
            key := aprintf("%s_%s", cities[cities_indexes[idx]], cities[cities_indexes[idx + 1]])
            //println(key)
            dist += routes[key]
        }
        return dist
    }

    permute_check_minimum :: proc(ary: ^[dynamic]int, n: int, distance: ^int, cities: ^[dynamic]string, routes : ^map[string]int) {
        i := 0
        for {
            if n > 2 do permute_check_minimum(ary, n - 1, distance, cities, routes)
            if i == (n - 1) do return 
            if n % 2 == 0 {
                t1, t2 := ary[i], ary[n - 1]
                ary[i] = t2
                ary[n - 1] = t1
            } else {
                t1, t2 := ary[0], ary[n - 1]
                ary[0] = t2
                ary[n - 1] = t1
            }            
            dist := calc_total_dist(ary, cities, routes)
            //if dist < distance^ do distance^ = dist // [1]
            if dist > distance^ do distance^ = dist // [2]
            i += 1           
        }
    }

    result_dist := calc_total_dist(&cities_idx, &cities_arr, &routes)
    println("first dist:", result_dist)
    permute_check_minimum(&cities_idx, len(cities_idx), &result_dist, &cities_arr, &routes)
    println("result dist:", result_dist)
}