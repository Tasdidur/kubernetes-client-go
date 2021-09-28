#FROM ubuntu
#
#COPY ./out-app-1 ./out-app-1
#
#CMD ["./out-app-1"]

FROM ubuntu
COPY ./in-app-1 ./in-app-1
CMD ["./in-app-1"]
