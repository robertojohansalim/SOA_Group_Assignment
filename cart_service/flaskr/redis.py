import json
import os
import redis

class RedisClient:
    def __init__(self, host:str = "",port:int = 0, useInMemoryOnly:bool=False):
        self.useInMemoryOnly = useInMemoryOnly
        self.storage = {}

        if not self.useInMemoryOnly:
            if host == "" and port == 0:
                host = os.getenv("REDIS_HOST")
                port = os.getenv("REDIS_PORT")
            self.redis = redis.Redis(
                host=host,
                port=port
            )
            print("[INFO] redis connection established to {}:{}".format(host,port))
            

    def get(self, key:str):
        if self.useInMemoryOnly:
            try:
                return self.storage[key]
            except:
                return None

        redisValue = self.redis.get(key)
        if redisValue is None:
            return None

        return json.loads(redisValue)

    def set(self, key:str, content:object):
        if self.useInMemoryOnly:
            self.storage[key] = content

            print(self.storage)
            return


        jsonString = json.dumps(content)
        self.redis.set(key, jsonString)
    
    def delete(self, key:str):
        if self.useInMemoryOnly:
            del(self.storage[key])
            return

        self.redis.delete(key)
        


    
