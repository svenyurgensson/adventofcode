package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:text/match"

main :: proc() {
    has_prohib :: proc(s: string) -> bool {
        matcher := match.matcher_init(s, "ab")
	    _, ok1 := match.matcher_match(&matcher)
        matcher = match.matcher_init(s, "cd")
	    _, ok2 := match.matcher_match(&matcher)     
        matcher = match.matcher_init(s, "pq")
	    _, ok3 := match.matcher_match(&matcher)
        matcher = match.matcher_init(s, "xy")
	    _, ok4 := match.matcher_match(&matcher)
        return ok1 || ok2 || ok3 || ok4
    }
    has_3_vowels :: proc(s: string) -> bool {
        matcher := match.matcher_init(s, "[aeiou]+.*[aeiou]+.*[aeiou]+")
	    _, ok := match.matcher_match(&matcher)
        return ok
    }
    has_twice_in_row :: proc(s: string) -> bool {
        for i in 0..<(len(s) - 1) {
            if s[i] == s[i+1] { return true }
        }
        return false
    }
    has_triple :: proc(s: string) -> bool {
        for i in 0..<(len(s) - 2) {
            if s[i] == s[i+2] { return true }
        }
        return false
    }
    has_good_pair :: proc(s: string) -> bool {
        l := len(s)
        i, j : int
        for i in 0..<(l - 2) {
            ii := i + 2
            ts := s[i:ii]
            for j in (i + 2)..<(l - 1) {
                jj := j + 2
                if ts == s[j:jj] { return true }
            }  
        }
        return false
    }

    txt, err := os.read_entire_file_from_filename("adventofcode.com_2015_day_5_input.txt")
    if err == false {
        fmt.println("Ошибка открытия файла!")
        os.exit(1)
    }
    test_strings := strings.split(cast(string)txt, "\n")
    count, count1 := 0, 0

    for s in test_strings {
        proh := has_prohib(s)
        vow3 := has_3_vowels(s)
        tw : = has_twice_in_row(s)
        if !proh && vow3 && tw { count += 1 } 
        if has_triple(s) && has_good_pair(s) { count1 += 1}
    }

    fmt.println(count) 
    fmt.println(count1) 
}