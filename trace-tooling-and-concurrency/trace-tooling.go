// manually adding tracing to your program

import (
    "runtime/trace"
    "os"
)

func main() {
    f, _ := os.Create("trace.out")
    defer f.Close()
    trace.Start(f)
    defer trace.Stop()

    // Your concurrent code here
}
