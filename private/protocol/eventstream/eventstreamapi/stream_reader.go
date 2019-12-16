package eventstreamapi

//type StreamReader struct {
//	eventReader *EventReader
//	stream      chan interface{}
//	errVal      atomic.Value
//
//	done      chan struct{}
//	closeOnce sync.Once
//
//	//	publishEvent func(Unmarshaler
//}
//
//func NewStreamReader(eventReader *EventReader,
//	logger aws.Logger, logLevel aws.LogLevelType,
//) *StreamReader {
//	r := &StreamReader{
//		stream: make(chan interface{}),
//		done:   make(chan struct{}),
//	}
//	r.eventReader.UseLogger(logger, logLevel)
//
//	return r
//}
//
//// Close will close the underlying event stream reader.
//func (r *StreamReader) Close() error {
//	r.closeOnce.Do(r.safeClose)
//
//	return r.Err()
//}
//
//func (r *StreamReader) safeClose() {
//	close(r.done)
//}
//
//func (r *StreamReader) Err() error {
//	if v := r.errVal.Load(); v != nil {
//		return v.(error)
//	}
//
//	return nil
//}
//
//func (r *StreamReader) Events() <-chan Unmarshaler {
//	return r.stream
//}
//
//func (r *StreamReader) readEventStream() {
//	defer close(r.stream)
//
//	for {
//		event, err := r.eventReader.ReadEvent()
//		if err != nil {
//			if err == io.EOF {
//				return
//			}
//			select {
//			case <-r.done:
//				// If closed already ignore the error
//				return
//			default:
//			}
//			r.errVal.Store(err)
//			return
//		}
//
//		select {
//		case r.stream <- event.(Unmarshaler):
//		case <-r.done:
//			return
//		}
//	}
//}
