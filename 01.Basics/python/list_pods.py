from kubernetes import client, config

# Load kube config from default location
config.load_kube_config()

v1 = client.CoreV1Api()
print("Listing pods in the default namespace:")
for pod in v1.list_namespaced_pod(namespace="default").items:
    print(f"{pod.metadata.name}\t{pod.status.phase}\t{pod.spec.containers[0].image}")
