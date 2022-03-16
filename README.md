This library exposes an interface for implementing lazy, concatenated queries over csv file.\
The library also provides lazy implementation of avg, ceil, filter-rows and get-column.\

The function _With_ concatenates the next operation to the current chain,
and the function _Run_ returns an element (channel) which can be iterated with _range_\

In order to add new operation, the interface _Operation_ have to be implemented.
(In go there is no abstract class, but you can get the basic implementation of _Init_ and _With_, you need to add the * prc.BaseOperation field to your struct)