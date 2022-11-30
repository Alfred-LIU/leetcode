package queue

/*
Design a hit counter which counts the number of hits received in the past 5 minutes (i.e., the past 300 seconds).

Your system should accept a timestamp parameter (in seconds granularity), and you may assume that calls are being made to the system in chronological order (i.e., timestamp is monotonically increasing). Several hits may arrive roughly at the same time.

Implement the HitCounter class:

HitCounter() Initializes the object of the hit counter system.
void hit(int timestamp) Records a hit that happened at timestamp (in seconds). Several hits may happen at the same timestamp.
int getHits(int timestamp) Returns the number of hits in the past 5 minutes from timestamp (i.e., the past 300 seconds).
*/

type Hit struct {
	Timestamp int
	Count     int
}

type HitCounter struct {
	HitsCount int
	HitsArr   []Hit
}

func Constructor() HitCounter {
	return HitCounter{0, []Hit{}}
}

func (this *HitCounter) Hit(timestamp int) {
	n := len(this.HitsArr)
	if n > 0 && this.HitsArr[n-1].Timestamp == timestamp {
		this.HitsArr[n-1].Count++
	} else {
		this.HitsArr = append(this.HitsArr, Hit{timestamp, 1})
	}
	this.HitsCount++
}

func (this *HitCounter) GetHits(timestamp int) int {
	for len(this.HitsArr) > 0 && (this.HitsArr[0].Timestamp <= timestamp-300) {
		this.HitsCount -= this.HitsArr[0].Count
		this.HitsArr = this.HitsArr[1:]
	}
	return this.HitsCount
}
