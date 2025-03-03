package progressbar

import (
	"github.com/schollz/progressbar/v3"
	"io"
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

type Options struct {
	options []progressbar.Option
}

func (p *ProgressBar) Total(total int) *ProgressBar {
	p.total = total
	return p
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

func ProgressOptions() *Options {
	return &Options{
		options: make([]progressbar.Option, 0),
	}
}

func (p *Options) Writer(w io.Writer) *Options {
	p.options = append(p.options, progressbar.OptionSetWriter(w))
	return p
}

// Width sets the width of the bar
func (p *Options) Width(width int) *Options {
	p.options = append(p.options, progressbar.OptionSetWidth(width))
	return p
}

// FullWidth sets the bar to be full width
func (p *Options) FullWidth() *Options {
	p.options = append(p.options, progressbar.OptionFullWidth())
	return p
}

func (p *Options) DisShowTotalBytes(val bool) *Options {
	p.options = append(p.options, progressbar.OptionShowTotalBytes(val))
	return p
}

func (p *Options) SpinnerChangeInterval(interval time.Duration) *Options {
	p.options = append(p.options, progressbar.OptionSetSpinnerChangeInterval(interval))
	return p
}

// SpinnerType sets the type of spinner used for indeterminate bars
func (p *Options) SpinnerType(spinnerType int) *Options {
	p.options = append(p.options, progressbar.OptionSpinnerType(spinnerType))
	return p
}

// SpinnerCustom sets the spinner used for indeterminate bars to the passed
// slice of string
func (p *Options) SpinnerCustom(spinner ...string) *Options {
	p.options = append(p.options, progressbar.OptionSpinnerCustom(spinner))
	return p
}

// Theme sets the elements the bar is constructed with.
// There are two pre-defined themes you can use: ThemeASCII and ThemeUnicode.
func (p *Options) Theme(t progressbar.Theme) *Options {
	p.options = append(p.options, progressbar.OptionSetTheme(t))
	return p
}

// DisEnableVisibility enable the visibility
func (p *Options) DisEnableVisibility() *Options {
	p.options = append(p.options, progressbar.OptionSetVisibility(false))
	return p
}

func (p *Options) RenderBlankState(r bool) *Options {
	p.options = append(p.options, progressbar.OptionSetRenderBlankState(r))
	return p
}

func (p *Options) Throttle(duration time.Duration) *Options {
	p.options = append(p.options, progressbar.OptionThrottle(duration))
	return p
}

// EnableShowCount will also print current count out of total
func (p *Options) EnableShowCount() *Options {
	p.options = append(p.options, progressbar.OptionShowCount())
	return p
}

// EnableShowIts will also print the iterations/second
func (p *Options) EnableShowIts() *Options {
	p.options = append(p.options, progressbar.OptionShowIts())
	return p
}

func (p *Options) Completion(cmpl func()) *Options {
	p.options = append(p.options, progressbar.OptionOnCompletion(cmpl))
	return p
}

// EnableColorCodes enables  support for color codes ,you need there is a color code library
// using mitchellh/colorstring
func (p *Options) EnableColorCodes() *Options {
	p.options = append(p.options, progressbar.OptionEnableColorCodes(true))
	return p
}

func (p *Options) ElapsedTime(elapsedTime bool) *Options {
	p.options = append(p.options, progressbar.OptionSetElapsedTime(elapsedTime))
	return p
}

// DisEnablePredictTime will also attempt to predict the time remaining.
func (p *Options) DisEnablePredictTime() *Options {
	p.options = append(p.options, progressbar.OptionSetPredictTime(false))
	return p
}

// EnableElapsedTimeOnFinish will keep the display of elapsed time on finish.
func (p *Options) EnableElapsedTimeOnFinish() *Options {
	p.options = append(p.options, progressbar.OptionShowElapsedTimeOnFinish())
	return p
}

func (p *Options) SetItsString(iterationString string) *Options {
	p.options = append(p.options, progressbar.OptionSetItsString(iterationString))
	return p
}

func (p *Options) ClearOnFinish() *Options {
	p.options = append(p.options, progressbar.OptionClearOnFinish())
	return p
}

// EnableShowBytes will update the progress bar
// configuration settings to display/hide kBytes/Sec
func (p *Options) EnableShowBytes() *Options {
	p.options = append(p.options, progressbar.OptionShowBytes(true))
	return p
}

// EnableANSICodes will use more optimized terminal i/o.
//
// Only useful in environments with support for ANSI escape sequences.
func (p *Options) EnableANSICodes() *Options {
	p.options = append(p.options, progressbar.OptionUseANSICodes(true))
	return p
}

// EnableIECUnits will enable IEC units (e.g. MiB) instead of the default
// SI units (e.g. MB).
func (p *Options) EnableIECUnits() *Options {
	p.options = append(p.options, progressbar.OptionUseIECUnits(true))
	return p
}

// EnableDescriptionAtLineEnd defines whether description should be written at line end instead of line start
func (p *Options) EnableDescriptionAtLineEnd() *Options {
	p.options = append(p.options, progressbar.OptionShowDescriptionAtLineEnd())
	return p
}

func (p *Options) MaxDetailRow(row int) *Options {
	p.options = append(p.options, progressbar.OptionSetMaxDetailRow(row))
	return p
}
