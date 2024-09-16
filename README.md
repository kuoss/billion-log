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
a=1 b=1 uuid=f6d8cd6d-ff19-4380-920e-6f13bfce1e73
a=1 b=2 uuid=2262be90-3974-4f7f-8b86-896a9eda9d7b
a=1 b=3 uuid=9b06a1cf-97a3-4f07-bc8c-637550b187fe
...
a=1000 b=98 uuid=bf887df1-83cf-442a-a1ee-de4067709af3
a=1000 b=99 uuid=88044fe4-1638-471d-b00e-389cc0ac3a89
a=1000 b=100 uuid=b6401365-7e81-42a5-9ff0-4347efaf7dfe
```
