FROM busybox

WORKDIR /app

MKDIR /data

ENV DATA_PATH /data/patients.json

COPY patients-api /app/patients-api

CMD ["/app/patients-api"]
