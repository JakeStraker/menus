FROM alpine
ADD bettermenus /bettermenus
RUN mkdir /Templates
ADD index.html /Templates/index.html
RUN chmod +x /bettermenus
EXPOSE 1323
ENTRYPOINT /bettermenus
