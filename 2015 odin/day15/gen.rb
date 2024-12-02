def to_sum_k(n, k)
    return [ [k] ] if n == 1 && k > 0
    return [] if n > k or n <= 0
        
    res = []
    (1..k).each do |i|
        sub_results = to_sum_k(n-1, k-i)
        for sub in sub_results
            res << sub + [i]
        end
    end
    res
end    


def to_sum2(n, k)
    (1..k).each do |a|
        (1..(k - a)).each do |b|
            (1..(k - a - b)).each do |c|
                d = k - a - b - c
                next if d < 1
                p [a, b, c, d]
            end
        end
    end
end

print to_sum2(3, 10)