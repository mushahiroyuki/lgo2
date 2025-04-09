package main

import (
	"fmt"
	"sync"
)

type MutexScoreboardManager struct {  //liststart1
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}  //listend1

func main() {
	msm := NewMutexScoreboardManager()
	teams := []string{"Lions", "Tigers", "Bears"}
	var wg sync.WaitGroup
	wg.Add(len(teams))
	for _, v := range teams {  // 各チームに関して
		go func(team string) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				curScore, found := msm.Read(team)
				if !found {
					curScore = 10  // 最初のスコアは10
				} else {
					curScore += len(team)  // チーム名の文字列だけ追加
				}
				msm.Update(team, curScore)
				// fmt.Println(team, curScore)
			}
		}(v)
	}
	wg.Wait()
	for _, v := range teams {
		score, found := msm.Read(v)
		fmt.Println(v, score, found)
	}
}
