package redlock

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMutex(t *testing.T) {
	value := rand.Intn(20)
	mutex := NewMutex("test", strconv.Itoa(value), 5)

	ok, err := mutex.Lock()
	assert.NoError(t, err)
	assert.Equal(t, true, ok)
	ret, err := mutex.Unlock()
	assert.NoError(t, err)
	assert.Equal(t, 1, ret)
}

func TestMutex2(t *testing.T) {
	value := rand.Intn(20)
	mutex := NewMutex("test", strconv.Itoa(value), 10)

	ok, err := mutex.Lock()
	assert.NoError(t, err)
	assert.Equal(t, true, ok)

	time.Sleep(2 * time.Second)
	notifyChan := make(chan struct{}, 1)
	go func() {
		defer close(notifyChan)

		ok, err := mutex.Lock()
		assert.NoError(t, err)
		assert.Equal(t, false, ok)
		notifyChan <- struct{}{}
	}()
	time.Sleep(2 * time.Second)

	<-notifyChan
	ret, err := mutex.Unlock()
	assert.NoError(t, err)
	assert.Equal(t, 1, ret)
}
