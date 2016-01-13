### SCRAPED.

### UPDATES

As the exercise was completed in under 2 hours improvements could of course be made. No time to add concurrency right now but at least now:

* The scraper is returned in a factory function.
* The document is now a field on the scraper struct.
* The document is created using an ```io.Reader``` value for easier testing.
* This means the tests can now pull raw HTML from tests fixtures for parsing.
* Some concurrency has been added for element processing via a ```WaitGroup```.

### NOTES

* This scraper has a single dependency, [goquery](http://github.com/PuerkitoBio/goquery), which simplifies DOM traversal. Otheriwse it's all standard lib (apart from [mux](http://github.com/gorilla/mux) for the web part).
* Dependency management is handled using the Go 1.5+ vendor flag
and git submodules.
* The current tests hit real HTML endpoints (in gists) which are
slower than mocks but goquery ```NewDocument``` takes a URI to operate so used them for this pass. Extensions to this code would
probably swap to some form of mocking.
* The exercise required a console application. This solution delivers both a console and a web version as I had time!
* Time taken to build this solution: 1 hour 45 minutes.

### BUILD IT

1. Clone this repository into your GOPATH.
2. Run ```source dev_env``` to set the Go 1.5+ vendor flag environment variable.
3. Run ```git submodule init``` and ```git submodule update```.

### RUN IT - CONSOLE MODE

1. In the ```console``` folder run ```go build```.
2. Run ```./console``` and bathe in the JSON goodness or ```./scraped > scraped.json```.
3. To hit another page with fruity goodness, run ```./console -URI=https://gist.githubusercontent.com/buddhamagnet/c6997464d84b8bf379a1/raw/0b2e8ca65fed073197cd7be1e91a163738488f2e/fruity.html```.

### RUN IT - WEB MODE

1. In the ```web``` folder run ```go build```.
2. Run ```./web``` and hit ```localhost:9494/scrape```.

### TEST IT

1. Run ```go test ./...```.
