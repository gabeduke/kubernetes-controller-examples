import time
from kubernetes import client, config, watch

def main():
    config.load_kube_config()
    v1 = client.CoreV1Api()
    w = watch.Watch()

    print("Starting custom controller to watch Pods in the default namespace...")
    while True:
        try:
            for event in w.stream(v1.list_namespaced_pod, namespace='default'):
                print(f"Event: {event['type']} Pod: {event['object'].metadata.name}")
        except Exception as e:
            print(f"Encountered exception: {e}")
            time.sleep(5)

if __name__ == '__main__':
    main()
