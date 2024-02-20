# Kubernetes Controllers Tutorial

Welcome to the Kubernetes Controllers Tutorial! This repository is designed to introduce you to the concepts and operation of Kubernetes controllers, using a hands-on approach. Whether you're familiar with Bash and Python or new to client-go, you'll find step-by-step guides to help you understand how controllers work in Kubernetes.

## Understanding Kubernetes Controllers

Kubernetes controllers are the brain behind Kubernetes' ability to manage the lifecycle of resources. They ensure the state of the cluster matches the desired state specified by the user. Here's a breakdown of the core components and concepts involved in a controller's operation:

### Core Components of a Kubernetes Controller

- **Queue**: A central component of a controller, the queue temporarily stores information about the resources that need processing. Controllers use queues to manage work in a reliable and efficient manner, ensuring that each item is processed while handling failures gracefully.

- **Informer**: Informers watch for changes to specific Kubernetes resources and update the controller about these changes. They play a crucial role in enabling controllers to react to changes in the cluster's state, such as creating, updating, or deleting resources.

- **Reconciler**: The reconciler is where the logic of a controller resides. It compares the actual state of the cluster to the desired state and performs actions to reconcile the two. This might involve creating, updating, or deleting resources to achieve the desired configuration.

- **Manager**: The manager is an overarching component that sets up the controller's various parts and starts the control loop. It typically handles the initialization of informers, queues, and the reconciler, orchestrating their operation to ensure the controller functions as intended.

### Key Concepts in Controller Design

- **State Management**: Controllers continuously monitor the state of resources in the Kubernetes cluster and take actions to ensure this state matches the desired configuration specified by users. This involves creating, updating, and deleting resources as necessary.

- **Desired vs. Actual State**: The desired state is the configuration specified by the user, typically through YAML files or other declarative means. The actual state is the current configuration of the cluster. Controllers work to minimize the difference between these two states.

- **Eventual Consistency**: Kubernetes controllers operate under the principle of eventual consistency. This means that while the actual state may not immediately match the desired state due to various factors (such as delays in resource provisioning or network issues), the controller will continue to attempt reconciling these states over time until they match.

- **Control Loop**: At the heart of a controller is the control loop. This continuous loop watches for changes in the cluster, triggers reconciliation when the actual state diverges from the desired state, and applies the necessary changes to bring the cluster's state in line with the user's specifications.

### The Lifecycle of a Controller

1. **Watch**: The controller uses informers to watch for changes to specific resources.
2. **Queue**: Upon detecting a change, the informer places an event in the queue.
3. **Reconcile**: The controller processes events from the queue, invoking the reconciler to adjust the cluster's state as needed.
4. **Repeat**: This process repeats continuously, ensuring the cluster maintains the desired state over time.

Controllers are a fundamental part of Kubernetes' self-healing mechanism, automatically managing the state of resources within the cluster. By understanding these core components and concepts, developers can better appreciate how Kubernetes maintains the desired state of applications and infrastructure, ensuring high availability and resilience.

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
