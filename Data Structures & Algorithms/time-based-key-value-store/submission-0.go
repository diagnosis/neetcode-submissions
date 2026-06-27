


type Entry struct {
	timeStamp int
	value string
}

type TimeMap struct {
	tMap map[string][]Entry
}

func Constructor() TimeMap {
 return TimeMap{tMap: make(map[string][]Entry)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.tMap[key] = append(this.tMap[key], Entry{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries := this.tMap[key]
	start := 0
	end := len(entries) - 1
	result := ""
	for start <= end{
		mid := (start + end) / 2
		if entries[mid].timeStamp <= timestamp {
			result = entries[mid].value
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result

}


