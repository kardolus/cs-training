The book "Introduction to algorithms" has a nice quick overview of hashing.

* The **Division method** works as followed: you hash all input values to an array of size *m*. 
  ```
  h(k) = k mod m
  ```
  You take the *modulus m* of the input value *k*. It is important to chose a prime number *m* that is not too 
  close to a power of 2.
 
 * In the **Multiplication method** you apply the following calculation
    ```
    h(k) = Floor(m (kA mod 1))
    ```
    Where A is between 0 an 1. `kA mod 1` is simply the remainder of a multiplication with k. 
    You then multiply with the size of your hash array *m* and take the Floor. The advantage compared to the Division 
    method is that the size m is unimportant. You can choose a power of 2.
  
  * With **Universal hashing** you pick a random hash function (from a pre-difined set) at the beginning of execution. 
