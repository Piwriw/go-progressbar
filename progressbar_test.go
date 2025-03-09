package progressbar

import (
	"errors"
	"fmt"
	"log/slog"
	"testing"
	"time"
)

func TestAddSuffix(t *testing.T) {
	bar := NewProgressBar().
		Total(10).
		Tasks(NewProgressTask(TaskTime, 1)).
		Options(DefaultOptions())
	bar.Describe("Xxx")
	bar.Suffix("Suffix")
	bar.AutoRun()
}

func TestAddPrefix(t *testing.T) {
	var tasks []ProgressTask
	for i := 0; i < 10; i++ {
		tasks = append(tasks, NewProgressTask(TaskTime))
	}
	bar := NewProgressBar().
		Total(10).
		Tasks(tasks...).
		Options(DefaultOptions()).Create()
	bar.Prefix("prefix")
	bar.AutoRun()
}

func TestElapsedTime(t *testing.T) {
	// 启用 EnablePredictTime
	barWith := NewProgressBar().Total(5).Create().Options(ProgressOptions().ElapsedTime(false))

	t.Log("Progress with ElapsedTime disenabled:")
	for i := 0; i < 5; i++ {
		Task(i)
		barWith.Next()
	}

	// 禁用 EnablePredictTime
	barWithout := NewProgressBar().Total(5).Create().Options(ProgressOptions().ElapsedTime(true))

	t.Log("Progress with ElapsedTime enable:")
	for i := 0; i < 5; i++ {
		Task(i)
		barWithout.Next()
	}
}

func TestSpinnerCustom(t *testing.T) {
	// 启用 EnableShowBytes
	barWith := NewProgressBar().Total(-1).Create().Options(ProgressOptions().SpinnerCustom("a", "c", "b"))

	t.Log("Progress with SpinnerCustom enable:")
	for i := 0; i < 10; i++ {
		//Task(i)
		time.Sleep(120 * time.Millisecond)
		barWith.Add(1)
	}
}

