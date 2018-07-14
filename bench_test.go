package bench

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"testing"

	intintmap "github.com/brentp/intintmap"
	christomic "github.com/chris-tomich/go-fast-hashmap"
	cornelk "github.com/cornelk/hashmap"
	lfmap "github.com/fastgeert/go-lfmap"
	suncat "github.com/suncat2000/hashmap"
)

func Benchmark_MathAbs_Positive(b *testing.B) {
	x := int64(20000000000)
	for i := 0; i < b.N; i++ {
		absed := int64(math.Abs(float64(x)))
		_ = absed
	}
}

func Benchmark_MathAbs_Negative(b *testing.B) {
	x := int64(-20000000000)
	for i := 0; i < b.N; i++ {
		absed := int64(math.Abs(float64(x)))
		_ = absed
	}
}

func Benchmark_Manual_Positive(b *testing.B) {
	x := int64(20000000000)
	for i := 0; i < b.N; i++ {
		var absed int64
		if x > 0 {
			absed = x
		} else {
			absed = x * -1
		}
		_ = absed
	}
}

func Benchmark_Manual_Negative(b *testing.B) {
	x := int64(-20000000000)
	for i := 0; i < b.N; i++ {
		var absed int64
		if x > 0 {
			absed = x
		} else {
			absed = x * -1
		}
		_ = absed
	}
}

func Benchmark_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		status := "open"
		if i%2 == 0 {
			status = "closed"
		}
		_ = status
	}
}

func Benchmark_If(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var status string
		if i%2 == 0 {
			status = "closed"
		} else {
			status = "open"
		}
		_ = status
	}
}

func allocate() {
	for allocs := 1; allocs <= 10; allocs++ {
		fmt.Sprint(allocs)
	}
}

func BenchmarkAllocateWithoutStopStartReset(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		allocate()
	}
}

func BenchmarkAllocateWithReset(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		allocate()
	}
	b.ResetTimer()
	for iter := 0; iter < b.N; iter++ {
		allocate()
	}
}

func BenchmarkAllocateStopStart(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		b.StopTimer()
		allocate()
		b.StartTimer()
		allocate()
	}
}

func BenchmarkSlice_Append(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		x := []int{}
		for i := 0; i < 10; i++ {
			x = append(x, i)
		}
	}
}

func BenchmarkSlice_MakeLength_Set(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		x := make([]int, 10)
		for i := 0; i < 10; i++ {
			x[i] = i
		}
	}
}

func BenchmarkSlice_MakeCapacity_Append(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		x := make([]int, 0, 10)
		for i := 0; i < 10; i++ {
			x = append(x, i)
		}
	}
}

func BenchmarkSliceDeclareVar(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		var s []int
		for i := 0; i < 3; i++ {
			s = append(s, i)
		}
	}
}

func BenchmarkSliceDeclareInitialize(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		s := []int{}
		for i := 0; i < 3; i++ {
			s = append(s, i)
		}
	}
}

func BenchmarkSliceDeclareUnderlying(b *testing.B) {
	var s []int
	for iter := 0; iter < b.N; iter++ {
		for i := 0; i < 3; i++ {
			s = append(s, i)
		}

		s = s[:]
	}
}

func getBool(i int) bool {
	return i%2 == 0
}

func BenchmarkVar_InLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			var value bool
			value = getBool(i)
			_ = value
		}
	}
}

func BenchmarkVar_OutLoop(b *testing.B) {
	var value bool
	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			value = getBool(i)
			_ = value
		}
	}
}

func BenchmarkVar_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			value := getBool(i)
			_ = value
		}
	}
}

func BenchmarkMaps_StdMap_Typed_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := map[int]int64{}
		for k := 0; k < 10000; k++ {
			table[k] = 1
		}
	}
}

func BenchmarkMaps_StdMap_WithLock_Typed_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := map[int]int64{}
		mutex := &sync.Mutex{}
		for k := 0; k < 10000; k++ {
			mutex.Lock()
			table[k] = 1
			mutex.Unlock()
		}
	}
}

func BenchmarkMaps_StdMap_Interface_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := map[interface{}]interface{}{}
		for k := 0; k < 10000; k++ {
			table[k] = 1
		}
	}
}

func BenchmarkMaps_SyncMap_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := sync.Map{}
		for k := 0; k < 10000; k++ {
			table.Store(k, 1)
		}
	}
}

func BenchmarkMaps_CornelkHashmap_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := cornelk.New(cornelk.DefaultSize)
		for k := 0; k < 10000; k++ {
			table.Set(k, 1)
		}
	}
}

func BenchmarkMaps_Intintmap_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := intintmap.New(100000, 1)
		var k int64
		for k = 0; k < 10000; k++ {
			table.Put(k, 1)
		}
	}
}

func BenchmarkMaps_SuncatHashmap_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := suncat.NewHashMap(16)
		var k int64
		for k = 0; k < 10000; k++ {
			table.Set(k, 1)
		}
	}
}

func BenchmarkMaps_ChrisTomichHashmap_Fill10K(b *testing.B) {
	keys := []string{}
	for k := 0; k < 10000; k++ {
		keys = append(keys, strconv.Itoa(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		table := christomic.New(10000)
		for k := 0; k < 10000; k++ {
			table.Set(keys[k], 1)
		}
	}
}
func BenchmarkMaps_LFMap_Fill10K(b *testing.B) {
	keys := []string{}
	for k := 0; k < 10000; k++ {
		keys = append(keys, strconv.Itoa(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		table := lfmap.NewLFmap()
		for k := 0; k < 10000; k++ {
			table.Set(keys[k], 1)
		}
	}
}
