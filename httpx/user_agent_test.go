package httpx

import (
	"fmt"
	"testing"
)

func TestRandomUserAgent(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomUserAgent())
	}

}
