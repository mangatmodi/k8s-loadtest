# User will run load test with master having number of users, hatch rate and outputfile
FROM python:3.7

#install locust
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y build-essential libncursesw5-dev libreadline-dev libssl-dev libgdbm-dev libc6-dev libsqlite3-dev libxml2-dev libxslt-dev python python-dev python-setuptools && apt-get clean
RUN pip install locustio pyzmq

COPY run.sh /run.sh
COPY ./locustfile.py /locustfile.py
RUN chmod 755 /run.sh

EXPOSE 5557 5558 8089
ENTRYPOINT ["/run.sh"]
