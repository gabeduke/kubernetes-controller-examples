# Kubernetes Controllers Tutorial

Welcome to the Kubernetes Controllers Tutorial! This repository is designed to introduce you to the concepts and operation of Kubernetes controllers, using a hands-on approach. Whether you're familiar with Bash and Python or new to client-go, you'll find step-by-step guides to help you understand how controllers work in Kubernetes.

## Repository Structure

This repository is organized into three main sections, each catering to different levels of complexity and understanding:

- **Basics**: Start here if you're new to Kubernetes or want a refresher on some of the foundational concepts. You'll learn how to interact with Kubernetes using Bash, Python, and Go, focusing on listing Pods within a namespace.

- **Intermediate**: This section introduces the concept of watching for changes in Kubernetes resources. You'll explore how to implement watchers in Bash, Python, and Go, enhancing your understanding of dynamic resource management in Kubernetes.

- **Advanced**: Dive deep into the world of Kubernetes controllers. Here, you'll learn how to build custom controllers in Python and Go, managing and reacting to changes in Kubernetes resources programmatically.

### Getting Started

Before you dive into the tutorials, make sure you have the following prerequisites installed and configured:

- **Kubernetes Cluster**: You'll need access to a Kubernetes cluster. You can set one up locally using [Minikube](https://minikube.sigs.k8s.io/docs/start/) or [kind](https://kind.sigs.k8s.io/), or use a managed Kubernetes service provided by cloud providers like AWS, Google Cloud, or Azure.
- **kubectl**: The Kubernetes command-line tool, `kubectl`, is essential for interacting with your cluster. [Install kubectl](https://kubernetes.io/docs/tasks/tools/) following the official Kubernetes documentation.
- **Python 3**: For the Python examples, ensure you have Python 3 installed. You might also need to install the Kubernetes Python client using `pip install kubernetes`. The `pipenv` file is included in the Python examples to help you set up a virtual environment.
- **Go**: For the Go examples, you'll need Go installed on your machine. Additionally, install the client-go library with `go get k8s.io/client-go@latest`. A `go.mod` file is included in the Go examples to help you manage dependencies.

### Additional Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/home/): The official documentation for Kubernetes is an excellent resource for understanding concepts and finding additional details.
- [Client-go Documentation](https://pkg.go.dev/k8s.io/client-go): For more information on using the Go client for Kubernetes.
- [Kubernetes Python Client](https://github.com/kubernetes-client/python): The GitHub repository for the Kubernetes Python client, containing examples and API documentation.
