# Demo Downloads Dashboard

This app is a demo of a dashboard that shows info about the downloads of an app. I developed it using:
* DB: RethinkDB
* Backend: Go
* Frontend: React

## Run

The best way to run this app is using my docker container:

`sudo docker run -p 8000:8000 -p 8001:8001 svallejo/downloads-dashboard`

This app is only a demo, so it doesn't need an external volume and includes the data for testing it.

## Expansion

This app was designed to be easy to expand and to maintain. It could be easily expanded to include a real time map, that would show the new data as soon as them are saved in the database.
To achieve this functionality the frontend and the backend need to be connected using the WebSockets protocol. In the backend, when the connection is made, a go routine would listen for the changes in the database and send the data to the client using a go Channel.
I chose RethinkDB for this demo because it is optimised for scalable real-time applications. A client can subscribe to a table and be notified of the changes and the new data. This is important especially if multiple applications are writing to the same database.