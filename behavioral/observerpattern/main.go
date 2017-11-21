package main

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

type AbstractObserver interface {
	Update(abstractSubject AbstractSubject) error
}

type AbstractSubject interface {
	Attach(abstractObserver AbstractObserver)
	Detach(abstractObserver AbstractObserver)
	Notify()
	GetFav() string
}

type PatternObserver struct {
}

func (p *PatternObserver) Update(s AbstractSubject) error {
	spew.Dump(" new favorite patterns", s.GetFav())
	return nil
}

type PatternSubject struct {
	FavoritePatterns []string
	Observers        []AbstractObserver
	Fav              string
}

func (p *PatternSubject) Attach(observer AbstractObserver) {
	p.Observers = append(p.Observers, observer)
}

func (p *PatternSubject) Detach(observer AbstractObserver) {
	for k, v := range p.Observers {
		if reflect.DeepEqual(observer, v) {
			p.Observers = append(p.Observers[:k], p.Observers[k+1:]...)
			break
		}
	}
}

func (p *PatternSubject) Notify() {
	for _, v := range p.Observers {
		v.Update(p)
	}
}

func (p *PatternSubject) updateFav(fav string) {
	p.Fav = fav
	p.Notify()
}

func (p *PatternSubject) GetFav() string {
	return p.Fav
}

func main() {
	patternGossiper := &PatternSubject{}
	patternGossipFan := &PatternObserver{}
	patternGossiper.Attach(patternGossipFan)
	patternGossiper.updateFav("abstract factory, decorator, visitor")
	patternGossiper.updateFav("abstract factory, observer, decorator")
	patternGossiper.Detach(patternGossipFan)
	patternGossiper.updateFav("abstract factory, observer, paisley")
}
