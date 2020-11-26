package catserver

import (
	"fmt"
	"strings"
	"sync"
)

func checkServerStatus(ch chan string) {
	dataCh := make(chan string)
	go status(dataCh)
	res := <-dataCh
	ch <- res

}
func fetchBannerContent(ch chan map[string]string) {
	copyCh := make(chan map[string]string)
	go banner(copyCh)
	res := <-copyCh
	res["copyrightYear"] = "2020"
	ch <- res
}

func fetchAllOwners(ch chan []string) {
	OwnerCh := make(chan []string)
	go getOwners(OwnerCh)
	owners := <-OwnerCh
	for i, o := range owners {
		owners[i] = strings.ToLower(o)
	}
	ch <- owners
}

func fetchCatsByOwner(owner string, ch chan map[string][]string, errCh chan error) {
	dCh := make(chan litter)
	eCh := make(chan error)
	go getCatsByOwner(owner, dCh, eCh)

	select {
	case cats := <-dCh:
		ch <- map[string][]string{
			owner: cats,
		}
	case err := <-eCh:
		errCh <- err
	}
}

func fetchCatPics(cats []string) {
	dCh := make(chan string)
	eCh := make(chan error)
	wg := sync.WaitGroup{}
	catPics := []string{}
	for _, cat := range cats {
		wg.Add(1)
		go getPic(cat, dCh, eCh, &wg)
		select {
		case cat := <-dCh:
			catPics = append(catPics, cat)
		case <-eCh:
			catPics = append(catPics, "placeholder.jpg")
		}
	}
	wg.Wait()
}
func fetchOwnersWithCats() {
	oCh := make(chan []string)
	cCh := make(chan map[string][]string, 10)
	eCh := make(chan error)
	wg := sync.WaitGroup{}
	go fetchAllOwners(oCh)
	owners := <-oCh
	for _, owner := range owners {
		wg.Add(1)
		go func(owner string) {
			fetchCatsByOwner(owner, cCh, eCh)
			wg.Done()
		}(owner)
	}
	wg.Wait()
	close(cCh)
	var petsWithCats []map[string][]string
	for cats := range cCh {
		petsWithCats = append(petsWithCats, cats)
	}
	fmt.Println(petsWithCats)
}

// function kickLegacyServerUntilItWorks() {}

// function buySingleOutfit() {}
