package s3manager

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestMaxSlicePool(t *testing.T) {
	pool := newMaxSlicePool(1)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// increase pool capacity by 2
			pool.ModifyCapacity(2)

			// remove 2 items
			bsOne := pool.Get()
			bsTwo := pool.Get()

			// attempt to remove a 3rd
			unblocked := make(chan struct{})
			go func() {
				defer close(unblocked)
				bs := pool.Get()
				pool.Put(bs)
			}()

			pool.Put(bsOne)

			<-unblocked

			pool.ModifyCapacity(-1)

			pool.Put(bsTwo)

			pool.ModifyCapacity(-1)

			rando := make([]byte, 1)
			pool.Put(&rando)
		}()
	}
	wg.Wait()

	pool.Empty()
}

type recordedPartPool struct {
	recordedAllocs      uint64
	recordedGets        uint64
	recordedOutstanding int64
	*maxSlicePool
}

func newRecordedPartPool(sliceSize int64) *recordedPartPool {
	sp := newMaxSlicePool(sliceSize)

	rp := &recordedPartPool{}

	allocator := sp.allocator
	sp.allocator = func() *[]byte {
		atomic.AddUint64(&rp.recordedAllocs, 1)
		return allocator()
	}

	rp.maxSlicePool = sp

	return rp
}

func (r *recordedPartPool) Get() *[]byte {
	atomic.AddUint64(&r.recordedGets, 1)
	atomic.AddInt64(&r.recordedOutstanding, 1)
	return r.maxSlicePool.Get()
}

func (r *recordedPartPool) Put(b *[]byte) {
	atomic.AddInt64(&r.recordedOutstanding, -1)
	r.maxSlicePool.Put(b)
}

func swapByteSlicePool(f func(sliceSize int64) byteSlicePool) func() {
	orig := newByteSlicePool

	newByteSlicePool = f

	return func() {
		newByteSlicePool = orig
	}
}
