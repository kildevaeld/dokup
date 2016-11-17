var Builder = require("docker.builder");
var ui = require("ui");
var n = Builder.Notification;
var argv = require("minimist")(process.argv.slice(1));
var fs = require("fs");
var fp = process.cwd() + "/dokup.js";

return fs.stat(fp)
    .then(function (stat) {

        var dokup = require(fp);

        var env = get_env()

        var arg
        if (argv[env] || argv[env[0]]) arg = argv[env] || argv[env[0]]
        else arg = argv._.length ? argv._[0] : ''

        if (!!!~['start', 'stop', 'remove', 'build'].indexOf(arg)) {
            console.log("Usage: dokup <s(tart)|stop|r(emove)|b(uild)>")
            process.exit(1);
        }

        return run(dokup, arg, get_env());

    }).catch(function (e) {
        console.log("No dokup.js exists!", e);
    });


function get_env() {
    if (argv.production || argv.p) return "production";
    if (argv.staging || argv.s) return "staging";
    return 'development';
}

function run(mod, cmd, env) {
    Builder.createBuilder(mod, env)
        .then(function (builder) {

            initUI(builder);

            switch (cmd) {

                case "start": return builder.start(true);

                case "remove": return builder.remove(true, true);
                case "stop": return builder.stop();

                case "build": return builder.build(true);
                default:
                    console.log("Usage: dokup <start|stop|remove>")
                    process.exit(1);
            }


        }).catch(function (e) {
            console.log(e);
        })
}

function initUI(builder) {
    var p;
    builder.on(Builder.NotificationEvent, function (e, m) {
        var str = Builder.Notification[e] + " " + (Array.isArray(m) ? m.map(function (z) {
            return z.name
        }).join(" ") : m.name)
        switch (e) {
            case n.Building:
            case n.Starting:
            case n.Creating:
            case n.Stopping:
                p = ui.Process(str + " ...");
                p.Start()
                break;
            case n.Build:
            case n.Started:
            case n.Created:
            case n.Stopped:
                p.Success("done");
                p = null;
                break;
        }

    })
}