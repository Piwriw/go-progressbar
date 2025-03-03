package progressbar

import (
	"github.com/schollz/progressbar/v3"
	"io"
	"time"
)

type Options struct {
	options []progressbar.Option
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
