FROM nginx:alpine

COPY nginx.conf /etc/nginx/nginx.conf

WORKDIR /app

COPY myapp /app/myapp

RUN chmod +x /app/myapp

EXPOSE 80

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]