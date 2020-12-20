import subprocess
from requests import get
import re
def check(HOST="172.17.0.5:8080"):
    def recaptcha():
        if get(f"http://{HOST}/insert?value=123").content == b'Incorrect Recaptcha':
            print("OK")
        else:
            print("Recaptcha not OK")

    def pwn():
        resp = get(f"http://{HOST}/").content
        ts = int(re.findall(b"Table recreated at: (.+?)</", resp)[0])
        x = subprocess.run(["go", "run", "test/main.go", str(ts)], stdout=subprocess.PIPE, text=True)
        sol = x.stdout.replace("[", "").replace("]", "").replace("\n", "").split(" ")
        res = ""
        for i in sol:
            res = get(f"http://{HOST}/3ce979d9-602d-4c9c-b713-e3183a3ec252?value={i}").content
        if b'we{' in res:
            print("OK")
        else:
            print("Cannot get flag") 
    recaptcha()
    pwn()
check()
