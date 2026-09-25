fn factorial(n: int): int
    if n <= 1 then
        return 1
    else
        return n * factorial(n - 1)
    end
end

fn main(): int
    let x: int := 5
    print(factorial(x))

    var total: int := 0
    let nums: [5]int := [1, 2, 3, 4, 5]

    for i in 0..5 do
        total := total + nums[i]
    end
    print(total)
    
