package main

import (
	"errors"
	"fmt"
)

func main() {
	err := runWithRecover()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func runWithRecover() error {
	defer func() {
		if r := recover(); r != nil {
			// 将 panic 转换为 error
			errors.New(fmt.Sprintf("Recovered from panic: %v", r))
		}
	}()

	// 模拟触发 panic
	panic("Something went wrong!")

	// 这里的代码不会被执行，因为已经发生了 panic
	return nil
}
