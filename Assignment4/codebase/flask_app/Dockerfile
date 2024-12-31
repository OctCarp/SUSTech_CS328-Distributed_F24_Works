FROM python:3.13.1-slim

ARG VERSION=1.0.1

WORKDIR /app

COPY requirements.txt ./${VERSION}/app.py ./
RUN pip3 install --no-cache-dir -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

EXPOSE 5000

CMD [ "python", "app.py" ]
