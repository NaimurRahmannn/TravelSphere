# TravelSphere

## Run With Docker Desktop

Make sure Docker Desktop is running, then open PowerShell in the project folder:

```powershell
cd D:\TravelSphere
```

Build the Docker image:

```powershell
docker build -t travelsphere .
```

Run the container:

```powershell
docker run -d --name travelsphere -p 8080:8080 travelsphere
```

Open the app in your browser:

```text
http://localhost:8080
```

If port `8080` is already in use, remove the failed container and run it on another local port:

```powershell
docker rm travelsphere
docker run -d --name travelsphere -p 8081:8080 travelsphere
```

Then open:

```text
http://localhost:8081
```

Useful Docker commands:

```powershell
docker ps
docker logs -f travelsphere
docker stop travelsphere
docker rm travelsphere
```

The Dockerfile exposes port `8080` inside the container. The `-p 8081:8080` format means local port `8081` maps to container port `8080`.
