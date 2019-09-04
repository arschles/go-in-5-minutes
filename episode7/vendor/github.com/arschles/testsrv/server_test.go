package testsrv

import (
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/arschles/assert"
)

const (
	recvWaitTime = 10 * time.Millisecond
)

func TestNoMsgs(t *testing.T) {
	srv := StartServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer srv.Close()
	recvCh := make(chan []*ReceivedRequest)
	// ensure that it returns after waitTime
	go func() {
		recvCh <- srv.AcceptN(20, recvWaitTime/2)
	}()
	select {
	case recv := <-recvCh:
		assert.Equal(t, len(recv), 0, "number of received messages")
	case <-time.After(recvWaitTime):
		t.Errorf("AcceptN didn't return after [%+v]", recvWaitTime)
	}
}

func TestMsgs(t *testing.T) {
	const numSends = 20
	srv := StartServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	ch := make(chan []*ReceivedRequest)
	waitTime := 10 * time.Millisecond

	for i := 0; i < numSends; i++ {
		resp, err := http.Get(srv.URLStr())
		assert.NoErr(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "status code")
	}

	go func() {
		ch <- srv.AcceptN(numSends, waitTime)
	}()
	select {
	case r := <-ch:
		assert.Equal(t, len(r), numSends, "number of recevied messages")
	case <-time.After(recvWaitTime):
		t.Errorf("AcceptN didn't return after [%+v]", recvWaitTime)
	}
}

func TestConcurrentClose(t *testing.T) {
	srv := StartServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer srv.Close()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			srv.Close()
		}()
	}
	wg.Wait()
}

func TestConcurrentSend(t *testing.T) {
	const numSends = 100
	srv := StartServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	var wg sync.WaitGroup
	for i := 0; i < numSends; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(srv.URLStr())
			if err != nil {
				t.Errorf("error on GET [%s]", err)
				return
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected code [%d] got [%d]", http.StatusOK, resp.StatusCode)
				return
			}
		}()
	}
	wg.Wait()
	recv := srv.AcceptN(numSends, 10*time.Millisecond)
	assert.Equal(t, numSends, len(recv), "number of received requests")
}

func TestHandlerSleep(t *testing.T) {
	const numSends = 10
	ch := make(chan struct{})
	srv := StartServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.(http.Flusher).Flush()
		<-ch
	}))
	defer srv.Close()
	defer close(ch)

	var wg sync.WaitGroup
	for i := 0; i < numSends; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(srv.URLStr())
			if err != nil {
				t.Errorf("got error on GET when not expected [%s]", err)
				return
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("got status [%d] expected [%d]", resp.StatusCode, http.StatusOK)
				return
			}
		}()
	}
	wgChan := make(chan struct{})
	go func() {
		wg.Wait()
		close(wgChan)
	}()
	select {
	case <-wgChan:
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("not all senders completed within 100ms")
	}
	recv := srv.AcceptN(numSends, 10*time.Millisecond)
	assert.Equal(t, 0, len(recv), "num received messages")
}

func TestConcurrenctClose(t *testing.T) {
	srv := StartServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			srv.Close()
		}()
	}
	wg.Wait()
}
