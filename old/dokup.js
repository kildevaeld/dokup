
var name = "livejazz-admin";
var mysql = "livejazz-mysql-" + (process.env.RUNN_ENV||'development');

module.exports = {
    name: name,
    link: {
        $staging: {
            mysql: mysql
        }
    },
    $development: {
        publish: "3000:3000"
    },
    tty: true,
    attachStdin: true,
    envFile: "",
    env: {

    },
    initialize: function () {
        return fetch('http://jsonip.com').then(function (res) { return res.json()})
        .then(function (json) {
            console.log(JSON.stringify(json))
        })
    },
    prestart: function () {
        //console.log('prestart')
    },
    prebuild: function () {
        //console.log('prebuild')
    },
    postbuild: function () {
        //console.log('post build')
    },
    build: {
        path: process.cwd() + "/admin"
    },
    dependencies: [
        {
            name: mysql,
            phase: ["development", "staging"],
            $darwin: {
                publish: ['3306:3306']
            },
            image: 'alpine',
            cmd: ['sh'],
            tty: true,
            prebuild: function () {

            },
            /*build: {
                path: process.cwd() + "/mysql"
            },*/
            dependencies: [
                {name: 'nginx', phase: ['staging']}
            ]
        }, {
            name: "nginx",
            image: "softshag/alpine-nginx-confd",
            phase: ["production", "staging"],
            volumne: [
                "./:/etc/confd"
            ],
            prestart: function () {
                console.log('prestart')
            },
            publish: ["80:80","443:443"]

        }
    ]
}