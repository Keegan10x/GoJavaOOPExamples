package example

import (
	"fmt"
)

// define type and constructor
type Engine struct {
	capacity   int
	cylinders  int
	engineType string
	isRunning  bool
}

// NewEngine is a contstructor that returns an instance of an engine type
func NewEngine(inputCapacity int, inputCylinders int, inputEngineType string) *Engine {
	e := &Engine{
		capacity:   inputCapacity,
		cylinders:  inputCylinders,
		engineType: inputEngineType,
	}
	return e
}

func (e *Engine) Start() {
	if e.isRunning {
		fmt.Println("already running...")
		return
	}
	e.isRunning = true
	fmt.Println("engine has started...")
}

func (e *Engine) Stop() {
	if !e.isRunning {
		fmt.Print("already stopped...")
		return
	}
	e.isRunning = false
	fmt.Print("engine has stopped...")
}
