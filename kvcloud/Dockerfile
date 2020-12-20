FROM ubuntu:20.04
RUN apt update
RUN apt install -y redis-server python3-pip
RUN echo 'rename-command CONFIG ""' >> /etc/redis/redis.conf
RUN echo 'rename-command DEBUG ""' >> /etc/redis/redis.conf
RUN echo 'rename-command SHUTDOWN ""' >> /etc/redis/redis.conf
RUN echo 'rename-command FLUSHDB ""' >> /etc/redis/redis.conf
RUN echo 'rename-command FLUSHALL ""' >> /etc/redis/redis.conf
RUN echo 'rename-command EVAL ""' >> /etc/redis/redis.conf
RUN echo 'rename-command SLAVEOF ""' >> /etc/redis/redis.conf
RUN echo 'rename-command MODULE ""' >> /etc/redis/redis.conf
RUN pip3 install func_timeout flask
RUN mkdir /app
COPY app.py /app
WORKDIR /app
RUN echo "we{00000000-0000-0000-0000-000000000000@demo-flag}" > /flag.txt
CMD service redis-server start && python3 app.py
