package main
import "core:os"
import "core:fmt"
import "core:strings"
import "core:strconv"
import "core:text/match"

main :: proc() {
    isdigit :: proc(s: string) -> bool {
        matcher := match.matcher_init(s, "^%d+$")
	    _, ok := match.matcher_match(&matcher)
        return ok
    }

    txt, err := os.read_entire_file_from_filename("input.txt")
    if err == false {
        fmt.println("Ошибка открытия файла!")
        os.exit(1)
    }
    test_strings := strings.split_lines(cast(string)txt)
    
    Operation :: enum{Pass, And, Or, Not, Lshift, Rshift}
    Gate :: struct {
        inp1: u16,
        inp1_wire: string,
        inp2: u16,
        inp2_wire: string,
        operation: Operation
    }
    circuit := make(map[string]Gate)
    defer delete(circuit)

    parse_strings :: proc(ts: []string, circuit: ^map[string]Gate) {
        for s in ts {
            ss := strings.split(s, " ")
            size := len(ss)
            out_wire, operand, inp1_wire, inp2_wire := ss[size - 1], "","",""

            switch size {
            case 3: // dq -> dr  |  44430 -> b
                if isdigit(ss[0]) {
                    circuit[out_wire] = Gate{ inp1 = cast(u16)strconv.atoi(ss[0]), operation = .Pass }
                } else {
                    circuit[out_wire] = Gate{ inp1_wire = ss[0], operation = .Pass }
                }            
            case 4: // NOT dq -> dr
                circuit[out_wire] = Gate{ inp1_wire = ss[1], operation = .Not }
            case 5: // ep OR eo -> eq | ep AND eo -> eq 
                gt := Gate{}
                switch ss[1] {
                    case "OR":
                        gt.operation = .Or
                    case "AND":
                        gt.operation = .And
                    case "RSHIFT":
                        gt.operation = .Rshift
                    case "LSHIFT":
                        gt.operation = .Lshift
                    case:
                        fmt.println(ss[1])
                }
                if isdigit(ss[0]) {
                    gt.inp1 = cast(u16)strconv.atoi(ss[0])
                } else {
                    gt.inp1_wire = ss[0]
                }
                if isdigit(ss[2]) {
                    gt.inp2 = cast(u16)strconv.atoi(ss[2])
                } else {
                    gt.inp2_wire = ss[2]
                }      
                circuit[out_wire] = gt            
            case: // default
                break
            }
        } // for
    }


    eval_gate :: proc(inp_gate: string, circuit: ^map[string]Gate) -> u16 {
        op1, op2 : u16
        inp := &circuit[inp_gate]
        //fmt.println(inp_gate, inp)
        switch inp.operation {
            case .Pass:
                if inp.inp1_wire == "" { return inp.inp1 }
                else { inp1 := eval_gate(inp.inp1_wire, circuit); inp.inp1 = inp1; inp.inp1_wire = ""; return inp1 }
            case .And:                
                if inp.inp1_wire == "" { op1 = inp.inp1 }
                else { op1 = eval_gate(inp.inp1_wire, circuit); inp.inp1_wire = ""; inp.inp1 = op1 }
                if inp.inp2_wire == "" { op2 = inp.inp2 }
                else { op2 = eval_gate(inp.inp2_wire, circuit); inp.inp2_wire = ""; inp.inp2 = op2 }
                return op1 & op2                
            case .Or:
                if inp.inp1_wire == "" { op1 = inp.inp1 }
                else { op1 = eval_gate(inp.inp1_wire, circuit); inp.inp1_wire = ""; inp.inp1 = op1 }
                if inp.inp2_wire == "" { op2 = inp.inp2 }
                else { op2 = eval_gate(inp.inp2_wire, circuit); inp.inp2_wire = ""; inp.inp2 = op2 }
                return op1 | op2            
            case .Not:
                if inp.inp1_wire == "" { return ~ inp.inp1 }
                else { inp1 := ~ eval_gate(inp.inp1_wire, circuit); inp.inp1_wire = ""; inp.inp1 = inp1; return inp1 }            
            case .Lshift:
                if inp.inp1_wire == "" { op1 = inp.inp1 }
                else { op1 = eval_gate(inp.inp1_wire, circuit); inp.inp1_wire = ""; inp.inp1 = op1 }
                if inp.inp2_wire == "" { op2 = inp.inp2 }
                else { op2 = eval_gate(inp.inp2_wire, circuit); inp.inp2_wire = ""; inp.inp2 = op2 }
                return op1 << op2
            case .Rshift:
                if inp.inp1_wire == "" { op1 = inp.inp1 }
                else { op1 = eval_gate(inp.inp1_wire, circuit); inp.inp1_wire = ""; inp.inp1 = op1 }
                if inp.inp2_wire == "" { op2 = inp.inp2 }
                else { op2 = eval_gate(inp.inp2_wire, circuit); inp.inp2_wire = ""; inp.inp2 = op2}
                return op1 >> op2
            case:
                fmt.println("Error gate!\n", inp)
                os.exit(1)
        } 
        return 0
    }

    parse_strings(test_strings, &circuit)
    res_a := eval_gate("a", &circuit)
    fmt.println(res_a) // [1] 3176

    parse_strings(test_strings, &circuit) 
    (&circuit["b"]).inp1 = res_a
    fmt.println(eval_gate("a", &circuit)) // [2] 14710
    
}