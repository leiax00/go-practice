package main

import (
	"encoding/json"
	"fmt"
	"listenerror/task"
	"os"
)

func main() {
	//listen_error.Err1Start()
	//try.StartErrUnwrap()
	QuerySingle(10)
	QuerySingle(11) // sql.ErrNoRows

	if err := task.CloseDb(); err != nil {
		fmt.Printf("Failed to close db: %v", err)
		os.Exit(0)
	}
	fmt.Println("Progress Success!!!")
	os.Exit(1)

}

func QuerySingle(userId int) {
	fmt.Printf("=============  Current user Id:%d ============\n", userId)
	user, err := task.QueryAll(userId)
	if err != nil {
		fmt.Printf("error happen:%+v\n", err)
		os.Exit(0)
	}

	fmtRst(user)
}

func fmtRst(user *task.User) {
	if user == nil {
		fmt.Println("No Data")
		return
	}

	tmp, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("parse user error:%v, user id: %d", err, user.Id)
		os.Exit(0)
	}
	fmt.Println(string(tmp))
}
