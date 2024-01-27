package canplaceflowers

import "fmt"

func CanPlaceFlowers(flowerbed []int, n int) bool {
    l := len(flowerbed)
    if n == 0 || (l == 1 && flowerbed[0] == 0 && n == 1) {
        return true
    }
    limit := l / 2
    if l % 2 != 0 {
        limit += 1
    }
    if n > limit {
        return false
    }

    for i := 0; i < l && n > 0; i++ {
        if flowerbed[i] == 0 {
            if i > 0 && i + 1 < l {
                if flowerbed[i - 1] == 0 && flowerbed[i + 1] == 0 {
                    n -= 1
                    i += 1
                }
           } else if i == 0 { // at the left edge
               if flowerbed[i + 1] == 0 {
                   n -= 1
                   i += 1
               }
           } else { // at the right edge
               if flowerbed[i - 1] == 0 {
                   n -= 1
                   i += 1
               }
           }
        }
    }

    return n == 0
}

/*
 if encounter nested if-else that cannot simply by unfolding and flattening 
 the if-else, try variable assignment but before that, pen-and-paper and 
 boolean logic! I seems to afraid to doing this and straight into write code
*/
func CanPlacePowerBetter(flowerbed []int, n int) bool {
    l := len(flowerbed)
    if n == 0 || (l == 1 && flowerbed[0] == 0 && n == 1){ return true }
    limit := l / 2
    if l % 2 != 0 {
        limit += 1
    }
    if n > limit { return false }
    for i := 0; i < l && n > 0; i++ {
        if flowerbed[i] == 1 {
            continue
        }
        
        // OR fall-off
        prev_zero := i == 0 || flowerbed[i - 1] == 0
        next_zero := i == l - 1 || flowerbed[i + 1] == 0
        if prev_zero && next_zero {
            n -= 1
            flowerbed[i] = 1
        }
    }

    fmt.Println(flowerbed, n)

    return n == 0
}
