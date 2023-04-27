# STARS

> Every star has its own story.

![stars](images/stars.png)

STARS is a component library for Go.

## Install

```shell
go get github.com/B1NARY-GR0UP/stars
```

## Component List

### Algorithm

| star           | code                                       |
|----------------|--------------------------------------------|
| consistenthash | [consistenthash](algorithm/consistenthash) |
| saltencryption | [saltencryption](algorithm/saltencryption) |

### Data Structure

| star      | code                                 |
|-----------|--------------------------------------|
| hashtable | [hashtable](datastructure/hashtable) |
| unionfind | [unionfind](datastructure/unionfind) |

### Sync

| star          | code                                   |
|---------------|----------------------------------------|
| chanmutex     | [chanmutex](sync/chanmutex.go)         |
| reentrantlock | [reentrantlock](sync/reentrantlock.go) |
| mutexinfo     | [mutexinfo](sync/mutexinfo.go)         |
| singleflight  | [singleflight](sync/singleflight.go)   |
| once          | [once](sync/once.go)                   |

### Other

| star      | code                         |
|-----------|------------------------------|
| mapreduce | [mapreduce](other/mapreduce) |

## License

STARS is distributed under the [Apache License 2.0](./LICENSE). The licenses of third party dependencies of STARS are explained [here](./licenses).