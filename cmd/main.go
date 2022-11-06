package main

import (
    "go-mongo-example/mgolearn"
    "go-mongo-example/mongolearn"
    "go-mongo-example/qmgolearn"
)

func main() {
    mgolearn.TestDial()

    qmgolearn.TestConn()

    mongolearn.TestConnUseDb()
}
