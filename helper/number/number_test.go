package number

import (
	"github.com/emacsist/go-common/helper/time"
)

import "testing"

const n = 11

func TestHeapSort(t *testing.T) {
	data := GenerateInt(n, n)
	// t.Logf("before %v\n", data)
	start := time.CurrentMillis()
	HeapSort(data)
	t.Logf("heap sort cost %v ms\n", time.CurrentMillis()-start)
	// t.Logf("after %v\n", data)
}

func TestInsertSort(t *testing.T) {
	data := GenerateInt(n, n)
	// t.Logf("before %v\n", data)
	start := time.CurrentMillis()
	InsertSort(data)
	t.Logf("insert sort cost %v ms\n", time.CurrentMillis()-start)
	// t.Logf("after %v\n", data)
}

func TestMergeSort(t *testing.T) {
	data := GenerateInt(n, n)
	// t.Logf("before %v\n", data)
	start := time.CurrentMillis()
	data = MergeSort(data)
	t.Logf("merge sort cost %v ms\n", time.CurrentMillis()-start)
	// t.Logf("after %v\n", data)
}

func TestQuickSort(t *testing.T) {
	data := GenerateInt(n, n)
	// t.Logf("before %v\n", data)
	start := time.CurrentMillis()
	QuickSort(data)
	t.Logf("quick sort cost %v ms\n", time.CurrentMillis()-start)
	// t.Logf("after %v\n", data)
}

func TestSelectSort(t *testing.T) {
	data := GenerateInt(n, n)
	// t.Logf("before %v\n", data)
	start := time.CurrentMillis()
	SelectSort(data)
	t.Logf("select sort cost %v ms\n", time.CurrentMillis()-start)
	// t.Logf("after %v\n", data)
}

func TestShellSort(t *testing.T) {
	data := GenerateInt(n, n)
	// t.Logf("before %v\n", data)
	start := time.CurrentMillis()
	ShellSort(data)
	t.Logf("shell sort cost %v ms\n", time.CurrentMillis()-start)
	// t.Logf("after %v\n", data)
}
