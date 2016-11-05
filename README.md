# golang-etl
A sample Go application showing the basics of ETL


## Version-1

This version doesn't take advantage of any of the concurrency features in Go. This gives us a benchmark to show how long execution takes if extract, transform and load are performed in series.

```bash
cd .\golang-etl
go build
.\etl
19.661943479s

head dest.txt
         Part Number        Quantity    Unit Cost   Unit Price      Total Cost    Total Price
            76502369               7        12.35        16.47           86.45         271.26
            56886748              18        24.94        30.92          448.92         956.05
            62283911               7        15.50        21.11          108.50         445.63
            40749682              96        14.45        19.50         1387.20         380.25
            59333352              73        25.88        29.43         1889.24         866.12
            32895121              61        28.99        32.57         1768.39        1060.80
            30174319              22        29.51        33.43          649.22        1117.56
            41555059              44        27.38        32.97         1204.72        1087.02
            81912584              99        26.86        31.94         2659.14        1020.16
```

**Time taken:** 19.66s

## Version-2

This version uses one Go routine for each of the three stages; extract, transform and load. This shaves a little time off the overall processing but isn't a dramatic improvement. If we think about it, each Go routine includes a simulated delay for processing each record. The Go routine is tied up waiting for each of these delays to comlete sequentially. It is these delays that make up the bulk of our processing time.

**Time taken:** 13.66s

## Credit

Source: [Go Concurrent Programming](http://www.pluralsight.com/courses/go-concurrent-programming)
