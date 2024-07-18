# Data Structures and Algorithms in Go

A recap to data structures and algorithms enjoying [ThePrimeagen](https://github.com/ThePrimeagen) and his [free course](https://frontendmasters.com/courses/algorithms/).

Took the chance to try out go the first time.

## 1. Arrays / Array Lists

I used slices to implement those algorithms. Technically
they are no arrays but I wanted to be flexible regarding the length
of the lists.

### Search Algorithms

- [x] [Linear Search](./arrays/linearsearch.go)
- [x] [Binary Search](./arrays/binarysearch.go)

### Sorting Algorithms

- [x] [Bubble Sort](./arrays/bubblesort.go)
- [ ] Quick Sort
- [ ] Merge Sort

### Other Algorithms

- [x] [Two Crystal Ball Problem](./arrays/bubblesort.go)
- [ ] Trapping Rain Water Problem

### Abstract data structures

- [x] [Ring Buffer](./arrays/ringbuffer.go)

## 2. Linked Lists

Linked Lists have there advantages when it comes to removing data
from the end of a queue. The run time will be always O(n).

While deleting items from the end of an array is also an operation of O(n),
the deleting from the beginning of an array is not. In order to delete the
first item in a queue using a list you have to delete the first item and
shift all other items one index to the front. This runs with O(n).

Therefore, when you don't or rarely need to access the items
in the middle of those, a linked list is a solid choice for implementing queues.

- [x] [Queue](./lists/queue.go)
- [x] [Stack](./lists/stack.go)
