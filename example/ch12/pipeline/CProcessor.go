package main

import "context"

type cProcessor struct {  //liststart1
	outC chan COut
	errs chan error
}

//ファクトリ関数
func newCProcessor() *cProcessor {
	return &cProcessor{
		outC: make(chan COut, 1),
		errs: make(chan error, 1),
	}
}

// cProcessorのメソッドstart
func (p *cProcessor) start(ctx context.Context, inputC cIn) {
	go func() {
		cOut, err := getResultC(ctx, inputC)
		if err != nil {
			p.errs <- err
			return
		}
		p.outC <- cOut
	}()
}

func getResultC(ctx context.Context, c cIn) (COut, error) {
	return COut{}, nil
}

// cProcessorのメソッドwait
func (p *cProcessor) wait(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}  //listend1
