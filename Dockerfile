FROM centurylink/ca-certs

WORKDIR /drone-xmpp

COPY drone-xmpp /drone-xmpp/

ENTRYPOINT ["/drone-xmpp/drone-xmpp"]
