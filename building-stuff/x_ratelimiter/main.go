package main

import (
	"fmt"
	"os"
)

/*
Token bucket ratelimiter TODO
- token rate /sec
- bucket size fixed
- incoming request size
*/
// const tokenRatePerSec int = 3
const bucketSize int = 5

func main() {
	var bucketStatus int = 0
	var timePassed int = 0
	var requestSize int // get how many request will be made
	for {               // infinite loop
		fmt.Printf("At time : %d, the bucket has %d tokens\n", timePassed, bucketStatus)
		fmt.Print("Enter your request size: ")
		fmt.Scanln(&requestSize)

		if requestSize >= 0 {
			if !isBucketFilled(bucketStatus) {
				bucketStatus = addTokens(bucketStatus, requestSize)
				timePassed += 1
				fmt.Printf("At time : %d, Bucket has %d tokens\n", timePassed, bucketStatus)
			}
		} else {
			fmt.Printf("At time : %d, Bucket has %d tokens\n", timePassed, bucketStatus)
			os.Exit(0) // exit the applciation, end of req
		}
	}
}

// if the bucket capacity has reached
func isBucketFilled(status int) bool {
	if status == bucketSize {
		return true
	} else {
		return false
	}
}

// add tokens as per the capacity allowed
func addTokens(status, requestCount int) int {
	if status == bucketSize {
		return bucketSize
	} else {
		if requestCount+status < bucketSize {
			return requestCount + status
		} else {
			return bucketSize - status
		}
	}
}
