# codebase/python/logging_service/Dockerfile

# Use Python 3.11 as the base image
FROM python:3.11

WORKDIR /app

# Copy all files from current directory to container
COPY . .

# Install Python dependencies without caching pip packages
RUN pip install --no-cache-dir -r requirements.txt

# Expose port for the service
EXPOSE 50052

# Start the logging service script
CMD ["python", "local_publisher.py"]