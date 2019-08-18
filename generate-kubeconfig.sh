#!/bin/bash
set -eu -o pipefail

# Set OS specific values.
if [[ "$OSTYPE" == "linux-gnu" ]]; then
    REQUEST_ID=$(uuid)
    BASE64_DECODE_FLAG="-d"
    BASE64_WRAP_FLAG="-w 0"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    REQUEST_ID=$(uuidgen)
    BASE64_DECODE_FLAG="-D"
    BASE64_WRAP_FLAG=""
else
    echo "Unknown OS ${OSTYPE}"
    exit 1
fi

mkdir -p build
pushd build
cat > csr <<EOF
{
  "hosts": [
  ],
  "CN": "k8smanager",
  "names": [{
        "O": "system:masters"
    }],
  "key": {
    "algo": "ecdsa",
    "size": 256
  }
}
EOF

cat csr | cfssl genkey - | cfssljson -bare server

# Create Kubernetes CSR
cat <<EOF | kubectl create -f -
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: ${REQUEST_ID}
spec:
  groups:
  - system:authenticated
  request: $(cat server.csr | base64 | tr -d '\n')
  usages:
  - digital signature
  - key encipherment
  - client auth
EOF
kubectl certificate approve ${REQUEST_ID}

kubectl get csr ${REQUEST_ID} -o jsonpath='{.status.certificate}' \
    | base64 ${BASE64_DECODE_FLAG} > server.crt

kubectl -n kube-system exec $(kubectl get pods -n kube-system -l k8s-app=kube-dns  -o jsonpath='{.items[0].metadata.name}') -c kubedns -- /bin/cat /var/run/secrets/kubernetes.io/serviceaccount/ca.crt > ca.crt

# Extract cluster IP from the current context
CURRENT_CONTEXT=$(kubectl config current-context)
CURRENT_CLUSTER=$(kubectl config view -o jsonpath="{.contexts[?(@.name == \"${CURRENT_CONTEXT}\"})].context.cluster}")
CURRENT_CLUSTER_ADDR=$(kubectl config view -o jsonpath="{.clusters[?(@.name == \"${CURRENT_CLUSTER}\"})].cluster.server}")

cat > kubeconfig <<EOF
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: $(cat ca.crt | base64 ${BASE64_WRAP_FLAG})
    server: ${CURRENT_CLUSTER_ADDR}
  name: cluster
contexts:
- context:
    cluster: cluster
    user: k8smanager
  name: cluster
current-context: cluster
kind: Config
preferences: {}
users:
- name: k8smanager
  user:
    client-certificate-data: $(cat server.crt | base64 ${BASE64_WRAP_FLAG})
    client-key-data: $(cat server-key.pem | base64 ${BASE64_WRAP_FLAG})
EOF

popd