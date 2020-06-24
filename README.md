# Problem Solving Practice

## Solved:
1. Alphapermutation: Find all permutations of a string
2. bitcount: Count the number of bits in a number
3. bytelandian problem:
4. DFS, BFS, TopSort: implementation of graphs
5. Graph problems:
6. integer_division: Division without using divide operation
7. GCD of two numbers
8. hammingwt: Calculate hamming weight
9. Heapsort of a vector
10. Intbr: Integer break problem leetcode
11. Memread: Read chunk of memory from disk using Read4kBlock function
12. Mergesort implementation
13. pathwobstacles:  Find the unique path between two points given an obstacle matrix
14. Permutation: given number of digits n and integer k find the k-th permutation of string 123...n in lexicographic order
15. power:  given base and exponent calculate power of a number
16. power_four: Find if a number is a power of four
17. power_three: Find is a number is a power of three
18. Quicksort implementation
19. reversesentence: reverse the words of a sentence
20. threadpool: threadpool implementation
21. tmplt: C++ 11 template use cases
22. toggle: Toggle the alternate bulb problem (leetcode)
23. trailing_zero_in_factorial: Given the number find the tralling zeros in its factorial
24. wiggle:  Length of wiggle sequence in a vector
25. ipcheck: Check whether given string is a valid IP address
26. midsum: Find the index of an array where the sum of the left part equals the sum of the right part( const space, linear time)
27. twosum: Given an array and a target value find two elements in the array whose sum = target
28. threesum: Given an array and target find 3 numbers in array whose sum = target 
29. array_rotate: Rotate array of length N by K elements 
30. pointdist: Given N points in a 2D plain, find k points closest to the origin (0,0). Note: k << N find optimal solution in O(Nlogk)
31. linked_list_problems: Linked list implementation and functions from meetup group assignment
32. bst_lca: Find lowest common ancestor for Binary Search Tree (no parent pointers given)
33. dictionary: Implement a word dictionary with following funcs: (i) insert word, (ii) search is a word is present 
	(iii) find the number of words that start with a given prefix (prefix cannot be the whole word)
34. sorted_quadratic: Solution to problem 7 below
35. Implement stack using a queue
36. k_sorted_merge: merge k sorted lists in an efficient way Hint: O(nklogk) where n is the avg. len of lists
37. wsum: Find the first index i of an array (of integers) such that sum of elements over a window is equal to target.
38. max_sum_subarray: Given a integer array, find the max sum of a subarray (need to be contiguous in the original array)
39. btree_lca: Lowest Common Ancestor in a B Tree
40. Binary search tree find k-th largest node 
41. ks: Knapsack implementation in DP
42. mcnugget: Given you can buy mcnuggets in pack of 4, 6, and 9; can you buy N macnuggets? E.g. 17 can be bought by (4+4+9)?
	Hint: Try DP? Can you generalize for more than 3 pack sizes. Note gcd(size1, size2, size3, ...sizen) = 1
43. num_subsequence: Problem solving Workshop #21 March 4,2017
44. parentheses: Given a set of braces determine if the order of parentheses is correct. E,g, ((())()) is correct , but ()()) is wrong
45. thread_order: execute threads in a specific order; e.g. 3 threads with id 1,2, and 3 should execute first() second() and third() in order
46. TestCache: sample cache implementation using LRU policy
47. Wiggle_sort: Given an unsorted array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....
48. Wrap_Around_Substr_Search: Unique substrings in a given pattern within a infinite length wraparound string "....abcd....xyzabcd...."
49. Game Of Life: Design a game of life matrix: cell of 0's and 1's with conditions: 
	(i) any alive cell with <2 neighbor dies
	(ii) any alive cell with 2-3 neighbors lives
	(iii) any  alive cell with > 3 neighbors dies
	(iv) any dead cell with 3 live neighbors becomes alive
