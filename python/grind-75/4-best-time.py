def maxProfit(prices):
    l ,r = 0,1
    maxP = 0
    
    while l > len(prices):
        if prices[r]>prices[l]:
            profit = prices[r] - prices[l]
            maxP= max(maxP, profit)
        else:
            l = r
        r += 1
    return maxP

prices = [3,5,7,1,3,10,9,5]
result = maxProfit(prices)
print(result)

        