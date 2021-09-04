# Summary
```
PS: The tests were run on my local PC (8 GB RAM, AMD Ryzen 5 processor). Obviously on higher configuration machines, the performance would be better.
```

### Stress Testing

- 33,077 API calls spread amongst up till 400 concurrent users
- Average request duration of 0.77478 milliseconds
- 95% of the requests took less or equal to 0.489 milliseconds
- **Excellent performance** according to [Nielsen Norman Group](https://www.nngroup.com/articles/response-times-3-important-limits/)

### Spike Testing

- 56,550 API calls spread amongst up till 1,400 concurrent users
- Steady traffic of 100 concurrent users suddenly spiked to 1400 and then stabilized
- Average request duration of 0.8509 milliseconds
- **Excellent performance** according to [Nielsen Norman Group](https://www.nngroup.com/articles/response-times-3-important-limits/)
- **Mere 19.2% increase in average response time on 700% spike in concurrent load**

## Detailed results

##### Stress Testing (gradual increase to 400 concurrent users)

```
http_req_blocked...........: avg=24.41µs min=1µs med=4µs max=2.82ms p(90)=8µs p(95)=11µs
http_req_connecting........: avg=17.87µs min=0s med=0s max=2.76ms p(90)=0s p(95)=0s
http_req_duration..........: avg=774.78µs min=88µs med=660µs max=19.57ms p(90)=1.44ms p(95)=1.7ms
http_req_receiving.........: avg=44.83µs min=8µs med=35µs max=3.5ms p(90)=77µs p(95)=95µs
http_req_sending...........: avg=24.64µs min=5µs med=17µs max=3.31ms p(90)=34µs p(95)=54µs
http_req_tls_handshaking...: avg=0s min=0s med=0s max=0s p(90)=0s p(95)=0s
http_req_waiting...........: avg=705.3µs min=58µs med=584µs max=19.43ms p(90)=1.37ms p(95)=1.63ms
http_reqs..................: 99231 703.781021/s
iteration_duration.........: avg=1s min=1s med=1s max=1.02s p(90)=1s p(95)=1s
iterations.................: 33077 234.593674/s
vus_max....................: 400 min=400 max=400
```

##### Stress testing (steady traffic of 200 users)

```
http_req_blocked...........: avg=11.65µs  min=1µs  med=5µs   max=1.16ms  p(90)=8µs   p(95)=11µs
http_req_connecting........: avg=4.59µs   min=0s   med=0s    max=1.05ms  p(90)=0s    p(95)=0s
http_req_duration..........: avg=651.11µs min=91µs med=633µs max=35.6ms  p(90)=933µs p(95)=1.03ms
http_req_receiving.........: avg=46.09µs  min=10µs med=38µs  max=2.28ms  p(90)=76µs  p(95)=92µs
http_req_sending...........: avg=26.01µs  min=5µs  med=22µs  max=615µs   p(90)=37µs  p(95)=60µs
http_req_tls_handshaking...: avg=0s       min=0s   med=0s    max=0s      p(90)=0s    p(95)=0s
http_req_waiting...........: avg=579µs    min=61µs med=557µs max=35.54ms p(90)=860µs p(95)=962µs
http_reqs..................: 81090  574.98551/s
iteration_duration.........: avg=1s       min=1s   med=1s    max=1.03s   p(90)=1s    p(95)=1s
iterations.................: 27030  191.661837/s
vus_max....................: 200    min=200 max=200
```

##### Spike Testing (sudden spike to 1400 concurrent users from 100)

```
http_req_blocked...........: avg=45.25µs  min=0s   med=4µs   max=13.18ms  p(90)=8µs    p(95)=15µs
http_req_connecting........: avg=38.73µs  min=0s   med=0s    max=12.63ms  p(90)=0s     p(95)=0s
http_req_duration..........: avg=850.9µs  min=76µs med=570µs max=277.94ms p(90)=1.5ms  p(95)=1.92ms
http_req_receiving.........: avg=45.13µs  min=8µs  med=33µs  max=11.24ms  p(90)=70µs   p(95)=89µs
http_req_sending...........: avg=29.05µs  min=4µs  med=15µs  max=7.07ms   p(90)=36µs   p(95)=60µs
http_req_tls_handshaking...: avg=0s       min=0s   med=0s    max=0s       p(90)=0s     p(95)=0s
http_req_waiting...........: avg=776.71µs min=50µs med=490µs max=277.89ms p(90)=1.41ms p(95)=1.81ms
http_reqs..................: 169650 1466.844921/s
iteration_duration.........: avg=1s       min=1s   med=1s    max=1.27s    p(90)=1s     p(95)=1s
iterations.................: 56550  488.948307/s
vus_max....................: 1400   min=1400 max=1400
```
