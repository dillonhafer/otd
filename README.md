# OTD
On This Day In History

Display a random event that happended on this day in history

## Usage

A. Import library

```go
import (
  "github.com/dillonhafer/otd"
)
```

B. Use one of two methods

`Events` return a `[]string` of this day's events.

```go
otd.Events()
```

`RandomEvent` takes a `[]string` of events and returns a `string` of a random event

```go
events := otd.Events()
event  := otd.RandomEvent(events)
```
