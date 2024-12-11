/*
I created a grid of integers for the counts for part 2.
Loop through 0-9 and process the whole grid looking for those digits. If 0, set count to 1.
Otherwise total the counts of the surrounding cells that are one lower and set the count to that,
and if 9 then add that to the total. About 0.4ms
*/
func part2(contents string) interface{} {
    area := aoc.ParseArea(contents)
    counts := make([]int, len(area.Data))

    totalAround := func(position int) int {
        h := area.GetIndex(position)
        r, c := area.IndexToRowCol(position)
        total := 0
        for _, v := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
            r2, c2 := r+v[0], c+v[1]
            if area.Inside(r2, c2) && (area.Get(r2, c2) == (h - 1)) {
                total += counts[area.RowColToIndex(r2, c2)]
            }
        }
        return total
    }

    sum := 0
    for h := byte(0x30); h <= 0x39; h++ {
        for p := range area.Data {
            if h == byte(0x30) {
                counts[p] = 1
            } else {
                if area.GetIndex(p) == h {
                    counts[p] = totalAround(p)
                    if h == byte(0x39) {
                        sum += counts[p]
                    }
                }
            }
        }
    }

    return sum
}