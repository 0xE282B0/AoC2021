-- https://stackoverflow.com/a/11204889

bits = {}
count = 0
for line in io.lines('data/input.txt') do 
    count = count +1
    for i = 1, #line do
        local j = #line-i
        if bits[j] == nil then
            bits[j] = 0
        end
        local b = line:sub(i,i)
        if b == '1' then
            bits[j] = bits[j] +1
        end
    end
end

gamma = 0
epsilon = 0
for key, value in pairs(bits) do
    print(key, " -- ", math.floor((value/count)+.5), 2^(key))
    if math.floor((value/count)+.5) == 1 then
        gamma = gamma + (2^(key))
    else
        epsilon = epsilon + (2^(key))
    end

end

print('count',count, 'gamma', gamma, 'epsilon', epsilon, 'mul', gamma*epsilon)