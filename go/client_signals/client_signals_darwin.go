package client_signals

import "os"
import "os/signal"
import "syscall"

var handledSignals = []os.Signal{syscall.SIGINFO}

func HandleSignals(pid int) {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, handledSignals...)
	go func() {
		for sig := range sigChan {
			if sig == syscall.SIGINFO {
				syscall.Kill(pid, syscall.SIGINFO)
				syscall.Kill(os.Getpid(), syscall.SIGINFO)
			}
		}
	}()
}
