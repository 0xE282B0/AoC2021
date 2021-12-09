
function task2 (type)
    prefix = ''
    notUnique = true
    while notUnique
    do
        counter = 0
        bit1 = 0  
        notUnique = false
        result = ''
        for line in io.lines('data/input.txt') do
            -- https://stackoverflow.com/a/12929685
            if line:find('^' .. prefix) == nil then goto continue end
            if result == '' then
                result = line
            else
                result = ''
                notUnique = true
            end
            -- calculate most common bit
            counter = counter + 1
            if line:sub(#prefix+1,#prefix+1) == '1' then bit1 = bit1+1 end
            ::continue::
        end
        if notUnique then
            mostCommon = (bit1/counter)+.5 
            if type == 'oxy' then 
                if mostCommon >= 1 then
                    prefix = prefix .. '1'
                else
                    prefix = prefix .. '0'
                end
            else
                if mostCommon < 1 then
                    prefix = prefix .. '1'
                else
                    prefix = prefix .. '0'
                end    
            end
            print('prefix',prefix)
        else
            return tonumber(result,2)
        end
    end
end

oxy = task2 ('oxy')
co2 = task2 ('co2')
print('oxy',oxy,'co2',co2,'mul',oxy*co2)