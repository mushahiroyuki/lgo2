package main

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

func main() {
	calcWithTimeout(1)
	calcWithTimeout(2)
	calcWithTimeout(5)
	calcWithTimeout(10)
}

func calcWithTimeout(numSeconds time.Duration) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), numSeconds*time.Second)
	defer cancelFunc()
	start := time.Now()
	result, err := calcPi(ctx)
	calcTime := time.Since(start)
	fmt.Println(result)
	fmt.Println(calcTime)
	fmt.Println(err)
}

func calcPi(ctx context.Context) (string, error) {
	var sum big.Float
	sum.SetInt64(0)
	var d big.Float
	d.SetInt64(1)
	two := big.NewFloat(2)
	i := 0  //liststart1
	for {
		if err := context.Cause(ctx); err != nil {
			fmt.Println("cancelled after", i, "iterations")
			return sum.Text('g', 100), err
		}
		var diff big.Float
		diff.SetInt64(4)
		diff.Quo(&diff, &d)
		if i%2 == 0 {
			sum.Add(&sum, &diff)
		} else {
			sum.Sub(&sum, &diff)
		}
		d.Add(&d, two)
		i++
	}  //listend1
}
