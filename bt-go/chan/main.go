package main


func main() {

	// var ch chan int= make( chan int ,2 )
	c1 := make(chan int , 10)
	c1<-10
	// ch<-10


}

//type hchan struct {
 383                                                                                                  | 34         qcount   uint           // total data in the queue
 384 // Puts the current goroutine into a waiting state and unlocks the lock.                         | 35         dataqsiz uint           // size of the circular queue
 385 // The goroutine can be made runnable again by calling goready(gp).                              | 36         buf      unsafe.Pointer // points to an array of dataqsiz elements
 386 func goparkunlock(lock *mutex, reason waitReason, traceEv byte, traceskip int) {                 | 37         elemsize uint16
 387         gopark(parkunlock_c, unsafe.Pointer(lock), reason, traceEv, traceskip)                   | 38         closed   uint32
 388 }                                                                                                | 39         elemtype *_type // element type
 389                                                                                                  | 40         sendx    uint   // send index
 390 func goready(gp *g, traceskip int) {                                                             | 41         recvx    uint   // receive index
 391         systemstack(func() {                                                                     | 42         recvq    waitq  // list of recv waiters
 392                 ready(gp, traceskip, true)                                                       | 43         sendq    waitq  // list of send waiters
 393         })                                                                                       | 44         
 394 }                                                                                                | 45         // lock protects all fields in hchan, as well as several
 395                                                                                                  | 46         // fields in sudogs blocked on this channel.
 396 //go:nosplit                                                                                     | 47         //
 397 func acquireSudog() *sudog {                                                                     | 48         // Do not change another G's status while holding this lock
 398         // Delicate dance: the semaphore implementation calls                                    | 49         // (in particular, do not ready a G), as this can deadlock
 399         // acquireSudog, acquireSudog calls new(sudog),                                          | 50         // with stack shrinking.
 400         // new calls malloc, malloc can call the garbage collector,                              | 51         lock mutex
 401         // and the garbage collector calls the semaphore implementation                          | 52 }

// 1. chansend , runtime.chansend()
//   1.1 判断


