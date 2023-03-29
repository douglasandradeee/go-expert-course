package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// esse contexto leva 3s pra cancelar
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	// case o contexto tenha sido finalizado
	case <-ctx.Done():
		fmt.Println("Hotel booking canceled. Timeout reached.")
		return
	case <-time.After(2 * time.Second):
		fmt.Println("Hotel booked")
		return
	}
}
