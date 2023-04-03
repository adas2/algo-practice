package practice

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// LC # 1169

// TODO: fix this
func invalidTransactions(transactions []string) []string {
	// TODO: define custom map key: name; val: max heap of <timestamp, city> tuple
	tMap := make(map[string][]tcpair)
	// final result slice set
	res := []string{}
	res2 := make(map[string]struct{})
	// ans := []string{}

	var name, city string
	var time, amt int
	// first pass: add each transaction to map
	for i, transaction := range transactions {
		tx := strings.Split(transaction, ",")
		name, city = tx[0], tx[3]
		time, _ = strconv.Atoi(tx[1])
		amt, _ = strconv.Atoi(tx[2])

		if _, exists := tMap[name]; !exists {
			// create entry
			entry := tcpair{ts: time, city: city, index: i}
			entryList := []tcpair{entry}
			tMap[name] = entryList
			// fmt.Printf("Map: %s --> %v\n", name, entryList)
		} else {
			// add tcpair
			entry := tcpair{ts: time, city: city, index: i}
			tMap[name] = append(tMap[name], entry)
		}

	}

	// detect suspicious tx:
	for i, transaction := range transactions {
		fmt.Printf("Transaction %s\n", transaction)
		tx := strings.Split(transaction, ",")
		name, city = tx[0], tx[3]
		time, _ = strconv.Atoi(tx[1])
		amt, _ = strconv.Atoi(tx[2])
		// catch if invalid
		if amt > 1000 {
			res = append(res, transaction)
			// res[transaction] = struct{}{}
			continue
		}
		for _, tc := range tMap[name] {
			if tc.city != city && int(math.Abs(float64(time)-float64(tc.ts))) < 60 {
				// add both current and match
				// res = append(res, transaction)
				res2[transaction] = struct{}{}
				// TODO: do not add if already added due to high amt
				if amt, _ := strconv.Atoi(strings.Split(transactions[tc.index], ",")[2]); amt <= 1000 {
					res = append(res, transactions[tc.index])
					fmt.Printf("Added %s to res : %v\n", transactions[tc.index], res)
				}
				// res[transactions[j]] = struct{}{}
				fmt.Printf("Added %s to res2 : %v\n", transaction, res2)
			}
		}
		// add tcpair
		entry := tcpair{ts: time, city: city, index: i}
		tMap[name] = append(tMap[name], entry)

		fmt.Println("res list: ", res, "res2 set:", res2)

	}

	ans := res
	// for k := range res2 {
	// 	ans = append(ans, k)
	// }
	return ans
}

type tcpair struct {
	ts    int
	city  string
	index int
}

// Logic: pop the max timestamp such that city != curr_city,
// check if the (curr-ts> < 60, if not push back the contents popped
// if < 60 , add both curr and index which had conflict

// See python impl from LC forum:
/**

class Solution:
    def invalidTransactions(self, transactions: List[str]) -> List[str]:
        transaction_map = defaultdict(set)

        invalid = []
        for transaction in transactions:
            name, time, amount, city = transaction.split(",")
            time = int(time)
            transaction_map[(time, name)].add(city)


        for transaction in transactions:
            name, time, amount, city = transaction.split(",")
            time = int(time)
            amount = int(amount)

            if amount > 1000:
                invalid.append(transaction)
                continue

            # check to see if there is a time within 60 minutes of this transaction
            for time in range(time-60, time+61):
                if (time, name) not in transaction_map:
                    continue
                sus_transaction = transaction_map[(time, name)]

                # check to see if there is another transaction from a diff city if so its invalid
                if city in sus_transaction:
                    sus_transaction.remove(city)
                    if len(sus_transaction) >= 1:
                        invalid.append(transaction)
                        break
                    sus_transaction.add(city)
                else:
                    invalid.append(transaction)
                    break

        return invalid

**/
