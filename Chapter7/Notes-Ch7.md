# Head First Go - Notes - Chapter 7

## Note 1

### Maps

A map is a collection where each value is stored under a corresponding key. Whereas arrays and slices can only use integers as indexes, a map can use (almost) any type for keys.

All of a map's keys must be the same type, and all the values must be the same type, but the keys don't have to be the same type as the values.

## Bullet Points

+ When declaring a map variable, you must provide the types for its keys and its values:
`var myMap map[string]int`

+ To create anew map, call the *make* function with the type of the map you want:  
`myMap = make(map[string]int)`

+ To assigna value to a map, provide the key you want to assign it to in square brackets: 

*mayMap["my key"] = 12*

+ To get a value, you provide the key as well: 

`fmt>println(myMap["my key"])`

+ You can create a map and initialise it with data at the same time using a **map literal**:

`map[string]int{"a": 2, "b": 3}`

+ As with arrays and slices, if you access a map key that hasn't been assigned a value, you'll get a zero value back.

+ Getting a vlue from a map can return a second, optional Boolean value that indicates whether that value was assigned, or if it represents a default zero value:

`value, ok := myMap["c"]`

+ If you only want to test whether a key has had a value assigned, you can ignore the actual value using the _ blank identifier.

`_, ok := myMap["c"]`

+ You can delete keys and their corresponding values from a map using the *delete* built-in function:

`delete(myMap, "b")`

+ You can use *for...range* loops with maps, much like you can with arrays and slices. You prpvide one variable that will be assigned each key in turn, and a second variable that will be assigned each value in turn.

```go
for key, value := range myMap {
  fmt.Println(key, value)
}
```