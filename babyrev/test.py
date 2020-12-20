import requests

def check(HOST="172.17.0.2"):

    def pre():
        if b"we{" in (requests.get(f"http://{HOST}/").content):
            print("Flag appear directly")
        else:
            print("OK")

    def get_flag():
        if b"we{" not in (requests.get(f"http://{HOST}/", headers={
            "User-Agent": "Flag Viewer 2.0"
        }).content):
            print("Cannot get flag")
        else:
            print("OK")
