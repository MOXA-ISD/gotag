package gotag

import (
    "os"
    "sync"
    "time"
    "math/rand"
)

func getEnv(key, alter string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return alter
}

func genId(n int) string {
    rand.Seed(time.Now().UnixNano())
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// ****************************************************************

type WaitLock struct {
    wg	        *sync.WaitGroup
    timeout     time.Duration
}

func WaitSync(wg *sync.WaitGroup, timeout time.Duration) *WaitLock {
    if wg == nil || timeout < 0 {
        return nil
    }
    self := WaitLock{}
    self.wg = wg
    self.timeout = timeout
    self.wg.Add(1)
    return &self
}

func (w *WaitLock) PostDelay() bool {
    c := make(chan struct{})
    go func() {
        defer close(c)
        w.wg.Wait()
    }()
    select {
    case <-c:
        return true
    case <-time.After(w.timeout * time.Second):
        return false
    }
}
