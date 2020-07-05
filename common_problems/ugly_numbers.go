// Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. 
func nthUglyNumber(n int) int {
    if n <= 6 { // preprocess
        return n
    }
    
    count := 6
    
    for number := 7; true; number++ {
        // if the number is a prime larger than 5, it is not an ugly number
        if isPseudoPrime(number) && isPrime(number) {
            continue 
        }
        
        // if the number has a prime factor larger than 5, it is not an ugly number
        if number, _ := getPrimeFactor(number); number > 5 {
            continue
        }

        count++
        
        if count == n {
            return number
        }
    }
    
    return -1 
}

func isPrime(number int) bool { // Trial Division O(n^0.5)
    if number % 2 == 0 {
        return false
    }
    
    const errorMargin = 1e-9
    
    sqrt := math.Floor(math.Pow(float64(number), 0.5))
                                    
    for i := float64(3.0); i <= sqrt; i = i + 2.0 {
        division := float64(number) / i
        if math.Abs(math.Trunc(division) - division) < errorMargin {
            return false
        }
    }
    
    return true
}

// Fermat's little theorem: (base^{number-1} - 1) % number == 0
// Note that Rabin-Miller extends this idea and is more accurate
func isPseudoPrime(number int) bool { 
  const (
    base       = 2.0
    precission = 1e-9
  )
	
  return math.Mod(math.Pow(base, float64(number)-1.0)-1.0, float64(number)) < precission
}

// Pollard's rho algorithm
func getPrimeFactor(number int) (int, error) {
    input := number
    
    x := 2
    y := 2
    d := 1
    
    for d == 1 {
        x = polynomial(x) % number
        y = polynomial(polynomial(y) % number) % number
        d = greatestCommonDivisor(abs(x-y), number)
        
        if d == number { // TODO proper error handling
            fmt.Printf("Failure for number [%d], x [%d], and y [%d]\n", number, x, y)
            return 0, errors.New("Failure")
        }
    }
    
    // This is an attempt to return the largest prime factor (Not part of Pollard's Rho)
    if !isPseudoPrime(d) {
        d = number / d
    } else if number / d > d && isPseudoPrime(number / d) {
        d = number / d
    }
    
    fmt.Printf("Prime factor for [%d]: [%d]\n", input, d)
    
    return d, nil 
}

func polynomial(x int) int {
    squared := math.Pow(float64(x), 2.0)
    
    return (int(squared) + 1)
}

func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}

// Euclidian Algorithm
func greatestCommonDivisor(a, b int) int {
  for {
    if r := a % b; r == 0 {
      return b
    } else {
      a = b
      b = r
    }
  }
}

