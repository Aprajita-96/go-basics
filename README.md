# go-basics


goto-app 
    URL shortner application
        web interface to submit the long url
        generate short verion of the given url
        persisted
        when the short url is requested, redirect the user to the original (long) URL

verion-1
    In memory data store

    run the server
    hit http://localhost:8080/add -> to create a short verson of a long url
        => respond with the short url

Version-2
    Persisting the data in the file (gob)

Version-3
    Making the persistence logic async using goroutines and channels

Version-4
    replacing gob with json
