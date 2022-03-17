#`production` image used for runtime
FROM python-base as production
ENV FASTAPI_ENV=production
COPY --from=builder-base $PYSETUP_PATH $PYSETUP_PATH
COPY ./app /app/
WORKDIR /app
CMD ["gunicorn", "-k", "uvicorn.workers.UvicornWorker", "main:app"]