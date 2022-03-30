from flaskr.redis import RedisClient

redisClient = RedisClient(
    host="localhost",
    port=6379
    )


redisClient.set("Hello", "World")

redisClient.get("Hello")

redisClient.delete("Hello")
redisClient.get("Hello")

