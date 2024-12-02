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

    //initial := []u8{'h', 'e', 'p', 'x', 'c', 'r', 'r', 'q'}  // part one -> hepxxyzz
    initial := []u8{'h','e','p','x','x','y','z','z'} // part 2 -> heqaabcc
    //initial := []u8{ 'g','h','i','j','k','l','m','n' }

    check_pwd :: proc(pwd : []u8) -> bool {
        flag := false
        i := 0
        for i in 0..<(len(pwd) - 3) {
            if pwd[i] == (pwd[i+1] - 1) && pwd[i+1] == (pwd[i+2] - 1) { flag = true } 
        }
        if flag == false { return false }

        for i in 0..<(len(pwd)) {
            if pwd[i] == 'i' || pwd[i] == 'o' || pwd[i] == 'l' { return false } 
        }       

        letter : u8 = 0
        flag = false
        for i in 0..<(len(pwd) - 1) {
            if pwd[i] == pwd[i+1] { letter = pwd[i]; flag = true } 
        }
        if flag == false { return false }
        flag = false
        for i in 0..<(len(pwd) - 1) {
            if pwd[i] == pwd[i+1] && pwd[i] != letter { flag = true } 
        }
        if flag == false { return false }

        return true
    }

    next_pwd :: proc(pwd : []u8) -> []u8 {
        n_pwd := pwd
        borrow := true
        i := len(pwd) - 1
        for {
            if borrow {
                borrow = false
                n_pwd[i] += 1
                
                if n_pwd[i] > 'z' { 
                    borrow = true
                    n_pwd[i] = n_pwd[i] - 'z' + 'a' - 1
                } else { 
                    return n_pwd 
                }
            } else {
                n_pwd[i] += 1
                return n_pwd 
            }

            if i == 0 { break }
            i -= 1
        }
        return n_pwd
    }

    for {
        initial = next_pwd(initial)
        str := cast(string)initial
        if check_pwd(initial) {
            println(str)
            break
        }
    }
}