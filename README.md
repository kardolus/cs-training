# Computer Science Training ![progress 28/36](https://img.shields.io/badge/progress-78%25-blue) ![license](https://img.shields.io/github/license/kardolus/cs-training)

![gopher from ashleymcnamara](https://raw.githubusercontent.com/ashleymcnamara/gophers/master/TEACHING_GOPHER.png "Logo Title Text 1")

I suggest using this list as a warm-up for interview preperation. Follow it up with interview problems from [HackerRank](https://www.hackerrank.com/interview/interview-preparation-kit), [Cracking the coding interview](http://www.crackingthecodinginterview.com/) or [Programming Interviews Exposed](https://www.oreilly.com/library/view/programming-interviews-exposed/9781118283400/). Alternatively go old school with [Project Euler](https://projecteuler.net/).

It can help to practice using [coderpad.io](https://coderpad.io/), since it is used for many coding interviews.

## Concepts
| Concept | Description | Status |
|:-------:|:------:|:------:|
| [Big O](/concepts/bigO.md) | Growth of run time and space complexity |  ✔ |
| [Divide and Conquer](/concepts/divide_and_conquer.md) | Design paradigm based on multi-branched recursion | ✔ |
| Dynamic Programming | An optimization over plain recursion |  ✘ |
| [Memory (Stack vs Heap)](/concepts/memory_stack_heap.md)| Memory allocation | ✔ |
| [Bitwise Operations](/concepts/bitwise.go) | Operations on ints and uints at the binary level | ✔ |
| [Hashing](/concepts/hashing.md) | Map data of arbitrary size to fixed-size values | ✔ |

## Data Structures
| Structure | Access^ | Search^ | Insertion^ | Deletion^ | Status |
|:-------:|:------:|:------:|:------:|:------:|:------:|
| [Stack](/data_structures/stack.go) | Θ(n) |	Θ(n) | Θ(1) | Θ(1) | ✔ |
| [Queue](/data_structures/queue.go) | Θ(n) |	Θ(n) | Θ(1) | Θ(1) | ✔ |
| [ArrayList](/data_structures/array_list.go) | Θ(1) | Θ(n) | Θ(n) | Θ(n) | ✔ |
| [Singly LinkedList](/data_structures/singly_linked_list.go) | Θ(n) |	Θ(n) | Θ(1) | Θ(1) | ✔ |
| [Doubly LinkedList](/data_structures/hash_map_doubly_linked_list.go) | Θ(n) |	Θ(n) | Θ(1) | Θ(1) | ✔ |
| SkipList | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | ✘ |
| [Set](/data_structures/set.go) | Θ(n) |	Θ(n) | Θ(1) | Θ(n) | ✔ |
| [HashTable](/data_structures/hash_map_doubly_linked_list.go) | N/A | Θ(1) | Θ(1) | Θ(1)| ✔ |
| [Binary Search Tree](/data_structures/binary_tree.go) | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | ✔ |
| Trie | | | | | ✘ |
| Red-Black Tree | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | ✘ |

^ = average

## Sorting
| Algorithm | Best | Average | Worst | Space (Worst) | Status |
|:-------:|:------:|:------:|:------:|:------:|:------:|
| [Bubble Sort (Exchange)](/sorting/bubble_sort.go) | Ω(n) | Θ(n^2) | O(n^2) | O(1) | ✔ |
| [Quick Sort (Exchange)](/sorting/quick_sort.go) | Ω(n log(n)) | Θ(n log(n)) | O(n^2) | O(log(n)) | ✔ |
| [Insertion Sort](/sorting/insertion_sort.go) | Ω(n) | Θ(n^2) | O(n^2) | O(1) | ✔ |
| Heap Sort (Selection) | Ω(n log(n)) | Θ(n log(n)) | O(n log(n)) | O(1) | ✘ |
| [Merge Sort](/sorting/merge_sort.go) | Ω(n log(n)) | Θ(n log(n)) | O(n log(n)) | O(n) | ✔ |

## Common Problems
| Problem | Description | Status |
|:-------:|:------:|:------:|
| Maximum subarray | Find a contiguous subarray with the largest sum | ✘ |
| [Fibonacci iterative](/sorting/fibonacci.go) | Approximation in integers to the logarithmic spiral series | ✔ |
| [Fibonacci recursive](/sorting/fibonacci.go) | ,, | ✔ |
| [Fibonacci dynamic programming](/sorting/fibonacci.go) | ,, | ✔ |

## Creational Patterns
| Pattern | Description | Status |
|:-------:|:------:|:------:|
| Abstract Factory | Provides an interface for creating families of releated objects | ✘ |
| [Builder](/patterns/creational/builder.go) | Builds a complex object using simple objects | ✔ |
| [Factory Method](/patterns/creational/factory_method.go) | Defers instantiation of an object to a specialized function for creating instances | ✔ |
| Prototype | Co-opt one instance of a class for use as a breeder of all future instances | ✘ |
| [Singleton](/patterns/creational/singleton.go) | Restricts instantiation of a type to one object | ✔ |

## Structural Patterns
| Pattern | Description | Status |
|:-------:|:------:|:------:|
| [Adapter](/patterns/structural/adapter.go) | Wrap an existing class with a new interface | ✔ |
| Bridge | Decouples an interface from its implementation so that the two can vary independently | ✘ |
| Composite | Encapsulates and provides access to a number of different objects | ✘ |
| Decorator | Adds behavior to an object, statically or dynamically | ✘ |
| Facade | Uses one type as an API to a number of others | ✘ |
| Flyweight | Reuses existing instances of objects with similar/identical state to minimize resource usage | ✘ |
| [Proxy](/patterns/structural/proxy.go) | Provides a surrogate for an object to control it's actions | ✔ |

## Behavioral Patterns
| Pattern | Description | Status |
|:-------:|:------:|:------:|
| Chain of Responsibility | Avoids coupling sender and receiver | ✘ |
| Command | Bundles a command and arguments to call later | ✘ |
| Interpreter | Define a grammatical representation for a language | ✘ |
| [Iterator](/patterns/behavioral/iterator.go) | Access the elements of an aggregate object sequentially | ✔ |
| Mediator | Connects objects and acts as a proxy | ✘ |
| Memento | Generate an opaque token that can be used to go back to a previous state | ✘ |
| [Observer](/patterns/behavioral/observer.go) | Provide a callback for notification of events/changes to data | ✔ |
| State | Encapsulates varying behavior for the same object based on its internal state | ✘ |
| [Strategy](/patterns/behavioral/strategy.go) | Enables an algorithm's behavior to be selected at runtime | ✔ |
| Template Method | Defines a skeleton class which defers some methods to subclasses | ✘ |
| Visitor | Separates an algorithm from an object on which it operates | ✘ |

## Acknowledgement
Got the table design and a bunch of pattern-definitions from this cool repo: https://github.com/tmrts/go-patterns 
