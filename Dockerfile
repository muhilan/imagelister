FROM busybox 

COPY imagelister /imagelister

CMD [ "/imagelister"]
