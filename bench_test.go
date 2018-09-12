package main

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
	"github.com/vmihailenco/msgpack"
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

func BenchmarkMaps_Intintmap_Size1_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := intintmap.New(1, 0.6)
		var k int64
		for k = 0; k < 10000; k++ {
			table.Put(k, 1)
		}
	}
}

func BenchmarkMaps_Intintmap_Size1_WithLock_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := intintmap.New(1, 0.6)
		mutex := &sync.Mutex{}
		var k int64
		for k = 0; k < 10000; k++ {
			mutex.Lock()
			table.Put(k, 1)
			mutex.Unlock()
		}
	}
}

func BenchmarkMaps_Intintmap_Size10K_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := intintmap.New(10000, 0.6)
		var k int64
		for k = 0; k < 10000; k++ {
			table.Put(k, 1)
		}
	}
}

func BenchmarkMaps_Intintmap_Size10K_WithLock_Fill10K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := intintmap.New(10000, 0.6)
		mutex := &sync.Mutex{}
		var k int64
		for k = 0; k < 10000; k++ {
			mutex.Lock()
			table.Put(k, 1)
			mutex.Unlock()
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

func BenchmarkQueue_Chan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pipe := make(chan struct{})
		go func() {
			for k := 0; k < 10000; k++ {
				pipe <- struct{}{}
			}
		}()

		results := []struct{}{}
		for k := 0; k < 10000; k++ {
			results = append(results, <-pipe)
		}
	}
}

type (
	Side int8
)

const SideFoo Side = 10
const SideBar Side = 20

func getConstant(i int) Side {
	if i%2 == 0 {
		return SideFoo
	}

	return SideBar
}

func BenchmarkCondition_Switch_TwoCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foos := 0
		bars := 0
		for j := 0; j < 100; j++ {
			value := getConstant(j)
			switch value {
			case SideFoo:
				foos++
			case SideBar:
				bars++
			}
		}
	}
}

func BenchmarkCondition_Switch_CaseDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foos := 0
		bars := 0
		for j := 0; j < 100; j++ {
			value := getConstant(j)
			switch value {
			case SideFoo:
				foos++
			default:
				bars++
			}
		}
	}
}

func BenchmarkCondition_TypeSwitch_Native(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		ints := 0
		strings := 0
		for j := 0; j < 100; j++ {
			var value interface{}
			if j%2 == 0 {
				value = int(1)
			} else {
				value = string("1")
			}

			switch value := value.(type) {
			case float64:
			case float32:
			case bool:
			case complex64:
			case []byte:

			case int:
				result = value
				ints++
			case string:
				result = value
				strings++
			}
		}
	}

	_ = result
}

func BenchmarkCondition_TypeSwitch_Assisted(b *testing.B) {
	var result interface{}

	const kindInt = "int"
	const kindString = "string"
	for i := 0; i < b.N; i++ {
		ints := 0
		strings := 0
		for j := 0; j < 100; j++ {
			var value interface{}
			var kind string

			if j%2 == 0 {
				value = int(1)
				kind = kindInt
			} else {
				value = string("1")
				kind = kindString
			}

			switch kind {
			case "float64":
			case "float32":
			case "bool":
			case "complex64":
			case "[]byte":

			case kindInt:
				result = value.(int)
				ints++
			case kindString:
				result = value.(string)
				strings++
			}
		}
	}

	_ = result
}

func BenchmarkCondition_IfElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foos := 0
		bars := 0
		for j := 0; j < 100; j++ {
			value := getConstant(j)
			if value == SideFoo {
				foos++
			} else {
				bars++
			}
		}
	}
}

func BenchmarkCondition_IfElseIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foos := 0
		bars := 0
		for j := 0; j < 100; j++ {
			value := getConstant(j)
			if value == SideFoo {
				foos++
			} else if value == SideBar {
				bars++
			}
		}
	}
}

func BenchmarkSerializers_EncodeGotinyStruct(b *testing.B) {
	s := createStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeGotinyStruct(s)
		_ = res
	}
}

func BenchmarkSerializers_EncodeGobStruct(b *testing.B) {
	s := createStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeGobStruct(s)
		_ = res
	}
}

func BenchmarkSerializers_EncodeMsgpackStruct(b *testing.B) {
	s := createStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeMsgpackStruct(s)
		_ = res
	}
}

func BenchmarkSerializers_EncodeGotiny(b *testing.B) {
	m := createMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeGotiny(&m)
		_ = res
	}
}

func BenchmarkSerializers_EncodeMsgpack(b *testing.B) {
	m := createMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeMsgpack(m)
		_ = res
	}
}

func BenchmarkSerializers_EncodeGobMap(b *testing.B) {
	m := createMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeGob(m)
		_ = res
	}
}

func BenchmarkSerializers_EncodeJSONMap(b *testing.B) {
	m := createMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeJSON(m)
		_ = res
	}
}

func BenchmarkSerializers_DecodeGobMap(b *testing.B) {
	m := createMap(1000)
	byt := encodeGob(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[int64]float64
		decodeGob(byt, &result)
	}
}

func BenchmarkSerializers_DecodeJSONMap(b *testing.B) {
	m := createMap(1000)
	byt := encodeJSON(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[int64]float64
		decodeJSON(byt, &result)
	}
}

func BenchmarkSerializers_EncodeGobSliceMap(b *testing.B) {
	m := createSliceMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeGob(m)
		_ = res
	}
}

func BenchmarkSerializers_EncodeJSONSliceMap(b *testing.B) {
	m := createSliceMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeJSON(m)
		_ = res
	}
}

func BenchmarkSerializers_DecodeGobSliceMap(b *testing.B) {
	m := createSliceMap(1000)
	byt := encodeGob(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result []map[int64]float64
		decodeGob(byt, &result)
	}
}

func BenchmarkSerializers_DecodeJSONSliceMap(b *testing.B) {
	m := createSliceMap(1000)
	byt := encodeJSON(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result []map[int64]float64
		decodeJSON(byt, &result)
	}
}

func BenchmarkSerializers_DecodeGobStruct(b *testing.B) {
	m := AI(createStruct())
	byt := encodeGobStruct(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := decodeGobStruct(byt)
		_ = result
		//if result.GetName() != "blah" {
		//    panic(result)
		//}
	}
}

func BenchmarkSerializers_DecodeGotinyStruct(b *testing.B) {
	m := AI(createStruct())
	byt := encodeGotinyStruct(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := decodeGotinyStruct(byt)
		_ = result
		//if result.GetName() != "blah" {
		//    panic(result)
		//}
	}
}

func BenchmarkSerializers_DecodeMsgpackStruct(b *testing.B) {
	m := createStruct()
	byt := encodeMsgpackStruct(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := decodeMsgpackStruct(byt)
		_ = result
		//if result.GetName() != "blah" {
		//    panic(result)
		//}
	}
}

func BenchmarkSerializers_EncodeMsgpackStructConcrete(b *testing.B) {
	s := createStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(s)
	}
}

func BenchmarkSerializers_DecodeMsgpackStructConcrete(b *testing.B) {
	m := createStruct()
	byt, _ := msgpack.Marshal(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result interface{}
		msgpack.Unmarshal(byt, &result)
		//if result.(*A).GetName() != "blah" {
		//    panic(result)
		//}
	}
}
