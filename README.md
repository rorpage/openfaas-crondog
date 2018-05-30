# crondog for OpenFaaS

Clone this repo and build a crondog container with the following:
```
docker build -t rorpage/crondog .
```

The Docker build will compile the crondog binary and copy it to an image for you.

Next, run the container and pass in 3 environment variables:
```
docker run -it --rm --net="host" \
  -e "cron_schedule=@every 1s" \
  -e "function_url=http://127.0.0.1:8080/function/func_wordcount" \
  -e "function_data=This is a super long test string woohoo" rorpage/crondog
```