# tribe-generator

A generator for a tribe history : births, deaths, and people joining the tribe. Its goal is to create a plausible history and number of people at the end of the given period.

It has been tuned to output approximatly 200 people after arount 300 years.

## running

```
$ go run tribe.go
```

## tuning
Lots of parameters can be tuned : either the events (when the death and natality have been impacted by outside events), the probabilities for birth and death, or even the names of the people can be tuned.

You can change all that by forking the repository and editing the relevant variables in the code
