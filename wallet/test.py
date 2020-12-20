from requests import *

def check(HOST="172.17.0.2:3000"):
    s = Session()
    def pwn():
        if b'html' in s.get(f"http://{HOST}/").content:
            print("OK")
        else:
            print("Cannot visit HTML")
        if b'html' in s.post(f"http://{HOST}/style", data={"style": "\"Raw"}).content:
            print("HTML in raw")
        else:
            print("OK")
        if b'style="Raw,addresses=Unknown Address: x"' == s.post(f"http://{HOST}/address", data={"address": "x\""}).content:
            print("OK")
        else:
            print("Cannot pwn")
    pwn()

check()