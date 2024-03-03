# Number of unequal triplets

## Problem

You are given a 0-indexed array of positive integer nums. Find the number of triplets (i, j, k) that meet the following conditions:

0 <= i < j < k < nums.length
nums[i], nums[j], and nums[k] are pairwise distinct.
In other words, nums[i] != nums[j], nums[i] != nums[k], and nums[j] != nums[k].

Return the number of triplets that meet the conditions.

## Solution

- A brute force way to do this is to use 3 nested for loops that look at 3 different numbers at a time. This gives us an O(n^3) time complexity
- A more efficient way to do this is to create a hashmap of the numbers that are in the list and their frequencies. You can then multiply all the frequencies on the right, all the frequencies on the left, and all the frequenccies in the middle. This runs at an O(n) time complexity.

## Source

[Rearrange characters to make target string leetcode](https://leetcode.com/problems/rearrange-characters-to-make-target-string/)
