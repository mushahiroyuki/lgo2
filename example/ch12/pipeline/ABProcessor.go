package main

import "context"

type abProcessor struct {  //liststart1
	outA chan aOut
	outB chan bOut
	errs chan error
}

func newABProcessor() *abProcessor { // ファクトリ関数
	return &abProcessor{
		outA: make(chan aOut, 1),
		outB: make(chan bOut, 1),
		errs: make(chan error, 2),
	}
}  //listend1

// abProcessorのメソッド start  //liststart2
func (p *abProcessor) start(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- aOut
	}()

	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()

}  //listend2

// abProcessorのメソッド wait  //liststart3
func (p *abProcessor) wait(ctx context.Context) (cIn, error) {
	var cData cIn
	for count := 0; count < 2; count++ {
		select {
		case a := <-p.outA:
			cData.a = a
		case b := <-p.outB:
			cData.b = b
		case err := <-p.errs:
			return cIn{}, err
		case <-ctx.Done():
			return cIn{}, ctx.Err()
		}
	}
	return cData, nil
}  //listend3

type aOut struct {
}

type bOut struct {
}

type cIn struct {
	a aOut
	b bOut
}

func getResultA(ctx context.Context, in string) (aOut, error) {
	return aOut{}, nil
}

func getResultB(ctx context.Context, in string) (bOut, error) {
	return bOut{}, nil
}
