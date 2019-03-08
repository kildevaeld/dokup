
local function prestart_fn()
    print("running prebuild")
end

service(function () 
    
    docker {
        image = "mhart/alpine-node:7.10.0",
        steps = [[
        RUN apk add --update
        
        ]],
        TTY = true,
        attachStdin = true,
        attachStdout = true,
        attachStderr = true,
        openStdin = true,
        openStdinOnce = true,
        stream = true,
        link = {
            "mysql:mysql"    
        },
        cmd = {"sh"}
    }

    prebuild = prestart_fn

    name = "boellefesten"

    dependOn {
        "mysql"
    }

end)

service(function ()
    name = "mysql"
    docker {
        image = "mysql",
        publish = {
            "3306:3306"
        },
        env = {
            "MYSQL_ROOT_PASSWORD=password"
        }
    }
    
end)