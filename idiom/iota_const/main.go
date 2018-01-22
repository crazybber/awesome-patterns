package main

import "github.com/davecgh/go-spew/spew"

type FeedLockLinkType int

const (
	Racing FeedLockLinkType = iota
	Market
)

var feedLockLinkTypes = [...]string{
	"racing.Races.raceID",
	"racing.Markets.marketID",
}

func (f FeedLockLinkType) String() string { return feedLockLinkTypes[f] }

func main() {
	test := FeedLockLinkType(Racing).String()
	spew.Dump(test)
}
