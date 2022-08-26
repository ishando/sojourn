# Maths

A simple maths api that will perform operations on a provided set of data.

# Endpoints

The api listens to `localhost:8338` with the following endpoints:

- `/min` - returns the _n_ smallest numbers from the dataset. If _n_ is greater that the size of the array the whole array is returned. If the array is empty it will return an empty set, otherwise will always return a set with at least one value.
- `/max` - returns the _n_ largest numbers from the dataset. If _n_ is greater that the size of the array the whole array is returned. If the array is empty it will return an empty set, otherwise will always return a set with at least one value.
- `/avg` - returns the arithmetic mean of the dataset
- `/median` - returns the median of the dataset
- `/percentile` - return the value of the _pth_ percentile for the dataset. This uses the nearest rank method, so the value returned is always one that is in the data set. This method a value _v_ such that no more than _p_ percent of the data is strictly less than _v_ and at least _p_ percent of the data is less than or equal to _v_.

# Request

The request accepts a json object with two attributes:
- nums - an array of numbers to perform the action on
- qualifier - varies per operation
  - min - number of values to return
  - max - number of values to return
  - avg - not used
  - median - not used
  - percentile - the percentile value to return

```json
{
  "qualifier": 1,
  "nums": [1,2,3]
}
```

# Response

The response is a json object which will have one of two attributes:
- answer - the single value result for `avg`, `median` and `percentile`
- answers - an array of values for `min` and `max`

```json
{
  "answer": 2
}

or

{
  "answers": [1,2]
}
```