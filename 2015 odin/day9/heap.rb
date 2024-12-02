arr = [1, 2, 3]

def heap(ary, cnt)
    if cnt > 2 
        heap(ary, cnt - 1)
    end
    for i in 0..(cnt - 2)  do
        if cnt.even? 
            ary[i], ary[cnt-1] = ary[cnt-1], ary[i]
        else
            ary[0], ary[cnt-1] = ary[cnt-1], ary[0]
        end
        p ary
        if cnt > 2
            heap(ary, cnt - 1)
        end
    end
end

p arr
heap(arr, 3)
p arr.permutation.to_a