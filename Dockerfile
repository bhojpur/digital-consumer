FROM moby/buildkit:v0.9.3
WORKDIR /consumer
COPY consumer README.md /consumer/
ENV PATH=/consumer:$PATH
ENTRYPOINT [ "/bhojpur/consumer" ]