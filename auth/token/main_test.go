package token

import (
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M){
    result := m.Run()
    os.Exit(result)
}

func randStr (length int) string {
    alphnum := "abcdefghijklmnopqrstuvwxyz"
    k := len(alphnum)
    rand_str := make([]byte, length)
    for i := 0; i < length; i++ {
	rand_str[i] = alphnum[rand.Intn(k)]
    }
    return string(rand_str)
}
