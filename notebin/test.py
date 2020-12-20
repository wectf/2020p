from requests import get, post

def check(HOST="172.17.0.5:8080"):
    u = post(f"http://{HOST}/note", data={
        "content": "123",
        "title": "234"
    }).url
    u = u.split("#")
    if len(u) < 2:
        print("No redirect")
        return
    else:
        print("OK")
    res = get(f"http://{HOST}/note/{u[1]}").json()
    if res["content"] == "123" and res["title"]== "234":
        print("OK")
    else:
        print("Wrong resp")

check()