//	func TestSpinnerType(t *testing.T) {
//		// 启用 EnableShowBytes
//		barWith := Add(-1,
//			ProgressOptions().SpinnerType(10),
//		)
//
//		t.Log("Progress with SpinnerType enable:")
//		for i := 0; i < 10; i++ {
//			//Task(i)
//			time.Sleep(120 * time.Millisecond)
//			barWith.Add(1)
//		}
//	}
//
//	func TestClearOnFinish(t *testing.T) {
//		// 启用 EnableShowBytes
//		barWith := Add(5,
//			ProgressOptions().ClearOnFinish(),
//		)
//
//		t.Log("Progress with ClearOnFinish enable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//		barWith.Finish()
//
//		// 禁用 EnableShowBytes
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableShowBytes disenable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//		barWithout.Finish()
//	}
//
//	func TestOptionShowBytes(t *testing.T) {
//		// 启用 EnableShowBytes
//		barWith := Add(5,
//			ProgressOptions().EnableShowBytes(),
//		)
//
//		t.Log("Progress with EnableShowBytes enable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnableShowBytes
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableShowBytes disenable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestEnableDescriptionAtLineEnd(t *testing.T) {
//		// 启用 EnableDescriptionAtLineEnd
//		barWith := Add(5,
//			ProgressOptions().EnableDescriptionAtLineEnd(),
//		)
//
//		t.Log("Progress with EnableDescriptionAtLineEnd enable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Describe("xxx")
//			barWith.Next()
//		}
//
//		// 禁用 EnablePredictTime
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableDescriptionAtLineEnd disenable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Describe("xxx")
//			barWithout.Next()
//		}
//	}
//
//	func TestEnableANSICodes(t *testing.T) {
//		// 启用 EnablePredictTime
//		barWith := Add(5,
//			ProgressOptions().EnableANSICodes(),
//		)
//
//		t.Log("Progress with EnableANSICodes enable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnablePredictTime
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableANSICodes disenable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestEnableIECUnits(t *testing.T) {
//		// 启用 EnablePredictTime
//		barWith := Add(5,
//			ProgressOptions().EnableIECUnits().DisShowTotalBytes(true),
//		)
//
//		t.Log("Progress with EnableIECUnits enable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnablePredictTime
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableIECUnits disenable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestFullWidth(t *testing.T) {
//		// 启用 EnablePredictTime
//		barWith := Add(5,
//			ProgressOptions().FullWidth(),
//		)
//
//		t.Log("Progress with FullWidth enable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnablePredictTime
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with FullWidth disenable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestWidth(t *testing.T) {
//		// 启用 EnablePredictTime
//		barWith := Add(5,
//			ProgressOptions().Width(10),
//		)
//
//		t.Log("Progress with Width enable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnablePredictTime
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with Width disenable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestEnablePredictTime(t *testing.T) {
//		// 启用 EnablePredictTime
//		barWith := NewProgressBar().Total(5).
//			Options(ProgressOptions().DisEnablePredictTime()).
//			AddBar()
//		barWith.AutoRun()
//		t.Log("Progress with DisEnablePredictTime disenabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnablePredictTime
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with DisEnablePredictTime enable:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestEOptionShowCount(t *testing.T) {
//		// 启用 EnableShowCount
//		barWith := Add(5,
//			ProgressOptions().EnableShowCount(),
//		)
//
//		t.Log("Progress with EnableShowCount enabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnableShowCount
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableShowCount disenabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestEnableElapsedTimeOnFinish(t *testing.T) {
//		// 启用 EnableVisibility
//		barWith := Add(5,
//			ProgressOptions().EnableElapsedTimeOnFinish(),
//		)
//
//		t.Log("Progress with EnableElapsedTimeOnFinish enabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnableColorCodes
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableElapsedTimeOnFinish disenabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestEnableVisibility(t *testing.T) {
//		// 启用 EnableVisibility
//		barWith := Add(5,
//			ProgressOptions().DisEnableVisibility(),
//		)
//
//		t.Log("Progress with EnableVisibility disenabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnableColorCodes
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableVisibility enabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestShowCount(t *testing.T) {
//		// 启用 EnableShowCount
//		barWith := Add(5,
//			ProgressOptions().EnableShowCount(),
//		)
//
//		t.Log("Progress with EnableShowCount enabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnableColorCodes
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableShowCount disabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
//
//	func TestEnableColorCodes(t *testing.T) {
//		// 启用 EnableColorCodes
//		barWith := Add(5,
//			ProgressOptions().EnableColorCodes().Theme(progressbar.Theme{
//				Saucer:        " ",
//				AltSaucerHead: "[yellow]<[reset]",
//				SaucerHead:    "[yellow]-[reset]",
//				SaucerPadding: "[white]•",
//				BarStart:      "[blue]|[reset]",
//				BarEnd:        "[blue]|[reset]",
//			}),
//		)
//
//		t.Log("Progress with EnableColorCodes enabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWith.Next()
//		}
//
//		// 禁用 EnableColorCodes
//		barWithout := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with EnableColorCodes disabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithout.Next()
//		}
//	}
func TestCompletion(t *testing.T) {
	// 启用 Completion
	barWithCompletion := NewProgressBar().Total(5).Create().Options(ProgressOptions().Completion(func() {
		fmt.Println("Progress completed!")
	}))

	t.Log("Progress with Completion enabled:")
	for i := 0; i < 5; i++ {
		Task(i)
		barWithCompletion.Next()
	}

	// 禁用 Completion
	bar := NewProgressBar().Total(5).Create().Options(ProgressOptions())

	t.Log("Progress with Completion disabled:")
	for i := 0; i < 5; i++ {
		Task(i)
		bar.Next()
	}
}

func TestShowIts(t *testing.T) {
	// 启用 ShowIts

	barWithShowIts := NewProgressBar().Total(5).Create().Options(ProgressOptions().EnableShowIts())

	t.Log("Progress with ShowIts enabled:")
	for i := 0; i < 5; i++ {
		Task(i)
		barWithShowIts.Next()
	}

	// 禁用 ShowIts
	bar := NewProgressBar().Total(5).Create().Options(ProgressOptions())

	t.Log("Progress with ShowIts disabled:")
	for i := 0; i < 5; i++ {
		Task(i)
		bar.Next()
	}
}

//	func TestShowTotalBytes(t *testing.T) {
//		// 启用 ShowTotalBytes
//		barWithTotalBytes := Add(10,
//			ProgressOptions().EnableShowIts().DisShowTotalBytes(true),
//		)
//
//		t.Log("Progress with ShowTotalBytes enabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithTotalBytes.Next()
//		}
//
//		// 禁用 ShowTotalBytes
//		barWithoutTotalBytes := Add(5,
//			ProgressOptions(),
//		)
//
//		t.Log("Progress with ShowTotalBytes disabled:")
//		for i := 0; i < 5; i++ {
//			Task(i)
//			barWithoutTotalBytes.Next()
//		}
//	}
//
//	func TestAutoRun(t *testing.T) {
//		tasks := make([]ProgressTask, 0)
//		for i := 0; i < 10; i++ {
//			tasks = append(tasks, NewProgressTask(TaskTime))
//		}
//		tasks = append(tasks, NewProgressTask(TaskTimeErr, 1))
//		tasks = append(tasks, NewProgressTask(Task, 1))
//		if err := AutoRun(ProgressOptions(), tasks...); err != nil {
//			t.Error(err)
//		}
//	}
//
//	func TestDescribe(t *testing.T) {
//		bar := Add(10,
//			ProgressOptions().
//				Writer(os.Stderr).
//				Width(10).
//				Throttle(65*time.Millisecond).
//				EnableShowCount().
//				EnableShowIts().
//				FullWidth().Completion(func() {
//				fmt.Fprint(os.Stderr, "\n")
//			}))
//		for i := 0; i < 10; i++ {
//			bar.Describe(fmt.Sprintf("Task %d", i))
//			Task(i)
//			bar.Next()
//		}
//	}
//
//	func TestAdd(t *testing.T) {
//		bar := Add(10,
//			ProgressOptions().
//				Writer(os.Stderr).
//				Width(10).
//				DisShowTotalBytes(true).
//				Throttle(65*time.Millisecond).
//				EnableShowCount().EnableShowIts().
//				SpinnerType(14).
//				FullWidth())
//		for i := 0; i < 10; i++ {
//			bar.Next()
//			Task(i)
//		}
//	}
func TaskTimeErr(num int) error {
	slog.Info("Task Done")
	time.Sleep(time.Duration(1) * time.Second)
	return errors.New("Task Validate Fail")
}

func TaskTime() error {
	slog.Info("Task Done")
	time.Sleep(time.Duration(1) * time.Second)
	return nil
}

func Task(num int) {
	slog.Info("Task Done")
	time.Sleep(time.Duration(num) * time.Second)
}
