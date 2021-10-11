// signalchanel - Task#2
// Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее, чем за одну секунду (установить таймаут).
package signalchanel

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TreatmentSignal create notify context and send SIGNTERM signal to signal channel
func TreatmentSignal() error {
	ctxNotify, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()
	ctxTimeout, cancel := context.WithTimeout(ctxNotify, time.Second*1)
	defer cancel()

	//wg := sync.WaitGroup{}
	//defer wg.Wait()

	// do something endlessly
	//wg.Add(1)
	go func() {
		for {
		}
		//wg.Done()
	}()

	// send SIGNTERM signal in to signal channel
	if err := sendSignTerm(); err != nil {
		return err
	}

	select {
	case <-ctxTimeout.Done():
		return ctxTimeout.Err()
	case <-ctxNotify.Done():
		return ctxNotify.Err()
		//default:
		//wg.Wait()
	}

	return nil
}

// sendSignTerm send SIGNTERM signal to current process
func sendSignTerm() error {
	// test delay send SIGNTERM
	//time.Sleep(time.Second * 1)

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		return err
	}
	if err := p.Signal(syscall.SIGTERM); err != nil {
		return err
	}

	return nil
}
