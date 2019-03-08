function print_table(node)
    -- to make output beautiful
    local function tab(amt)
        local str = ""
        for i=1,amt do
            str = str .. "\t"
        end
        return str
    end

    local cache, stack, output = {},{},{}
    local depth = 1
    local output_str = "{\n"

    while true do
        local size = 0
        for k,v in pairs(node) do
            size = size + 1
        end

        local cur_index = 1
        for k,v in pairs(node) do
            if (cache[node] == nil) or (cur_index >= cache[node]) then

                if (string.find(output_str,"}",output_str:len())) then
                    output_str = output_str .. ",\n"
                elseif not (string.find(output_str,"\n",output_str:len())) then
                    output_str = output_str .. "\n"
                end

                -- This is necessary for working with HUGE tables otherwise we run out of memory using concat on huge strings
                table.insert(output,output_str)
                output_str = ""

                local key
                if (type(k) == "number" or type(k) == "boolean") then
                    key = "["..tostring(k).."]"
                else
                    key = "['"..tostring(k).."']"
                end

                if (type(v) == "number" or type(v) == "boolean") then
                    output_str = output_str .. tab(depth) .. key .. " = "..tostring(v)
                elseif (type(v) == "table") then
                    output_str = output_str .. tab(depth) .. key .. " = {\n"
                    table.insert(stack,node)
                    table.insert(stack,v)
                    cache[node] = cur_index+1
                    break
                else
                    output_str = output_str .. tab(depth) .. key .. " = '"..tostring(v).."'"
                end

                if (cur_index == size) then
                    output_str = output_str .. "\n" .. tab(depth-1) .. "}"
                else
                    output_str = output_str .. ","
                end
            else
                -- close the table
                if (cur_index == size) then
                    output_str = output_str .. "\n" .. tab(depth-1) .. "}"
                end
            end

            cur_index = cur_index + 1
        end

        if (size == 0) then
            output_str = output_str .. "\n" .. tab(depth-1) .. "}"
        end

        if (#stack > 0) then
            node = stack[#stack]
            stack[#stack] = nil
            depth = cache[node] == nil and depth + 1 or depth - 1
        else
            break
        end
    end

    -- This is necessary for working with HUGE tables otherwise we run out of memory using concat on huge strings
    table.insert(output,output_str)
    output_str = table.concat(output)

    print(output_str)
end


local function is_hook_field(name)
    return name == "prestart" or name == "prestop" or name == "prebuild" or name == "preremove"
end




local function add_field(dict, depth, name, options)
    
    if name == "name" then
        if type(options) ~= "string" then
            error("not a string")
        end
    
    elseif is_hook_field(name) then
        if type(options) ~= "function" then
            error(name.." should be a function")
        end

    elseif name == "dependOn" or name == "dependencies" then
        if type(options) ~= "table" then 
            error(name.." should be a table")
        end
    end

    local len = #depth
    for key,val in ipairs(depth) do
        
        if key == len then
            if val == "dependOn" then
                val = "dependencies"
            end
            dict[val] = options
        elseif dict[val] == nil then
            dict[val] = {}
        end
    end
   
    
end


function service(fn)

    local description = {}
    local depth = {}
    setfenv(fn, setmetatable({
    print = print
    }, {
    __index = function(self, field_name)
        return function(opts)
            table.insert(depth, field_name)
            local ret = add_field(description, depth, field_name, opts)
            table.remove(depth)
            return ret
        end
    end,
    __newindex = function (self, field_name, value)
        table.insert(depth, field_name)
        add_field(description, depth, field_name, value)
        table.remove(depth)    
    end,
    
  }))

  local ret = fn()
  
  _registerService(description)

  return ret

end

