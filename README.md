This library exposes an interface for implementing lazy, concatenated queries over csv file.\
The library also provides lazy implementation of avg, ceil, filter-rows and get-column.\

The function _With_ concatenates the next operation to the current chain,
and the function _Run_ returns an element (channel) which can be iterated with _range_\

In order to add new operation - \
If the operation is _Map_ operation create it with _NewMapOperation_. \
If the operation is _Reduce_ operation create it with _NewReduceOperation_.