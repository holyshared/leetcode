package main

type Station struct {
	Name string
	At   int
}

type Result struct {
	Start *Station
	End   *Station
}

func (this *Result) SectionKey() string {
	return this.Start.Name + "-" + this.End.Name
}

func (this *Result) Time() int {
	return this.End.At - this.Start.At
}

type UndergroundSystem struct {
	cards          map[int]*Result
	sectionResults map[string][]*Result
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		cards:          map[int]*Result{},
		sectionResults: map[string][]*Result{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	_, ok := this.cards[id]
	if ok {
		return
	}
	this.cards[id] = &Result{Start: &Station{Name: stationName, At: t}}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	card, ok := this.cards[id]
	if !ok {
		return
	}
	card.End = &Station{Name: stationName, At: t}
	delete(this.cards, id)

	results, hasResults := this.sectionResults[card.SectionKey()]
	if hasResults {
		results = append(results, card)
	} else {
		results = []*Result{card}
	}
	this.sectionResults[card.SectionKey()] = results
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	results, ok := this.sectionResults[startStation+"-"+endStation]
	if ok {
		total := 0
		for i := 0; i < len(results); i++ {
			total += results[i].Time()
		}
		return float64(total) / float64(len(results))
	} else {
		return 0
	}
}
