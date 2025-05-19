FROM ubuntu:latest

# Set the working directory
WORKDIR /app

# Copy requirements and app code
COPY requirements.txt /app
COPY devops /app

# Install dependencies and set up a virtual environment
RUN apt-get update && \
    apt-get install -y python3 python3-pip python3-venv && \
    python3 -m venv /venv && \
    /venv/bin/pip install -r requirements.txt

# Use the virtual environment
ENV PATH="/venv/bin:$PATH"

# Set the entrypoint and command
ENTRYPOINT ["python3"]
CMD ["manage.py", "runserver", "0.0.0.0:8000"]
