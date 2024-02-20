from kubernetes import client, config, watch

config.load_kube_config()
v1 = client.CoreV1Api()

print("Watching for Pod changes in the default namespace:")
w = watch.Watch()
for event in w.stream(v1.list_namespaced_pod, namespace='default'):
    print(f"Event: {event['type']} Pod: {event['object'].metadata.name}")
