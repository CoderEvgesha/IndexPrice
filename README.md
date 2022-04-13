# Explanations for the decision #

### A general idea ###

The algorithm works as follows: there is a set of channels from which price data comes. Within 5 seconds, the application reads the data, and only data received with a delay of no more than a minute relative to the current time is taken into account. The rest are ignored. After 5 seconds, the reading is briefly interrupted and the application calculates the average of the available data. After counting, the value is averaged with the previous one (calculated one iteration ago). The result obtained is available from the outside online, the maximum delay of information according to this scheme is in the 5-second mode. Of course, all components are thread-safe.

### An application structure ###

All business logic is presented in the **app** package. This is not quite the Go way, but given that the functionality is presented in a truncated form, I assume that such a structure is acceptable.

The **ticker** folder contains the initially specified models in unchanged form.

The folder **average** contains all the functionality related to the calculation of the average value. The **Price** structure is a thread-safe average. For the purposes of encapsulation, only two functions of this structure are available outside the package: **Get** and **Calculate**. For calculation, the function takes as input the **TickerPrice** channel, which is read-only. During reading content of the channel, the price represented by the string type is converted into a decimal. If the conversion is not possible, this value is simply ignored. After calculating the average value based on the data from the channel, it is averaged with the current value and the data is updated.

The **source** folder includes a functional responsible for working with data sources. Here is the initially given interface **PriceStreamSubscriber**. The main type in the package is the **Pool**. Similar to the previous structure, it is thread-safe, the internal representation is hidden as much as possible for encapsulation purposes. Inside the **Pool** is a slice of the **Pair** type objects. The following mechanism is assumed - two channels come through the interface: **TickerPrice channel** type and **error channel** type. Both are read-only. They are formed in **Pair** and added into a slice. The **isClosed** flag indicates one channel of the pair was closed or has signals an error that came from the channel. Such pairs are periodically removed from the **subs** slice. 

The structure rpovides the following methods available outside the package: **GetSubscribers**, **AddSubscribers**, **Clean** and **Merge**. The first two are trivial, the third involves removing the pair with the true flag **isClosed** by rearranging the pair at the end and truncating the slice. The most interesting is the method **Merge** that implements the process of reading and combining data. In the loop, we go through all available sources. Within 5 seconds, data is read in parallel from each source into a single channel. The reading of an individual channel is interrupted if it has been closed or an error has occurred. Otherwise, the reading continues until interrupted by the context. The wait group allows you to wait for all goroutines to complete their work.

The **service** file is the main file that integrates the application logic into a single whole. Here is the interface **Storer** that is implemented by the **Pool** type. Although there is no clear need to implement this interface, this decision was made to improve cohesion, coupling, and improve the testability of the application. 

The main type is the structure **IndexService** with two methods: **Run** and **GetFairPrice**. For the first in an infinite loop, the channels is read and the average value is recalculated. Periodically there is a cleansing of sources from closed ones. The function can be stopped via context. The second method simply provides the functionality of getting the current value, which can be further used for any purpose: published, printed, etc.

### Testing ###

All functionality is covered by unit tests. The folder **fixtures** contains auxiliary functionality for setting up test data.

### That can be improved ###

This is the most interesting part because in my opinion, the current code is not ready for production right now, but bringing it to such a state would require a huge amount of time, which is not justified for a test task. Therefore, in this section, I will list everything that in my opinion is strictly necessary for production, and which can significantly improve the health of the project.

1. Infrastructure
- adding linters, analyzers, calculating test coverage, and so on
- adding git hooks, not being able to push changes if the tools described above are on
- creation of .sh scripts or a makefile in the "one-button action" paradigm for building, starting, stopping the application
- creating pipeline for Github actions
- calculation of metrics (both business and code)
- adding logging and monitoring
- adding docker or its equivalents

2. Tests
- building a full-fledged testing pyramid (as far as possible)
- adding component, integration and mutation tests

3. Implementation
- adding circuit breaker
- implementation in a more truthful Go way when adding functionality