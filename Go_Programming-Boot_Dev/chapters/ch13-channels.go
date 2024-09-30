package main

import (
	"fmt"
	"time"
)

/* func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)
	go func() {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()
	isOld := <-isOldChan
	fmt.Println("email 1 is old: ", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old: ", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old: ", isOld)
} */

/* func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
} */

/* func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
} */

/* func countReports(numSentCh chan int) int {
	countedReps := 0
	for {
		_, ok := <-numSentCh
		if ok {
			countedReps++
		} else {
			break
		}
	}
	return countedReps
} */

/* func concurrrentFib(n int) {
	chInts := make(chan int)
	go fibonacci(n, chInts)
	for item := range chInts {
		fmt.Println(item)
	}
}
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
} */

/* func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}
} */

/* func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(time.Millisecond * 500)
		}
	}
}
func takeSnapshot() { fmt.Println("Taking a snapshot...") }
func saveSnapshot() { fmt.Println("Saving a snapshot...") }
func waitForData()  { fmt.Println("Waiting for data.") } */

func main() {
	// concurrrentFib(13)
}
