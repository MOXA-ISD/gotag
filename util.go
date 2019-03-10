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

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
    c := make(chan struct{})
    go func() {
        wg.Add(1)
        defer close(c)
        wg.Wait()
    }()
    select {
    case <-c:
        return false
    case <-time.After(timeout):
        return true
    }
}
