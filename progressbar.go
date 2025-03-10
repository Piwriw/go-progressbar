package progressbar

import (
	"encoding/json"
	"errors"
	"log/slog"
	"strings"

	"github.com/schollz/progressbar/v3"
)

const (
	defaultMetricPort = "0.0.0.0:19999"
)

type ProgressBar struct {
	total int
	bar   *progressbar.ProgressBar
	opts  Options
	tasks []ProgressTask
	err   []error
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		total: 0,
		tasks: make([]ProgressTask, 0),
	}
}

// Error 返回所有累积的错误
func (p *ProgressBar) Error() error {
	if len(p.err) == 0 {
		return nil
	}
	return errors.New(strings.Join(p.errorMessages(), "; "))
}

// errorMessages 获取所有错误消息
func (p *ProgressBar) errorMessages() []string {
	messages := make([]string, len(p.err))
	for i, err := range p.err {
		messages[i] = err.Error()
	}
	return messages
}

func (p *ProgressBar) Create() *ProgressBar {
	p.genericBar()
	return p
}

func (p *ProgressBar) Total(total int) *ProgressBar {
	p.total = total
	return p
}

func (p *ProgressBar) Options(opts *Options) *ProgressBar {
	p.opts = *opts
	return p
}

func (p *ProgressBar) Tasks(tasks ...ProgressTask) *ProgressBar {
	p.tasks = append(p.tasks, tasks...)
	return p
}

// Prefix sets the prefix of the progress bar
func (p *ProgressBar) Prefix(prefix string) {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return
	}
	p.bar.Describe(prefix)
}

// Suffix sets the suffix of the progress bar
func (p *ProgressBar) Suffix(suffix string) {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return
	}
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

func (p *ProgressBar) AddBar() *ProgressBar {
	return &ProgressBar{
		bar: progressbar.NewOptions(p.total, p.opts.options...),
	}
}

func (p *ProgressBar) genericBar() {
	p.bar = progressbar.NewOptions(p.total, p.opts.options...)
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

func (p *ProgressBar) AutoRun() error {
	if p.Error() != nil {
		return p.Error()
	}
	for _, task := range p.tasks {
		if err := callFunc(task.fn, task.params...); err != nil {
			slog.Error("AutoRun", "err", err)
			return p.Exit()
		}
		if err := p.Next(); err != nil {
			return err
		}
	}
	return nil
}

// Next will increase the progress bar by 1
// example:step + 1
func (p *ProgressBar) Next() error {
	if p.Error() != nil {
		return p.Error()
	}
	return p.Add(1)
}

// Add will add the specified amount to the progressbar
func (p *ProgressBar) Add(num int) error {
	if p.Error() != nil {
		return p.Error()
	}
	return p.bar.Add(num)
}

// Finish will fill the bar to full
func (p *ProgressBar) Finish() error {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return p.Error()
	}
	return p.bar.Finish()
}

// Exit will exit the bar to keep current state
func (p *ProgressBar) Exit() error {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return p.Error()
	}
	return p.bar.Exit()
}

// Clear erases the progress bar from the current line
func (p *ProgressBar) Clear() error {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return p.Error()
	}
	return p.bar.Clear()
}

// Set will set the bar to a current number
func (p *ProgressBar) Set(step int) error {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return p.Error()
	}
	return p.bar.Set(step)
}

// IsFinished returns true if progress bar is completed
func (p *ProgressBar) IsFinished() bool {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return false
	}
	return p.bar.IsFinished()
}

// IsStarted returns true if progress bar is started
func (p *ProgressBar) IsStarted() bool {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return false
	}
	return p.bar.IsStarted()
}

// State returns the current state
func (p *ProgressBar) State() progressbar.State {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return progressbar.State{}
	}
	return p.bar.State()
}

// Describe will change the description shown before the progress, which
// can be changed on the fly (as for a slow running process).
func (p *ProgressBar) Describe(description string) {
	if p.bar == nil {
		p.err = append(p.err, ErrNilBar)
		return
	}
	p.bar.Describe(description)
}

func (p *ProgressBar) JSON() string {
	data, _ := json.Marshal(p.bar.State())
	return string(data)
}
