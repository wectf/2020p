import os
import socket
from func_timeout import func_timeout, FunctionTimedOut
from flask import Flask, request
from werkzeug.serving import WSGIRequestHandler
WSGIRequestHandler.protocol_version = "HTTP/1.1"

REDIS_ADDRESS = "127.0.0.1"
REDIS_PORT = 6379
NEWLINE = b'\x0d\x0a'
app = Flask(__name__)


def tcp_send(msg: bytes,
             redis_addr=REDIS_ADDRESS,
             redis_port=REDIS_PORT) -> [bytes]:
    # tcp init
    fd = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    fd.connect((redis_addr, redis_port))
    # send the message
    fd.sendall(msg + NEWLINE)
    # receive 2048 bytes, assume resp at most 2048 bytes
    # i'm lazy : )
    data = fd.recv(2048)
    # be a good programmer and clean up things
    fd.close()
    return data.split(NEWLINE)


# set timeout to be 1s, in case Redis hangs us
# also in case Redis is doing some weird things
def tcp_send_safe(msg: bytes, redis_addr=REDIS_ADDRESS, redis_port=REDIS_PORT) -> [bytes]:
    try:
        return func_timeout(1, tcp_send, args=(msg,), kwargs={
            "redis_addr": redis_addr,
            "redis_port": redis_port
        })
    except:
        return []


# execute GET [key]
def redis_get(key: str, redis_addr=REDIS_ADDRESS, redis_port=REDIS_PORT) -> str:
    result = tcp_send_safe(f"GET {key}".encode('ascii'),
                           redis_addr=redis_addr,
                           redis_port=redis_port)
    # redis response is $1[NEWLINE]balbalbla if found or $-1[NEWLINE] if not found
    if len(result) < 2 or result[1] == b'':
        return 'NO_SUCH_KEY'
    return result[1].decode('ascii')


# execute SET [key] [value]
def redis_set(key: str, value: str, redis_addr=REDIS_ADDRESS, redis_port=REDIS_PORT) -> bool:
    result = tcp_send_safe(f"SET {key} {value}".encode('ascii'),
                           redis_addr=redis_addr,
                           redis_port=redis_port)
    # redis response is +OK if success
    if len(result) < 1 or result[0] == b'':
        return False
    return result[0] == b'+OK'


# return a dict mapping keys redis_addr and redis_port to their values if included in request
def get_addr_port_helper():
    redis_addr = request.args.get('redis_addr', '')
    redis_port = request.args.get('redis_port', '')
    kwargs = {}
    if redis_addr:
        kwargs["redis_addr"] = redis_addr
    if redis_port:
        try:
            kwargs["redis_port"] = int(redis_port)
        except ValueError:
            pass
    return kwargs


@app.route('/')
def index():
    return 'Set a key value pair: PUT /set/[KEY]?value=[VALUE]<br>' \
           'Get a value by key: GET /get?key=[KEY]'


@app.route('/set/<key>', methods=['PUT'])
def set_kv(key):
    if redis_set(key, request.args.get('value', ''), **get_addr_port_helper()):
        return 'OK'
    return 'NOK'


@app.route('/get')
def get_v():
    return redis_get(request.args.get('key', ''), **get_addr_port_helper())


@app.route('/debug', methods=["POST"])
def debug():
    # only allow command execution if user is me : )
    if request.remote_addr != '127.0.0.1':
        return "Only localhost can execute command :("
    exec(request.form.get('cmd', 'nothing'))
    return "OK"


if __name__ == '__main__':
    app.run(port=80, host="0.0.0.0")
