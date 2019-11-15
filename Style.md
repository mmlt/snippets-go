## Go naming tips
https://github.com/golang/go/wiki/CodeReviewComments
https://golang.org/doc/effective_go.html#names
https://peter.bourgon.org/blog/2019/04/24/go-naming-tips.html

In my experience, the best way to name types in Go is as follows:

- Structs are plain nouns: API, Replica, Object
- Interfaces are active nouns: Reader, Writer, JobProcessor
- Functions and methods are verbs: Read, Process, Sync

Besides being consistent, this model has two nice secondary benefits. 

First, doc comments become natural and fluent.
```
// Sync the local replica state to the provided upstream.
func (r *Replica) Sync(u Upstream) error { ... }

// Process the next available job from the queue
// and emit results to the sink.
func Process(q JobQueuer, s Sinker) error { ... }
```
Second, if functions are verbs, it seems to help a block of code read more fluently. 
That is, each expression (or expression-block) becomes kind of like a sentence.
```
objectNoun := Verb(subjectNoun, subjectNoun, ...)
```
Through this lens, each function is like a coherent paragraph of prose. 
I think this is a great metaphor, as it naturally reinforces other virtues, 
like appropriate length, single responsibility, and independent testability.