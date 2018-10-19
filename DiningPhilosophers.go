package DiningPhilosophers

import (
       "fmt"
       "sync"
)

// Intializing Variables to be used later
var chopsticks [5]int
var pickup sync.Mutex
var exit int

// Funtion that checks if the chopsticks are available 
func checksChopsticks(pick bool, pid int) bool {
       var i, j = pid, (pid+1)%5
       var has_value int = 1
       if pick {
              has_value = 0
       }
       if chopsticks[i]==has_value && chopstics[j]==has_value {
              chopsticks[i] = (has_value+1)%2
              chopsticks[j] = (has_value+1)%2
              return true
       }
       return false
}

// Funtion that acts on if chopsticks are available to be used

func Action(pick bool, pid int) {
       for true {
              // Global var: pickup -> sync.Mutex
              pickup.Lock()
              if checkChopsticks(pick, pid) {
                     pickup.Unlock()
                     break
              }
              pickup.Unlock()
       }
}
