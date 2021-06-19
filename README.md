# Pq-Listen-Notify

postgresqlのListen, Notifyで遊ぶ


## Notifier

```
 go run ./notifier/main.go "user= dbname= password= host= sslmode=disable"
```

## Listner

```
 go run ./listener/main.go "user= dbname= password= host= sslmode=disable"
&{131 channel 37a825d5-51f8-4361-8f8a-833b2e7f9b96}
&{131 channel 4dcea0e3-850a-4b39-b28a-1747ab05be4d}
&{131 channel 84e29037-cc9e-460e-a33b-20d5a592b79c}
&{131 channel b2f120e7-2538-4a96-a6b2-3c9e54d4d604}
&{131 channel 42699ddf-db91-4cd4-8780-32f938363e04}
```
