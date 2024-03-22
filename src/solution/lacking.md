## Array and Hashing
---
- Word Pattern (https://leetcode.com/problems/word-pattern/description/)
---
## two pointer
---
- Valid Palindrome II (https://leetcode.com/problems/valid-palindrome-ii/description/)
- Minimum Difference Between Highest and Lowest of K Scores (https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/)
- Merge Sorted Array (https://leetcode.com/problems/merge-sorted-array/description/)
    - When thinking about state (reduce extra state) optimization, ask whether if a state can be derived from other states. In this case, total length can be from m + n. Since either m decrement or n decrement assuming they work as the intended way, m + n - 1 is an valid index for nums1 []int.
    - Ask why I have to use this approach when I get stuck on problem. Break the current chain of thought and try a different way.
    - What's the easy way or less complex way to solve this problem as a human? This can be one way of finding the solution, and the solution is more intuitive. In this case, why shuffling everything in nums2 [] when one can start at the end of nums1 [], find the max. for each index, and overwrite nums1[]
- Remove Duplicate Array (https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/1188312362/)
    - Utilize the properties of the given input. This can help formulate a solution or reduce the amount of state and condition check since these properties provide some form of implication
---
## Sliding Window
---
- Minimum Difference Between Highest and Lowest of K Scores (https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/)
