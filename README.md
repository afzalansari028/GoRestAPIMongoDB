# gomakeme

```bash
# build the image
docker build -t gofiber .

# run and publish with the name of gofiber
docker run --publish 5011:5011 --name gofiber gofiber

# stop
docker stop gofiber

# remove
docker image rm gofiber
```