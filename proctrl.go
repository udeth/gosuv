package main

// import (
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"sync"
// 	"syscall"
// 	"time"

// 	"github.com/codeskyblue/kexec"
// 	"github.com/qiniu/log"
// )

// const (
// 	ST_RUNNING   = "RUNNING"
// 	ST_STOPPED   = "STOPPED"
// 	ST_FATAL     = "FATAL"
// 	ST_RETRYWAIT = "RETRYWAIT" // some like python-supervisor EXITED
// )

// type Event int

// const (
// 	EVENT_START = Event(iota)
// 	EVENT_STOP
// )

// type Program struct {
// 	Name         string   `yaml:"name"`
// 	Command      string   `yaml:"command"`
// 	Environ      []string `yaml:"environ"`
// 	Dir          string   `yaml:"directory"`
// 	AutoStart    bool     `yaml:"autostart"` // change to *bool, which support unexpected
// 	StartRetries int      `yaml:"startretries"`
// 	StartSeconds int      `yaml:"startsecs"`
// 	LogDir       string   `yaml:"logdir"`
// }

// func (p *Program) buildCmd() *kexec.KCommand {
// 	cmd := kexec.CommandString(p.Command) // Not tested here, I think it should work
// 	// cmd := kexec.Command(p.Command[0], p.Command[1:]...)
// 	cmd.Dir = p.Dir
// 	cmd.Env = append(os.Environ(), p.Environ...)
// 	return cmd
// }

// type Process struct {
// 	*kexec.KCommand `json:"-"`
// 	Status          string         `json:"state"`
// 	Sig             chan os.Signal `json:"-"`
// 	Info            *Program       `json:"conf"`
// 	mu              sync.Mutex

// 	retry int
// 	stopC chan bool
// }

// func NewProcess(info *Program) *Process {
// 	// set default values
// 	if info.StartRetries <= 0 {
// 		info.StartRetries = 3
// 	}
// 	if info.StartSeconds <= 0 {
// 		info.StartSeconds = 3
// 	}
// 	return &Process{
// 		Status: ST_STOPPED,
// 		Sig:    make(chan os.Signal),
// 		Info:   info,
// 		stopC:  make(chan bool),
// 	}
// }

// func (p *Process) setStatus(st string) {
// 	p.Status = st
// }

// // STOP and START Should not run parallel
// func (p *Process) Operate(event Event) {
// 	p.mu.Lock()
// 	defer p.mu.Unlock()

// 	switch event {
// 	case EVENT_START:
// 		if p.Status == ST_STOPPED || p.Status == ST_FATAL {
// 			go p.RunWithRetry()
// 		}
// 	case EVENT_STOP:
// 		if p.Status == ST_RUNNING || p.Status == ST_RETRYWAIT {
// 			p.stop()
// 		}
// 	}
// }

// func (p *Process) LogFilePath() string {
// 	logDir := p.Info.LogDir
// 	if logDir == "" {
// 		logDir = filepath.Join(p.Dir, "logs")
// 	}
// 	return filepath.Join(logDir, p.Info.Name+".log")
// }

// func (p *Process) createLog() (*os.File, error) {
// 	logDir := filepath.Join(p.Info.LogDir, "logs")
// 	os.MkdirAll(logDir, 0755) // just do it, err ignore it
// 	logFile := p.LogFilePath()
// 	return os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
// }

// func (p *Process) RunWithRetry() {
// 	var err error
// 	for p.retry = 0; p.retry < p.Info.StartRetries+1; p.retry += 1 {
// 		// wait program to exit
// 		errC := GoFunc(p.run)

// 	PROGRAM_WAIT:
// 		// Here is RUNNING State
// 		select {
// 		case err = <-errC:
// 			log.Info(p.Info.Name, err)
// 		case <-time.After(time.Second * time.Duration(p.Info.StartSeconds)): // reset retry
// 			p.retry = 0
// 			goto PROGRAM_WAIT
// 		case <-p.stopC:
// 			return
// 		}

// 		// Enter RETRY_WAIT State
// 		if p.retry < p.Info.StartRetries {
// 			p.setStatus(ST_RETRYWAIT)
// 			select {
// 			case <-p.stopC:
// 				return
// 			case <-time.After(time.Second * 2):
// 			}
// 		}
// 	}
// 	p.setStatus(ST_FATAL)
// }

// func (p *Process) start() error {
// 	p.KCommand = p.Info.buildCmd()
// 	logFd, err := p.createLog()
// 	if err != nil {
// 		return err
// 	}
// 	p.Cmd.Stdout = logFd
// 	p.Cmd.Stderr = logFd
// 	return p.Cmd.Start()
// }

// func (p *Process) run() (err error) {
// 	if err = p.start(); err != nil {
// 		return
// 	}
// 	p.setStatus(ST_RUNNING)
// 	defer func() {
// 		if out, ok := p.Cmd.Stdout.(io.Closer); ok {
// 			out.Close()
// 		}
// 		log.Warnf("Process finish: %v", err)
// 	}()
// 	err = p.Wait()
// 	return
// }

// func (p *Process) stop() error {
// 	select {
// 	case p.stopC <- true: // stopC may not recevied, when in FATAL, in a very low probality
// 	case <-time.After(time.Millisecond * 200): // 0.2s
// 	}
// 	p.Terminate(syscall.SIGKILL)
// 	p.setStatus(ST_STOPPED)
// 	return nil
// }
