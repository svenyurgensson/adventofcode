package main
import "core:os"
import "core:fmt"

main :: proc() {
    txt, err := os.read_entire_file_from_filename("input-d1.txt")
    if err == false {
        fmt.println("Ошибка открытия файла!")
        os.exit(1)
    }

    count := 0

    for ch, idx in txt {
        if ch == '(' do count += 1
        if ch == ')' do count -= 1
        // step2
        if count == -1 {
            fmt.println(idx+1) // 1795
            os.exit(0)
        }
        // /step2
    }
    fmt.println(count)
}
