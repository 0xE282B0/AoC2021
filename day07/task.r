#https://stackoverflow.com/a/23678838/3801482
test <- scan("data/input.txt", integer(), sep = ",")
#test <- c(16,1,2,0,4,2,7,1,2,14)

# part1
sum(abs(test-median(test)))

#part2
# https://en.wikipedia.org/wiki/Triangular_number
triangular <- function(n) n*(n+1)/2
# A bit brute force:
x <- min(test):max(test)
y <- sapply(x,function(x){sum(sapply(test,function(t) triangular(abs(t-x))))})
#plot(x, y)
min(y)
