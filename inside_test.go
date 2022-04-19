package batchlog

// var alphabet = []string{"Alain", "Bernard"} //, "Celeste", "Didier"} //, "Emile", "Francois", "Gaston", "Henri", "Isabelle", "Jean"}

// func TestChaosScenario(t *testing.T) {
// 	logger := NewLogger(OptTimeout(time.Hour))
// 	n := 100
// 	goLenMax := 20
// 	maxVariablesCount := 2
// 	milliSleep := 500 * time.Millisecond
// 	valueMaxCount := 2

// 	go readHeapsAlloc()

// 	for i := 0; i < n; i++ {
// 		goLen, _ := mathfunc.RandomInt(1, goLenMax)
// 		for j := 0; j < goLen; j++ {
// 			event := randLogLevel(&logger)
// 			variablesCount, _ := mathfunc.RandomInt(1, maxVariablesCount+1)
// 			go randEvent(event, variablesCount, valueMaxCount)
// 		}

// 		sleepTime, _ := mathfunc.Random0ToInt(int(milliSleep))
// 		time.Sleep(time.Duration(sleepTime))
// 	}

// 	for {
// 		time.Sleep(5 * time.Second)
// 	}
// }

// func randEvent(e *event, variableCount int, valueCount int) {
// 	for i := 0; i < variableCount; i++ {
// 		randVariables(e, valueCount)
// 	}

// 	randId, _ := mathfunc.Random0ToInt(len(alphabet))
// 	if v, _ := mathfunc.Random0ToInt(2); v == 0 {
// 		e.BatchMsg(alphabet[randId])
// 	}
// }

// func randLogLevel(l *Logger) *event {
// 	randNum, _ := mathfunc.Random0ToInt(3)
// 	switch randNum {
// 	case 0:
// 		return l.Debug()
// 	case 1:
// 		return l.Info()
// 	case 2:
// 		return l.Error()
// 	}
// 	return l.Debug()
// }

// func randVariables(e *event, maxNum int) {
// 	randNum, _ := mathfunc.Random0ToInt(maxNum)
// 	randId, _ := mathfunc.Random0ToInt(len(alphabet))

// 	e.BatchStr(alphabet[randId], strconv.Itoa(randNum))
// }

// func readHeapsAlloc() {
// 	memStats := runtime.MemStats{}
// 	delay := time.Second
// 	go deleteBatcher()
// 	for {
// 		runtime.ReadMemStats(&memStats)
// 		if isClean {
// 			x := 0
// 			x = x + 1
// 		} else {
// 			// fmt.Println(memStats.Alloc, memStats.Frees)
// 		}

// 		time.Sleep(delay)
// 	}
// }

// var isClean = false

// func deleteBatcher() {
// 	for {
// 		time.Sleep(time.Second)
// 		isClean = true
// 		for _, level := range batcher.root.nexts {
// 			if len(level.nexts) != 0 {
// 				isClean = false
// 				break
// 			}
// 		}

// 		if isClean {
// 			memStats := runtime.MemStats{}
// 			runtime.ReadMemStats(&memStats)
// 			fmt.Println(memStats.Alloc)
// 			fmt.Println("--------------------- CLEAN -------------------")
// 			runtime.GC()
// 			runtime.ReadMemStats(&memStats)
// 			fmt.Println(memStats.Alloc)
// 			break
// 		}
// 	}
// }
