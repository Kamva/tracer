package tracer_test

import (
	"fmt"

	"github.com/kamva/tracer"
)

func ExampleNew() {
	err := tracer.New("whoops")
	fmt.Println(err)

	// Output: whoops
}

func ExampleNew_printf() {
	err := tracer.New("whoops")
	fmt.Printf("%+v", err)

	// Example output:
	// whoops
	// github.com/kamva/tracer_test.ExampleNew_printf
	//         /home/dfc/src/github.com/kamva/tracer/example_test.go:17
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
	//         /github.com/kamva/tracer/_test/_testmain.go:106
	// runtime.main
	//         /home/dfc/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /home/dfc/go/src/runtime/asm_amd64.s:2059
}

func ExampleWithMessage() {
	cause := tracer.New("whoops")
	err := tracer.WithMessage(cause, "oh noes")
	fmt.Println(err)

	// Output: oh noes: whoops
}

func ExampleWithStack() {
	cause := tracer.New("whoops")
	err := tracer.WithStack(cause)
	fmt.Println(err)

	// Output: whoops
}

func ExampleWithStack_printf() {
	cause := tracer.New("whoops")
	err := tracer.WithStack(cause)
	fmt.Printf("%+v", err)

	// Example Output:
	// whoops
	// github.com/kamva/tracer_test.ExampleWithStack_printf
	//         /home/fabstu/go/src/github.com/kamva/tracer/example_test.go:55
	// testing.runExample
	//         /usr/lib/go/src/testing/example.go:114
	// testing.RunExamples
	//         /usr/lib/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /usr/lib/go/src/testing/testing.go:744
	// main.main
	//         github.com/kamva/tracer/_test/_testmain.go:106
	// runtime.main
	//         /usr/lib/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /usr/lib/go/src/runtime/asm_amd64.s:2086
	// github.com/kamva/tracer_test.ExampleWithStack_printf
	//         /home/fabstu/go/src/github.com/kamva/tracer/example_test.go:56
	// testing.runExample
	//         /usr/lib/go/src/testing/example.go:114
	// testing.RunExamples
	//         /usr/lib/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /usr/lib/go/src/testing/testing.go:744
	// main.main
	//         github.com/kamva/tracer/_test/_testmain.go:106
	// runtime.main
	//         /usr/lib/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /usr/lib/go/src/runtime/asm_amd64.s:2086
}

func ExampleWrap() {
	cause := tracer.New("whoops")
	err := tracer.Wrap(cause, "oh noes")
	fmt.Println(err)

	// Output: oh noes: whoops
}

func fn() error {
	e1 := tracer.New("error")
	e2 := tracer.Wrap(e1, "inner")
	e3 := tracer.Wrap(e2, "middle")
	return tracer.Wrap(e3, "outer")
}

func ExampleCause() {
	err := fn()
	fmt.Println(err)
	fmt.Println(tracer.Cause(err))

	// Output: outer: middle: inner: error
	// error
}

func ExampleWrap_extended() {
	err := fn()
	fmt.Printf("%+v\n", err)

	// Example output:
	// error
	// github.com/kamva/tracer_test.fn
	//         /home/dfc/src/github.com/kamva/tracer/example_test.go:47
	// github.com/kamva/tracer_test.ExampleCause_printf
	//         /home/dfc/src/github.com/kamva/tracer/example_test.go:63
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
	//         /github.com/kamva/tracer/_test/_testmain.go:104
	// runtime.main
	//         /home/dfc/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /home/dfc/go/src/runtime/asm_amd64.s:2059
	// github.com/kamva/tracer_test.fn
	// 	  /home/dfc/src/github.com/kamva/tracer/example_test.go:48: inner
	// github.com/kamva/tracer_test.fn
	//        /home/dfc/src/github.com/kamva/tracer/example_test.go:49: middle
	// github.com/kamva/tracer_test.fn
	//      /home/dfc/src/github.com/kamva/tracer/example_test.go:50: outer
}

func ExampleWrapf() {
	cause := tracer.New("whoops")
	err := tracer.Wrapf(cause, "oh noes #%d", 2)
	fmt.Println(err)

	// Output: oh noes #2: whoops
}

func ExampleErrorf_extended() {
	err := tracer.Errorf("whoops: %s", "foo")
	fmt.Printf("%+v", err)

	// Example output:
	// whoops: foo
	// github.com/kamva/tracer_test.ExampleErrorf
	//         /home/dfc/src/github.com/kamva/tracer/example_test.go:101
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
	//         /github.com/kamva/tracer/_test/_testmain.go:102
	// runtime.main
	//         /home/dfc/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /home/dfc/go/src/runtime/asm_amd64.s:2059
}

func Example_stackTrace() {
	type stackTracer interface {
		StackTrace() tracer.StackTrace
	}

	err, ok := tracer.Cause(fn()).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Printf("%+v", st[0:2]) // top two frames

	// Example output:
	// github.com/kamva/tracer_test.fn
	//	/home/dfc/src/github.com/kamva/tracer/example_test.go:47
	// github.com/kamva/tracer_test.Example_stackTrace
	//	/home/dfc/src/github.com/kamva/tracer/example_test.go:127
}

func ExampleCause_printf() {
	err := tracer.Wrap(func() error {
		return func() error {
			return tracer.New("hello world")
		}()
	}(), "failed")

	fmt.Printf("%v", err)

	// Output: failed: hello world
}
