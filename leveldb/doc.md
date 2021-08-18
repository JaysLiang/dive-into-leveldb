# slice

- slice is []byte wrapper

# BloomFilter

## murmur hash

## bloom filter

- double hash
  ```
  first murmur hash
  second shift hash (hash>>17|hash<<15) according k to hash again
  ```
# data struct
    internal key: 
        |key string| (seq|type) |
    Lookup Key
        |varint length | key string| (seq|type)|
    that means internal is part of Lookup key