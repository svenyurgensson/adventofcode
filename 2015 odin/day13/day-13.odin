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

    push_if_not_exists :: proc(elm: string, ary: ^[dynamic]string) {
        for i in ary {
            if i == elm { return }
        }
        append(ary, elm)
    }

    test_strings := strings.split_lines(cast(string)txt)
    uniq_names := [dynamic]string{}
    participants := map[string]int{}

    for i in test_strings {
        if i == "" { break }        
        ss := strings.split(i, " ") 
        who := ss[0]
        whom := strings.cut(ss[10], 0, len(ss[10]) - 1)
        amount := strconv.atoi(ss[3])
        if ss[2] == "lose" { amount *= -1 }
        key := fmt.aprintf("%s_%s", who, whom)
        participants[key] = amount
        push_if_not_exists(who, &uniq_names)
    }
    append(&uniq_names, "Me") // [2]
    for i in uniq_names {
        if i == "Me" do continue 
        key := fmt.aprintf("%s_%s", i, "Me")
        participants[key] = 0
        key = fmt.aprintf("%s_%s", "Me", i)
        participants[key] = 0
    }
    println(uniq_names)
    println(participants)

    calc_happiness :: proc(positions: ^[dynamic]string, places: ^map[string]int) -> int {
        sum := 0
        for person, idx in positions {
            left := idx - 1
            if left < 0 { left = len(positions) - 1 }
            right := idx + 1
            if right >= len(positions) { right = 0 }
            key_left := fmt.aprintf("%s_%s", person, positions[left])
            key_right := fmt.aprintf("%s_%s", person, positions[right])
            sum += places[key_left] + places[key_right]
        }
        return sum
    }

    permute_find_maximum :: proc(positions: ^[dynamic]string, n: int, happiness: ^int, places: ^map[string]int) {
        i := 0
        for {
            if n > 2 do permute_find_maximum(positions, n - 1, happiness, places)
            if i == (n - 1) do return 
            if n % 2 == 0 {
                t1, t2 := positions[i], positions[n - 1]
                positions[i] = t2
                positions[n - 1] = t1
            } else {
                t1, t2 := positions[0], positions[n - 1]
                positions[0] = t2
                positions[n - 1] = t1
            }            
            happy := calc_happiness(positions, places)
            if happy > happiness^ do happiness^ = happy
            i += 1           
        }
    }

    max_happiness := 0
    permute_find_maximum(&uniq_names, len(uniq_names), &max_happiness, &participants)
    println(max_happiness) // [1] 618 [2] 601  
}