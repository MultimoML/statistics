# Statistics

Microservice for managing grocery items.

Available endpoints:
- [`/live`](https://multimo.ml/stats/live): Liveliness check
- [`/ready`](https://multimo.ml/stats/ready): Readiness check
- [`/all`](https://multimo.ml/stats/all): Returns all statistics

Branches:
- [`main`](https://github.com/MultimoML/statistics/tree/main): Contains latest development version
- [`prod`](https://github.com/MultimoML/statistics/tree/prod): Contains stable, tagged releases

## Setup/installation

Prerequisites:
- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)

Example usage:
- See all available options: `make help`
- Run microservice in a container: `make run`
- Release a new version: `make release ver=x.y.z`

All work should be done on `main`, `prod` should never be checked out or manually edited.
When releasing, the changes are merged into `prod` and both branches are pushed.
A GitHub Action workflow will then build and publish the image to GHCR, and deploy it to Kubernetes.

## License

Multimo is licensed under the [GNU AGPLv3 license](LICENSE).