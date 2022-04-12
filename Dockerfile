FROM scratch
COPY /ui/css /css
COPY /ui/html html
COPY /ui/js js
COPY /linux /
CMD ["/brankad"]