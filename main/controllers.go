package catserver

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// func checkLegacyStatus(errors) string {

//   s := rand.NewSource(time.Now().UnixNano())
//   r := rand.New(s)
//   if r.Float64() < 0.35 {
//   //  return error
//   } else{
//    // return status string
//   }
//   // else return '200 - the legacy server is up';
// };

func status(ch chan string) {
	time.Sleep(time.Millisecond * 2)
	ch <- "200 - the server is good"
}

func banner(ch chan map[string]string) {
	time.Sleep(time.Millisecond * 2)
	ch <- bannerContent
}

func getOwners(ch chan []string) {
	time.Sleep(time.Millisecond * 2)
	ch <- ownerContent
}

func getCatsByOwner(o string, dCh chan litter, eCh chan error) {
	time.Sleep(time.Millisecond * 2)
	if cats, ok := catsByOwner[o]; ok {
		dCh <- cats
	} else {
		eCh <- errors.New("Owner not found")
	}
}

func getPic(pic string, dCh chan string, eCh chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 1)
	if strings.Contains(pic, "cat") {
		dCh <- fmt.Sprintf("%s.jpg", pic)
	} else {
		eCh <- fmt.Errorf("%q file not found", pic)
	}

}

// exports.fetchCatsByOwner = (errors, db, owner) => {
//   const cats = db.catsByOwner[owner];
//   if (cats) return cats;
//   else errors.push(`404 - ${owner} not found`);
// };

// exports.buyBuyBuy = (errors, db, outfit, handleResponse) => {
//   const cost = db.outfits[outfit];
//   const checkout = { quantity: 0, outfit, totalCost: 0 };
//   if (!cost) {
//     errors.push(`404 - ${outfit} not found`);
//     return;
//   } else {
//     for (let i = 0; i < Math.floor(Math.random() * 1000); i++) {
//       checkout.totalCost += cost;
//       checkout.quantity++;
//       handleResponse(null, checkout);
//     }
//     return checkout;
//   }
// };
