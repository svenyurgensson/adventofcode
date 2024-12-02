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
    
    ReinDeer :: struct {
        name: string,
        speed, fly_duration, rest_duration: int,
        score: int,
        dist: int
    }

    deers := [dynamic]ReinDeer{}

    for i in test_strings {
        if i == "" { break }        
        ss := strings.split(i, " ") 
        append(&deers, 
                ReinDeer{
                    name = ss[0],
                    speed = strconv.atoi(ss[3]),
                    fly_duration = strconv.atoi(ss[6]),
                    rest_duration = strconv.atoi(ss[13]),
                    score = 0, dist = 0
                }
        )
    }

    time_stop :: 2503 // [1] 2660 [2]
    //time_stop :: 1000 // [t1] 1092 [t2] 689
    max_dist := 0

    for d in deers {
        rounded := time_stop / (d.fly_duration + d.rest_duration)
        rem := time_stop % (d.fly_duration + d.rest_duration)
        add_dist := 0
        if rem >= d.fly_duration { 
            add_dist = d.fly_duration * d.speed 
        } else { 
            add_dist = rem * d.speed 
        }
        dist := rounded * d.fly_duration * d.speed + add_dist
        //println(rounded, rem, add_dist, dist)
        if dist > max_dist { max_dist = dist }
    }
    println("Max dist [1]:", max_dist)

    calc_dist_at_time :: proc(d: ^ReinDeer, second: int) -> int {
        rounded := second / (d.fly_duration + d.rest_duration)
        rem := second % (d.fly_duration + d.rest_duration)
        add_dist := 0
        if rem > d.fly_duration { // on rest
            add_dist = d.fly_duration * d.speed 
        } else { // on fly
            add_dist = rem * d.speed 
        }
        d.dist = rounded * d.fly_duration * d.speed + add_dist
        return d.dist
    }

    find_max_dist :: proc(deers: ^[dynamic]ReinDeer) -> int {
        max_dist := 0
        for d in deers {
            if d.dist > max_dist { max_dist = d.dist }
        }
        return max_dist
    }

    add_score_to_leads :: proc(deers: ^[dynamic]ReinDeer, max_dist: int) {
        for d in deers {
            if d.dist >= max_dist { d.score += 1 }
        }
    }

    find_winner_score_points :: proc(deers: ^[dynamic]ReinDeer) -> int {
        max_score := 0
        for d in deers {
            if d.score >= max_score { max_score = d.score }
        }
        return max_score
    }

    for curr_sec in 1..=time_stop {
        for _, idx in deers {
            calc_dist_at_time(&deers[idx], curr_sec)
        }
        max_dist = find_max_dist(&deers)
        add_score_to_leads(&deers, max_dist)
    }
    max_score := find_winner_score_points(&deers)
    println("Winning score[2]:", max_score)
}