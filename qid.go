package qid

import (
	"context"
	"fmt"
	"time"
)

func Sleep() {
	ctxto, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	tk := time.NewTicker(time.Second)
	count := 0
	for {
		select {
		case <-tk.C:
			fmt.Printf("%d seconds passed\n", count+1)
			count++
		case <-ctxto.Done():
			tk.Stop()
			return
		}
	}
}
