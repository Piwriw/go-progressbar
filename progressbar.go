package progressbar

import (
	"encoding/json"
	"github.com/schollz/progressbar/v3"
	"log/slog"
	"time"
)

const (
	defaultMetricPort = "0.0.0.0:19999"
)

type ProgressBar struct {
	total   int
	bar     *progressbar.ProgressBar
	options Options
}

func (p *ProgressBar) Total(total int) *ProgressBar {
	p.total = total
	return p
}

func (p *ProgressBar) AutoUpdate(interval time.Duration, stopChan chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			p.Next()
		case <-stopChan:
			return
		}
	}
}

// Prefix sets the prefix of the progress bar
func (p *ProgressBar) Prefix(prefix string) {
	p.bar.Describe(prefix)
}

// Suffix sets the suffix of the progress bar
func (p *ProgressBar) Suffix(suffix string) {
	p.bar.Describe(p.bar.State().Description + suffix)
}

// Metric starts an HTTP server dedicated to serving progress bar updates. This allows you to
// display the status in various UI elements, such as an OS status bar with an `xbar` extension.
// It is recommended to run this function in a separate goroutine to avoid blocking the main thread.
//
// hostPort specifies the address and port to bind the server to, for example, "0.0.0.0:19999".
func (p *ProgressBar) Metric(hostPort string) {
	if hostPort == "" {
		hostPort = defaultMetricPort
	}
	p.bar.StartHTTPServer(hostPort)
}

func Add(total int, ps *Options) *ProgressBar {
	return &ProgressBar{
		bar: progressbar.NewOptions(total, ps.options...),
	}
}

type ProgressTask struct {
	fn     any
	params []any
}

func NewProgressTask(fn any, params ...any) ProgressTask {
	return ProgressTask{
		fn:     fn,
		params: params,
	}
}

func AutoRun(ps *Options, tasks ...ProgressTask) error {
	bar := &ProgressBar{
		bar: progressbar.NewOptions(len(tasks), ps.options...),
	}
	for _, task := range tasks {
		if err := callFunc(task.fn, task.params...); err != nil {
			slog.Error("AutoRun", "err", err)
			return bar.Exit()
		}
		if err := bar.Next(); err != nil {
			return err
		}
	}
	return nil
}

// Next will increase the progress bar by 1
// example:step + 1
func (p *ProgressBar) Next() error {
	return p.Add(1)
}

// Add will add the specified amount to the progressbar
func (p *ProgressBar) Add(num int) error {
	return p.bar.Add(num)
}

// Finish will fill the bar to full
func (p *ProgressBar) Finish() error {
	return p.bar.Finish()
}

// Exit will exit the bar to keep current state
func (p *ProgressBar) Exit() error {
	return p.bar.Exit()
}

// Clear erases the progress bar from the current line
func (p *ProgressBar) Clear() error {
	return p.bar.Clear()
}

// Set will set the bar to a current number
func (p *ProgressBar) Set(step int) error {
	return p.bar.Set(step)
}

// IsFinished returns true if progress bar is completed
func (p *ProgressBar) IsFinished() bool {
	return p.bar.IsFinished()
}

// IsStarted returns true if progress bar is started
func (p *ProgressBar) IsStarted() bool {
	return p.bar.IsStarted()
}

// State returns the current state
func (p *ProgressBar) State() progressbar.State {
	return p.bar.State()
}

// Describe will change the description shown before the progress, which
// can be changed on the fly (as for a slow running process).
func (p *ProgressBar) Describe(description string) {
	p.bar.Describe(description)
}

func (p *ProgressBar) JSON() string {
	data, _ := json.Marshal(p.bar.State())
	return string(data)
}