50. PatternRule: Given a string and a pattern, find if the string follows the pattern rule. E.g. p="aba", str="boygirlboy" --> TRUE 
51. TreeZigZag: given a binary tree, find the zig zag traversal for the tree (leetcode)
52. MatrixZero: Given a matrix convert the row and column of an element == 0 with all zeros; do this with constant space O(1)
53. 01matrix: Given a matrix find the closest path of an element from the nearest zero. (leetcode)
54. btree_inc_decr_sequence: Given a btree find the longest increasing or decreasing sequence (does not have through go through root).
55. repeated_substring: Given a string S, find the longest repeated substring (i.e. s="abcdabc" o/p="abc") -- Eugene Problemset 4/29
56. count_repeated_elem: Count repeated elements in a sorted array (hint: using binary search) 
57. longest_palindrome (leetcode): Find the longest palindromic substring in a given string
58. integre_dist: Given an array (i) Determine whether there exists 2 values in array with a diff at most k (i.e. diff <=k)
	(ii) whether two values in the array are identical and their diff in position is at most D (i.e. a[i]==a[j] && j-i<=D
	(iii) both (i) and (ii) hold true i.e. two values separated by at most D distance and diff in values of at most K 
	Note: for iii. better complexity is O(NlogN) and best is O(n)
59. sbbst: Self balancing BST methods and functions
60. Array_shuffle:
61. Given a non-empty integer array, find the minimum number of moves required to make all array elements equal, where a move 
	is incrementing a selected element by 1 or decrementing a selected element by 1.
62. recursive_file_read: Read a directory recursively and print all file
63. islands: Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water 
	and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are 
	all surrounded by water. Leetcode (#200)
64. Given a network in the form of a binary tree where edge weight represent network latency, find the path between two nodes with max latency
	NOTE: Path can not pass through root; HINT: similar to diameter of binary tree
65. unique_array_element: Given an array of repeated elements find the following (i) arra has all element repeated twice except one
	(ii) array has all elements repeated twice except 2; (iii) array has all elements repeated 3 times except 1 (tricky)
	All time complexity O(n) and space O(1)
66. test_recursive_file: Iterate through the current directory structure and list all files recursively. (sample code from internet)
67. Producer_consumer: Sample producer consumer problem with shared global Queue using mutex and cond_vars
68. coin_change: Given a denomination vector e.e. {1,2,5,10} find how many ways can you generate change for value V (e.g. 17)
	Hint: DP similar to Knapsack problem; assume we have infinite supply of coins for the given denominations
	Part (ii): slight modification: find the minimum number of coins that can be used to make up the denomination
69. sort_0_1: Given an array or 0's and 1's, sort them such that 0's appear and then 1's appear. e.g. Arr={0,1,1,0, 1,1} --> {0,0,1,1,1,1} 
70. Clone a linked list with next and random pointer - ideally in O(1) space. You are given a linked list which has a next pointer and a random pointer that points
 to a random node in the linked list. You have to clone this linked list.
71. AlphaOrder: Given the alphabet encoded as numbers (e.g., a=1, b=2, ..., z=26), and a sequence of numbers (e.g., "23413259802"), how many strings can be generated
72. SubArrayMaxAverage: Given an array with positive and negative numbers, find the maximum average subarray of given length
73. BigNum: Big number addition: Given two big numbers inputed as strings find the sum of the numbers and output it as a string of another big number
75. Calendar_conflict: Given n appointments/meetings each with a start time and end time, find all conflicting meetings; find slots which does not have conflicts.
76. BitAdder: Given two binary strings, print the binary sum of the two. E.g. A="101" B="110" C="1011"
77. BracketValidator: Find if the string has unmatched brackets; alternative version: String Parenthesis normalizer/sanitizer "((((a(b)" ==> "a(b)"  
78. ShuffledDeck: Find is a shuffled deck of cards is a single riffle. (check riffle definition) 
79. SearchRotated: Search a number is a arbitrarily rotated sorted array e.g. arr={4,5,6,7,1,2} search (arr,1) should return 4
80. NumberPad: Given a numerical string for a telephone number, print all possible strings for the same
81. SortedIntersection: Find intersection of two sorted arrays. Use constant space and linear time algorithms.
82. segment_tree: Sample Segment Tree implementation using multi-level arrays with UpdateSum() procedure.  
83. Longest repeated substring ( in O(logN)): Given a string S, find the longest repeated substring (i.e. s="abcdabc" o/p="abc") -- Eugene Problemset 4/29
84. Cow-stable problem: Given a set of x-coordinates and num of cows distribute the cows such that max separation is achieves (cows close to each other will fight). Hint: Exp Binary Seacrh?
85. Predecessor timestamp problem: Given hahs table with funcs: Get(key, timestamp) and Put(key, val, tstamp) find the last value with timestamp <= given timestamp(T);
86. Sorted intersection of arrays of size N and M in ~O(M) time where M ~ N; for N >> M expected complexity O(MlogN). Note previous implementation was O(M+N).   


## Todo Questions:
1. Robot Problem
2. Using Rand5() design Rand7()
4. Print perimeter of a tree
5. Closest 3sum problem
6. Given text file T and pattern P, find the number of anagrams of 'P' in T; (any ASCII character is a valid literal)
7. Suppose you have a sorted array of integers. Apply a function of the form f(x)=a*x^2 + b*x + c to each element 
	x of the array such that the resulting array is still sorted. Can you do it in better than O(n)
8. Given two n-ary trees how can you check if their structures are same? (Hard problem).
10. Given a array of integers find if there is an element whose sum in the left side == sum in the right side (exclude the element).
	Optimize the solutions to use O(1) space and O(n) time complexity
11. Given a string return whether there is a permutation that is a palindrome
12. Given a double array find the max product of a subarray (contiguous in original array). [Hint: Kaden's algo]
13. Find the number of primes within internal A and B (A, B large numbers) try to optimize the time complexity to O(B-A) i.e. range.
14. Given a list of nodes (integer values) in in order traversal, reconstruct a balanced binary tree.
	[Hint: Standard approach O(nlong); can you do it in O(n)? Create empty btree?]
15. Given a singly linked list; Zip it. Zip operation is defined as follows:
	1->2->3->4->5 ==> 1->5->2->4->3 (like a zipper); 1->2->3->4->5->6 ==> 1->6->2->5->3->4
	Notes: Try doing this with O(1) space
16. Find the longest increasing snake in a 2D matrix of integers. Note that you can go in all 8 directions  left right top down and diagonal.
	Return the longest path size. Tip: snake is a increasing subsequence where next elements differs by 1 (consecutive path)
17. Given a huge database with n records (equivalent to n rows), a user has changed k records (k << n) without remembering which k-s have been
	changed; how to sync the database without transmitting the entire database.
18. Given a number M find the smallest number consisting of 0's and 1's that is divisible by M; e.g. for M=3 Num=111, for M=7 Num=1001
	Hint: [link](http://math.stackexchange.com/questions/388165/how-to-find-the-smallest-number-with-just-0-and-1-which-is-divided-by-a-give)
	Target O(M^2)
19. Find the longest substring without repeating characters in a given string. target O(n) time. Hint: similar to max_sum
20. Word break: Given a sentence (in string) return true if the characters can be broken down into words which are contained in a dictionary.
	You are given the Dict structure with a method Dict.contains(string s) which return true or false.
21. Given a map with Parcels (small variable boundary area) and given a list of points find the parcels with the max points. 
	Note: you are given an API bool PointInParcel(lat, long, parcel_boundary) which returns true or false; but this is very expensive
	So ideally you would like to minimize the calls to this function as much a possible. If needed you can represent the parcels
	using your own data structures.
22. Anti-spiral traversal in a matrix
23. Given an NxM matrix representing a garden, each entry contains a number, which is the number of carrots in that plot. 
	A rabbit starts in the middle, or plot closest to middle with the most carrots and always goes up down left or right after eating 
	to the plot with the most carrots. No negative entries, write some code that given a garden matrix, you return the number of carrots 
	a rabbit would eat in it.  
24. Hills size and sea level problem. Given a matrix Hill MxN dimension Hill[i][j] represents the height of the hill for that point
	There is also sea level which can rise and cover the height of the hills.
	Given the Hill matrix a hiker wants to cross from the leftmost column to the rightmost column of this matrix;
	What is the minimum change in sea level that will ensure that hiker cannot reach the other side (i.e. the last column)
	Assume the initial sea level is 0; HINT: DP approach
25. Given an array of integers ,find a triplet i,j,k such that sum(arr[0..i-1])==sum(arr[i+1..j-1])==sum(arr[j+1...k-1])==sum(arr[k+1...end])
	Q: can you do better than O(n^3)
26. Given a running stream of bits find the occurrence of a given short integer. Note the the incoming stream is endless and every time the 
	short value is encountered a flag is raised. How to efficiently represent the data structs for doing this operation?
27. Print Binary Tree in level order:
28. Circular Queue implementation [queue which wraps around with the last element]
29. From an input stream read a short byte and raise a flag when it is of value k (eg. k=0xab)
	//try doing in single pass
30. Pizza Order: In a pizza shop customers order pizzas and based on the pizza then order it takes a certain amount of time. The pizza cook 
	can only cook one pizza at a time; if the customers are served on a first come first serve basis the waiting time can be significant;
	Given a times at which the orders come to the pizza shop and the time it takes to make each such order, how can you minimize the 
	average waiting time for each order. Waiting time for each order = (finish time)-(start time for that order);
	E.g. start_time = {1,2,3} cook_time = {2,9,6}; if the pizzas are cooked in order total wait time = (3-1)+(12-2)+(18-3)= 27
	However; if the pizzas are cooked in the order 1,3,2 then the wait time = (3-1)+(9-3)+(18-2)=24; hence avg waiting time is less in second case
31. Interval_Tree: Find calendar appointment conflicts using interval tree; Alternate implementation for Calendar_conflict problem
32. Implement strstr(string s1, string s2) method: Search a string s2 in a given string s1, return the position at which match occurs. 
        Tweak: return all positions of match if there are multiple occurrences or s2 in s1
33. Expression evaluation: Evaluate an expression containing integers, '+', '-', '*', '/' and '()'
34. SearchWord: Given a set of dictionary words search a query string. The query string can have a '*' in it indicating any character can take that spot.
    E.g. Word List = ["peak", "pete", "peek", "pint"}  Query: "p**t" -> true, "p*nk" -> false, "peak" -> true etc.
35. 

