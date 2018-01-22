# Read me
----------------------------

## Install
to get started please install those dependencies :

```
// router
go get github.com/gin-gonic/gin
// orm
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite
// excel bridge
go get github.com/360EntSecGroup-Skylar/excelize
```

## Running

there's 2 way to running the app

1. Using Visual studio code
2. Using command `go run`

---------------------------

1. Visual studio code

configure launch.json

```
{
	"version": "0.2.0",
	"configurations": [
			{
					"name": "Launch",
					"type": "go",
					"request": "launch",
					"mode": "exec",
					"port": 2345,
					"host": "127.0.0.1",
					"program": "${workspaceRoot}/debug",
					"preLaunchTask": "build-debug",
					"env": {},
					"args": [],
					"showLog": true
			}
	]
}
```

configure tasks.json

```
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "build-debug",
            "type": "shell",
            "command": "go",
            "group": "build",
            "presentation": {
                "echo": true,
                "reveal": "never",
                "focus": false,
                "panel": "shared"
            },
            "args": [
                "build",
                "-i",
                "-gcflags",
                "'-N -l'"
            ],
            "linux": {
                "args": [
                    "-o",
                    "debug",
                    "${workspaceRoot}/main.go",
                    "${workspaceRoot}/router.go"
                ]
            },
            "osx": {
                "args": [
                    "-o",
                    "debug",
                    "${workspaceRoot}/main.go",
                    "${workspaceRoot}/router.go"
                ]
            },
            "windows": {
                "args": [
                    "-o",
                    "debug.exe",
                    "\"${workspaceRoot}\\main.go\"",
                    "\"${workspaceRoot}\\router.go\""
                ]
            },
            "problemMatcher": [
                "$go"
            ]
        }
    ]
}
```

---------------------------------

2. Command Go

`go run main.go router.go` // because router.go and main.go are in the same package


----------------------------------

*Notes : the database are included in commit and no password for database
