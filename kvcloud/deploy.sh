docker build . -f Dockerfile.deploy -t docker.pkg.github.com/wectf/2020p/kvcloud
docker push docker.pkg.github.com/wectf/2020p/kvcloud
