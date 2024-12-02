def sums(length, total_sum):
    if length == 1:
        yield (total_sum,)
    else:
        for value in range(1, total_sum + 1):
            for pr in sums(length - 1, total_sum - value):
                yield (value,) + pr

L = list(sums(3,10))
print('total permutations:',len(L))
for i in L[:10] + L[-10:]:
    print(i) 
