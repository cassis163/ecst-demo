## Event-Carried State Transfer (ECST) Demo
This demo showcases how event-carried state transfer can be utilized to reuse data from multiple services. In this case, article and stock data are reused to build the dataset for the catalog. Each article along with its stock is saved in a NoSQL database that is optimized for searching.
### Why even bother duplicating data?
As a software engineer, you should always desire to use the right tool for the job. However, oftentimes it is very hard to be technically able to do so. For example, you might not be able to create the solution that you intended, leaving you with a suboptimal solution that harms you in the long run.

This is a problem that I encountered in the past. My job was to create a materialized view of data that comes from multiple microservices. However, there was no standard for doing that, so we settled for using the shared database schema approach, where we can directly query the data from both services. Not to mention that a NoSQL database was ideal for the job, which I did not get to use for technical reasons.

In hindsight, duplicating the data in the form of events would have prevented all this. If I could have pulled this off, I would have been able to programmatically create my materialized view in a NoSQL database system like MongoDB, ElasticSearch or MeiliDB.
### Ok, well is ECST then?
ECST is an event-sourcing approach where events are stored as idempotent create, read and update events that contain the entire state, rather than only the changed properties of an object.
### Why ECST instead of more traditional event-sourcing solutions like ES + CQRS?
* In this example, it is not a requirement to save any business domain events specifically. Rather, the customer wants to be able to access all of our products along with any data that might be relevant such as stock. Business domain data is not desired, yet the underlying data itself is.
* ECST only produces create, update and delete events.
* The event log can be optimized and will not grow as much as it would with traditional event sourcing approaches. For example, all previous events of a certain entity can be deleted if it has changed. The old state is not important per se, unless time-travel is desired.
* It is easier to migrate existing systems to a ECST approach than a ES + CQRS one. You can build ECST around an existing system.
## Decoupling strategy
All event versions and types are stored in the ects-events library. This library is accessed by all services that require ECST and serves as a collection of contracts that all services have to comply to.

There is no such thing as complete decoupling. The best-effort to decouple services is to utilize a limited form of strict communication where services agree on accepting certain event formats.
### Architecture
-- TODO
