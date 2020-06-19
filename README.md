# sampling

This repo offers implementations of a couple of random sampling algorithms -- simple reservoir sampling, slightly optimized reservoir sampling, and weighted reservoir sampling. Reservoir sampling is an algorithm that randomly selects *k* samples from a collection of *n* elements and is a technique best used when the collection is large enough to exceed memory limits.

To learn more: check out my [blog post](https://www.lukechui.com/post/reservoir-sampling)!

## Installation

To install the binaries, run the following command
```
$ go get -u github.com/epicchewy/sampling/cmd/...
```

This repo can also be imported as a package
```
$ go get -u github.com/epicchewy/sampling
```

## Installation + Usage

This repo provides two binaries (`ssampling` and `rsampling`) that users can run out of box. It can also be imported as a package and extended. 

#### Binary Example

Installing the binaries: 
```
$ go get -u github.com/epicchewy/sampling/cmd/...
```

Running the sampler:
```
$  seq 1 100 | rsampling --samples 5
27
37
85
80
45
```

#### Code Example

Import the package to use in a project
```
$ go get -u github.com/epicchewy/sampling
```

Sample Usage:
```go
import "github.com/epicchewy/sampling"

func main() {
    var items []int
    for i := 0; i < 100; i++ {
        items = append(items, i)
    }

    // use available iterators or pass in a custom one that implements the Iterator interface
    itr := sampling.NewIntInterator(items)
    r := sampling.NewSimpleReservoir(10, itr)
    samples := r.Sample()
}
```

## References

- [Wikipedia Page](https://en.wikipedia.org/wiki/Reservoir_sampling)
- [Reservoir Sampling Whitepaper](http://www.cs.umd.edu/~samir/498/vitter.pdf)
- [Weighted Reservoir Sampling Whitepaper](https://arxiv.org/pdf/1904.04126.pdf)
