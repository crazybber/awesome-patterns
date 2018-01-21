package main

// https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b

// Take io.Reader when you can
// If you’re designing a package or utility (even if it’s an internal thing that nobody will ever see)
// rather than taking in strings or []byte slices, consider taking in an io.Reader if you can for data sources.
// Because suddenly, your code will work with every type that implements io.Reader.

func main() {

}
