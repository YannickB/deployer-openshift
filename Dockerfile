FROM openwhisk/dockerskeleton

ENV GOPATH /opt
RUN apk add --update wget ca-certificates
RUN mkdir /.kube
RUN chmod -R 777 /.kube

RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub \
    && wget -q https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.27-r0/glibc-2.27-r0.apk -O /tmp/glibc.apk \
    && apk add /tmp/glibc.apk \
    && rm /tmp/glibc.apk

RUN wget -q https://github.com/openshift/origin/releases/download/v3.7.2/openshift-origin-client-tools-v3.7.2-282e43f-linux-64bit.tar.gz -O /tmp/oc.tar.gz \
    && tar -C /tmp -xzf /tmp/oc.tar.gz \
    && mv /tmp/openshift-origin-client-tools-v3.7.2-282e43f-linux-64bit/oc /usr/bin/oc \
    && rm -rf /tmp/oc.tar.gz /tmp/openshift-origin-client-tools-v3.7.2-282e43f-linux-64bit

ADD src /opt/src
# RUN apk add --no-cache --virtual .bootstrap-deps go \
RUN wget -q https://dl.google.com/go/go1.9.2.linux-amd64.tar.gz -O /tmp/go.tar.gz \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \
    && wget -q https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 -O /usr/bin/dep \
    && chmod +x /usr/bin/dep \
    && cd /opt/src \
    && /usr/local/go/bin/go build -o /action/exec \
    && rm -rf /tmp/go.tar.gz /usr/local/go /usr/bin/dep
    # && apk del .bootstrap-deps
