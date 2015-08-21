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

## Example App

I wrote a [slack bot](https://github.com/dillonhafer/historybot) that uses this library.

## License

   Copyright 2015 Dillon Hafer

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
