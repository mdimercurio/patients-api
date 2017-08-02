FROM busybox

WORKDIR /app

ENV DATA_PATH /patients.json

COPY patients-api /app/patients-api

CMD ["/app/patients-api"]
