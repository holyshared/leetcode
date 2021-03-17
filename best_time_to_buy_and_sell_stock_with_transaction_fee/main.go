package main

// prices = [1,3,2,8,4,9]
func maxProfit(prices []int, fee int) int {
	cash, hold := 0, -prices[0]

    for i := 1; i < len(prices); i++ {
    	soldCash := hold + prices[i] - fee
    	if cash < soldCash {
    		cash = soldCash
    	}
    	buyed := cash - prices[i]
    	if hold < buyed {
			hold = buyed
    	}
	}
	return cash
}
