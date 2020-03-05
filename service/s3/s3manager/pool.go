package s3manager

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
)

type byteSlicePool interface {
	Get(aws.Context) (*[]byte, error)
	Put(*[]byte)
	ModifyCapacity(int)
	SliceSize() int64
	Close()
}

type maxSlicePool struct {
	// allocator is defined as a function pointer to allow
	// for test cases to instrument custom tracers when allocations
	// occur.
	allocator sliceAllocator

	slices      chan *[]byte
	allocations chan struct{}

	max       int
	sliceSize int64

	mtx sync.RWMutex
}

func newMaxSlicePool(sliceSize int64) *maxSlicePool {
	p := &maxSlicePool{sliceSize: sliceSize}
	p.allocator = func() *[]byte {
		bs := make([]byte, p.sliceSize)
		return &bs
	}

	return p
}

func (p *maxSlicePool) Get(ctx aws.Context) (*[]byte, error) {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	// check if context is canceled before attempting to get a slice
	// this ensures priority is given to the cancel case first
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if p.max == 0 {
		return nil, fmt.Errorf("get called on zero capacity pool")
	}

	select {
	case bs := <-p.slices:
		return bs, nil
	case _ = <-p.allocations:
		return p.allocator(), nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (p *maxSlicePool) Put(bs *[]byte) {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	if p.max == 0 {
		return
	}

	select {
	case p.slices <- bs:
	default:
	}
}

func (p *maxSlicePool) ModifyCapacity(delta int) {
	if delta == 0 {
		return
	}

	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.max += delta

	if p.max == 0 {
		p.empty()
		return
	}

	origAllocations := p.allocations
	p.allocations = make(chan struct{}, p.max)

	for i := 0; i < len(origAllocations)+delta; i++ {
		p.allocations <- struct{}{}
	}
	if origAllocations != nil {
		close(origAllocations)
	}

	origSlices := p.slices
	p.slices = make(chan *[]byte, p.max)
	if origSlices == nil {
		return
	}

	close(origSlices)
	for bs := range origSlices {
		select {
		case p.slices <- bs:
		default:
		}
	}
}

func (p *maxSlicePool) SliceSize() int64 {
	return p.sliceSize
}

func (p *maxSlicePool) Close() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.empty()
}

func (p *maxSlicePool) empty() {
	p.max = 0

	if p.allocations != nil {
		close(p.allocations)
		for range p.allocations {
			// drain channel
		}
		p.allocations = nil
	}

	if p.slices != nil {
		close(p.slices)
		for range p.slices {
			// drain channel
		}
		p.slices = nil
	}
}

type returnCapacityPoolCloser struct {
	byteSlicePool
	returnCapacity int
}

func (n *returnCapacityPoolCloser) ModifyCapacity(delta int) {
	if delta > 0 {
		n.returnCapacity = -1 * delta
	}
	n.byteSlicePool.ModifyCapacity(delta)
}

func (n *returnCapacityPoolCloser) Close() {
	if n.returnCapacity < 0 {
		n.byteSlicePool.ModifyCapacity(n.returnCapacity)
	}
}

type sliceAllocator func() *[]byte

var newByteSlicePool = func(sliceSize int64) byteSlicePool {
	return newMaxSlicePool(sliceSize)
}
