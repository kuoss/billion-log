# log-loop

`log-loop` is a utility that generates nested loops based on environment variables and outputs indexed values along with UUIDs.

## Docker Image

The Docker image is available at:

`ghcr.io/kuoss/log-loop:latest`

## Environment Variables

- **LOOP_COUNTS**: Comma-separated values indicating the number of iterations for each loop depth.
  - Example: `LOOP_COUNTS=10,100,1000` will generate three loops: the outer loop running 10 times, the middle loop running 100 times, and the inner loop running 1000 times.

If `LOOP_COUNTS` is not provided, the default loop counts are `[1000, 1000]`.

## Build and Run Instructions

### Build the Docker Image

```bash
docker build -t ghcr.io/kuoss/log-loop:latest .
```

### Run the Application with Docker

```bash
docker run --rm -e LOOP_COUNTS="1000,100" ghcr.io/kuoss/log-loop:latest
```

### Example Output
```
a=000 b=000 uuid=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
a=000 b=001 uuid=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
...
a=999 b=099 uuid=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```
