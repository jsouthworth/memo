package memo

import (
	"reflect"
	"runtime"
	"unsafe"

	"jsouthworth.net/go/dyn"
	"jsouthworth.net/go/etm/atom"
	"jsouthworth.net/go/hash"
	"jsouthworth.net/go/immutable/hashmap"
)

type argList []interface{}

func (l argList) SeededHash(seed uintptr) uintptr {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&l))
	out := hash.Unsafe(unsafe.Pointer(hdr.Data), uintptr(hdr.Len), seed)
	runtime.KeepAlive(&l)
	return out
}

func (l argList) Equal(other interface{}) bool {
	ol, isArgList := other.(argList)
	return isArgList &&
		len(l) == len(ol) &&
		l.compareElements(ol)
}

func (l argList) compareElements(other argList) bool {
	for i, v := range l {
		if !dyn.Equal(v, other[i]) {
			return false
		}
	}
	return true
}

func Memoize(fn interface{}) func(args ...interface{}) interface{} {
	cache := atom.New(hashmap.Empty())
	return func(args ...interface{}) interface{} {
		out, inCache := cache.Deref().(*hashmap.Map).
			Find(argList(args))
		if inCache {
			return out
		}
		out = dyn.Apply(fn, args...)
		cache.Swap(func(m *hashmap.Map) *hashmap.Map {
			return m.Assoc(argList(args), out)
		})
		return out
	}
}
