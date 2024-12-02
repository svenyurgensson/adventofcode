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

    Ingr :: struct {
        name: string,
        capacity: int, 
        durability: int, 
        flavor: int, 
        texture: int, 
        calories: int
    }

    ingridients := [dynamic]Ingr{}

    for i in test_strings {
        if i == "" { break }
        ss := strings.split(i, " ") 
        append(&ingridients, Ingr{
            name = ss[0], capacity = strconv.atoi(ss[2]), durability = strconv.atoi(ss[4]), 
            flavor = strconv.atoi(ss[6]), texture = strconv.atoi(ss[8]), calories = strconv.atoi(ss[10])
        })
    }

    calc_score :: proc(ingr: ^[dynamic]Ingr, tsp_ary: [4]int) -> (int, int) {
        total_capacity := 0 
        total_durability := 0
        total_flavor := 0 
        total_texture := 0
        total_cals := 0
        for d, idx in ingr {
            total_capacity += d.capacity * tsp_ary[idx]
            total_durability +=  d.durability * tsp_ary[idx]
            total_flavor +=  d.flavor * tsp_ary[idx]
            total_texture +=  d.texture * tsp_ary[idx]
            total_cals +=  d.calories * tsp_ary[idx]
        }
        if total_capacity < 0 || total_durability < 0 || total_flavor < 0 || total_texture < 0 { return 0, 0 }
        return total_capacity * total_durability * total_flavor * total_texture, total_cals
    }

    best := 0

    // cheating here, as I know beforehead that here is only 4 products
    for i in 1..=100 {
        for j in 1..=(100-i) {
            for k in 1..=(100-i-j) {
                h := 100 - i - j - k
                if h < 1 { continue }
                teaspoons := [4]int{i, j, k, h}
                //score := calc_score(&ingridients, teaspoons) // [1]
                score, cals := calc_score(&ingridients, teaspoons) // [2]
                if cals == 500 && best < score { best = score }
            }
        }
    }
    println("Best score:", best) // [1] 222870 [2] 117936

}