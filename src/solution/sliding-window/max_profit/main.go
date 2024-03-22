package main

import "fmt"

func maxProfit(prices []int) int {
    max_profit := 0
    i := 0; var j int 
    for i < len(prices) && j < len(prices) {
        for j = i + 1; j < len(prices); j++ {
            profit := prices[j] - prices[i] 
            if profit <= 0 {
                i = j
                break
            } else {
                max_profit = max(profit, max_profit)
            }
        }
    }

    return max_profit
}

func main() {
    fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
