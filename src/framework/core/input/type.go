package input

import (
	"configcenter/src/framework/core/output"
	"context"
	"time"
)

// InputerStatus the inputer status type definition.
type InputerStatus int

// InputerType the inputer type definition
type InputerType int

// InputerKey the inputer name
type InputerKey string

// MapInputer inputer object
type MapInputer map[InputerKey]*wrapInputer

// ExceptionFunc the exception callback
type ExceptionFunc func(data interface{}, errMsg error)

// Inputer status type
const (

	// NormalStatus Inputer just created
	NormalStatus InputerStatus = iota
	// WaitingToRunStatus Inputer is waiting to run
	WaitingToRunStatus
	// RunningStatus Inputer is running
	RunningStatus
	// StoppingStatus Inputer normal exiting
	StoppingStatus
	// StoppedStatus Inputer normal exit
	StoppedStatus
	// ExceptionExitStatus Inputer abnormal exit
	ExceptionExitStatus
)

// InputerType definition
const (
	// ExecuteOnce only execute onece, non-blocking
	ExecuteOnce InputerType = iota

	// ExecuteTiming timing execute
	ExecuteTiming

	// ExecuteLoop loop execution does not exit, blocking
	ExecuteLoop
)

// InputerParams the inputer params
type InputerParams struct {
	IsTiming  bool
	Frequency time.Duration
	Target    Inputer
	Kind      InputerType
	Putter    output.Puter
	Exception ExceptionFunc
}

// Manager is the interface that must be implemented by every input manager.
type Manager interface {

	// AddInputer add a new inputer
	AddInputer(params InputerParams) InputerKey

	// RemoveInputer remove the Inputer by a WorkerKey
	RemoveInputer(key InputerKey)

	// Run start the business cycle until the stop method is called.
	Run(ctx context.Context, cancel context.CancelFunc)

	// Stop
	Stop() error
}

// Inputer is the interface that must be implemented by every Inputer.
type Inputer interface {

	// Description the Inputer description.
	// This information will be printed when the Inputer is abnormal, which is convenient for debugging.
	Name() string

	// Run execute the user logics
	Run() interface{}

	// Stop stop the run function
	Stop() error
}