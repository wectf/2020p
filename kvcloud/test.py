from requests import *
import urllib
import random
import string


def check(host="172.17.0.6"):
    port = 80
    letters = string.ascii_lowercase

    def setget():
        cmd = ''.join(random.choice(letters) for i in range(40))
        if put(f"http://{host}:%s/set/%s?value=dev_test" % (port, cmd)).content != b'OK':
            print("Cannot set")
        else:
            print("OK")
        if get(f"http://{host}:%s/get?key=%s" % (port, cmd)).content != b'dev_test':
            print("Cannot get")
        else:
            print("OK")

    def auth():
        if post(f"http://{host}:%s/debug" % port).content != b"Only localhost can execute command :(":
            print("Auth bypass")
        else:
            print("OK")

        if post(f"http://{host}:%s/debug" % port, headers={
            "X-Forwarded-For": "127.0.0.1"
        }).content != b"Only localhost can execute command :(":
            print("Auth bypass with XFF")
        else:
            print("OK")

    setget()
    auth()

check()

