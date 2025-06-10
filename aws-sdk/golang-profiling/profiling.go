// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"runtime/pprof"
// )

// func main() {
// 	cpuFile, err := os.Create("cpu.prof")
// 	if err != nil {
// 		log.Fatalf("failed to create CPU profile: %v", err)
// 	}
// 	defer cpuFile.Close()

// 	if err = pprof.StartCPUProfile(cpuFile); err != nil {
// 		log.Fatalf("failed to start CPU profile: %v", err)
// 	}
// 	defer pprof.StopCPUProfile()

// 	for i := 0; i < 50; i++ {
// 		fmt.Println(fibonaci(i))
// 	}

// 	// Write memory profile
// 	memFile, err := os.Create("mem.prof")
// 	if err != nil {
// 		log.Fatalf("failed to create memory profile: %v", err)
// 	}
// 	defer memFile.Close()
// 	if err := pprof.WriteHeapProfile(memFile); err != nil {
// 		log.Fatalf("failed to write memory profile: %v", err)
// 	}
// }

// func fibonaci(n int) int {
// 	if n <= 1 {
// 		return n
// 	}

// 	return fibonaci(n-1) + fibonaci(n-2)
// }

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // Import for side effect - enables /debug/pprof endpoints
	"os"
	"reflect"
	"runtime"
	"runtime/pprof"
	"time"
)

// CPU-intensive function for demonstration
func cpuIntensiveWork() {
	for i := 0; i < 1000000; i++ {
		_ = fibonacci(30)
	}
}

// Recursive fibonacci - intentionally inefficient for demo
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Memory-intensive function for demonstration
func memoryIntensiveWork() {
	var data [][]int
	for i := 0; i < 1000; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(100)
		}
		data = append(data, slice)
	}
	// Simulate some work with the data
	sum := 0
	for _, slice := range data {
		for _, val := range slice {
			sum += val
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

// Example 1: Manual CPU Profiling
func manualCPUProfiling() {
	// Create CPU profile file
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("Could not create CPU profile:", err)
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Could not start CPU profile:", err)
	}
	defer pprof.StopCPUProfile()

	fmt.Println("Starting CPU-intensive work...")
	cpuIntensiveWork()
	fmt.Println("CPU-intensive work completed")
}

// Example 2: Manual Memory Profiling
func manualMemoryProfiling() {
	fmt.Println("Starting memory-intensive work...")
	memoryIntensiveWork()

	// Force garbage collection to get accurate memory stats
	runtime.GC()

	// Create memory profile file
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("Could not create memory profile:", err)
	}
	defer f.Close()

	// Write heap profile
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("Could not write memory profile:", err)
	}

	fmt.Println("Memory profile written to mem.prof")
}

// Example 3: HTTP Server with pprof endpoints
func startHTTPProfiler() {
	// This goroutine runs some background work
	go func() {
		for {
			cpuIntensiveWork()
			memoryIntensiveWork()
			time.Sleep(2 * time.Second)
		}
	}()

	// HTTP endpoints for profiling
	fmt.Println("Starting HTTP server with profiling endpoints...")
	fmt.Println("Available endpoints:")
	fmt.Println("  http://localhost:6060/debug/pprof/          - Index of available profiles")
	fmt.Println("  http://localhost:6060/debug/pprof/profile   - CPU profile (30s by default)")
	fmt.Println("  http://localhost:6060/debug/pprof/heap      - Heap memory profile")
	fmt.Println("  http://localhost:6060/debug/pprof/goroutine - Goroutine profile")
	fmt.Println("  http://localhost:6060/debug/pprof/trace     - Execution trace")

	log.Fatal(http.ListenAndServe(":6060", nil))
}

// Example 4: Custom profiling with runtime stats
func runtimeStats() {
	var m runtime.MemStats

	fmt.Println("=== Before intensive work ===")
	runtime.ReadMemStats(&m)
	printMemStats(m)

	// Do some work
	memoryIntensiveWork()

	fmt.Println("\n=== After intensive work ===")
	runtime.ReadMemStats(&m)
	printMemStats(m)

	// Force GC and check again
	runtime.GC()
	fmt.Println("\n=== After forced GC ===")
	runtime.ReadMemStats(&m)
	printMemStats(m)
}

func printMemStats(m runtime.MemStats) {
	fmt.Printf("Allocated memory: %d KB\n", m.Alloc/1024)
	fmt.Printf("Total allocations: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("System memory: %d KB\n", m.Sys/1024)
	fmt.Printf("Number of GC cycles: %d\n", m.NumGC)
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [cpu|mem|http|stats]")
		fmt.Println("  cpu   - Manual CPU profiling")
		fmt.Println("  mem   - Manual memory profiling")
		fmt.Println("  http  - Start HTTP server with pprof endpoints")
		fmt.Println("  stats - Show runtime memory statistics")
		return
	}

	switch os.Args[1] {
	case "cpu":
		manualCPUProfiling()
	case "mem":
		manualMemoryProfiling()
	case "http":
		startHTTPProfiler()
	case "stats":
		runtimeStats()
	default:
		fmt.Println("Unknown command. Use cpu, mem, http, or stats")
	}
}

/*
How to use this profiling example:

1. Manual CPU Profiling:
   go run main.go cpu
   go tool pprof cpu.prof
   (pprof) top10
   (pprof) web

2. Manual Memory Profiling:
   go run main.go mem
   go tool pprof mem.prof
   (pprof) top10
   (pprof) list main.memoryIntensiveWork

3. HTTP Profiling (run in separate terminal):
   go run main.go http

   Then in another terminal:
   go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
   go tool pprof http://localhost:6060/debug/pprof/heap

   Or view in browser:
   http://localhost:6060/debug/pprof/

4. Runtime Statistics:
   go run main.go stats

Common pprof commands:
- top: Show top functions by resource usage
- list <function>: Show source code with annotations
- web: Generate web-based visualization (requires graphviz)
- png: Generate PNG image
- pdf: Generate PDF report
- traces: Show execution traces
- help: Show all available commands

For web visualization, install graphviz:
- Ubuntu/Debian: sudo apt-get install graphviz
- macOS: brew install graphviz
- Windows: Download from https://graphviz.org/download/
*/

func MyCustom() {
	var a interface{}
	a = 10

	fmt.Println(a)
}

type CustomStruct struct {
	Field1 int
	Field  string
}

func Hello() {
	c := CustomStruct{
		Field1: 10,
		Field:  "Hello",
	}
	t := reflect.TypeOf(c)
	fmt.Println(t)
}
