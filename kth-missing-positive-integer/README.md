# Kth Missing Positive Integer

## Problem

Given an array of positive integers sorted in ascending order, find the Kth missing integer in the array.


## Personal Solution

- Create a counter that goes through each number one by one, and adds a count if the number is not on the list
- Create a hashmap where array items are the keys, and use the hashmap to check if the numbers in the array are in hashmap.

## Source

[Kth Positive Integer Leetcode](https://leetcode.com/problems/kth-missing-positive-number/description/)