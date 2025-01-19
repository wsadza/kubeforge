<div align="center">
   <img src="../../.media/assets/badges/assets_badges_project_backgroundless.png" width="15%" height="auto"/>
</div>

<!---
#####################################################
# Installation - Standalone
#####################################################
--->
### Installation - Standalone
<sup>[(Back to Readme)](../../README.md#installation)</sup>
<br>
<!--- CONTENT --->

If you want to test <code>Kubeforge</code> without installing it directly on the Kubernetes cluster, you can follow the steps below to run it as a standalone Docker instance.

<!-- list -->   
<ul>

   <!-- element [0] -->    
   <li>
   <details>
   <summary>$\color{#FAFAD2}{\textsf{Preparation}}$</summary>
   <ul>
   <li>
   <p>Prepare the <code>Kubeforge</code> source configuration as a foundation for the next steps.</p>
      
    cat <<EOF > "${PWD}/sourceConfiguration.yml"
    Pod:
    - metadata:
        name: bannana-pod 
      spec:
        containers:
        - name: bannana 
          command: [ "tail", "-f", "/dev/null" ]
    EOF
      
   </li>
   <li>
   <p>Install <code>Kubeforge</code> custom resource definition.</p>
   
    ---
    apiVersion: apiextensions.k8s.io/v1
    kind: CustomResourceDefinition
    metadata:
      name: overlays.kubeforge.sh
    spec:
      group: kubeforge.sh 
      versions:
        - name: v1
          served: true
          storage: true
          schema:
    
            # schema used for validation
            openAPIV3Schema:
              type: object
              properties:
                spec:
                  type: object
                  # Allows any arbitrary structure under `spec` by omitting "properties"
                  # and adding the "x-kubernetes-preserve-unknown-fields" flag                
                  x-kubernetes-preserve-unknown-fields: true
                status:
                  type: object
                  properties:
                    data:
                      type: object
                      x-kubernetes-preserve-unknown-fields: true
          subresources:
            status: {}
      names:
        kind: Overlay 
        plural: overlays 
      scope: Namespaced
    ...
   </li>
   </ul>
   </details>
   </li>   

   <!-- element [1] -->    
   <li>
   <details>
   <summary>$\color{#EEE8AA}{\textsf{Installation}}$</summary>
   <ul>
   <li>
   <p>Execute the <code>Kubeforge</code> docker container with mounted kubeconfig and source configuration.</p>
      
    docker run \
       --volume "${HOME}/.kube/config:/opt/.kube/config" \
       --volume "${PWD}/sourceConfiguration.yml:/opt/sourceConfiguration.yml"
       --environment KUBEFORGE_KUBERNETES_CONFIG=/opt/.kube/config \
       --environment KUBEFORGE_SOURCE_CONFIGURATION=/opt/sourceConfiguration.yml \
    ghcr.io/wsadza/kubeforge 

   </ul>
   </details>
   </li>
   <!-- element [1] --> 
   
   <!-- element [2] -->    
   <li>
   <details>
   <summary>$\color{#F0E68C}{\textsf{Usage}}$</summary>
   <ul>
   <li>
   <p>Create a <code>Kubeforge</code> overlay resource to provision the "banana-pod"</p>
    
    cat <<EOF | kubectl apply -f -
    apiVersion: kubeforge.sh/v1
    kind: Overlay
    metadata:
      name: "bannana" 
    spec:
      data:
        Pod:
          - metadata:
              name: bannana-pod 
            spec:
              containers:
              - name: bannana 
                image: busybox 
    EOF
    
   </li>
   </ul>
   </details>
   </li>
   <!-- element [2] --> 

</ul>

##

<!---
#####################################################
# Installation - Kubernetes
#####################################################
--->
### Installation - Kubernetes
<sup>[(Back to Readme)](../../README.md#installation)</sup>
<br>
<!--- CONTENT --->

The default method of working with <code>Kubeforge</code> is to install it directly on a Kubernetes cluster - installation is streamlined through Helm chart templates.

<!-- list -->   
<ul>

   <!-- element [0] -->    
   <li>
   <details>
   <summary>$\color{#FAFAD2}{\textsf{Preparation}}$</summary>
   <ul>
   <li>
   <p>Add the Helm chart repository.</p>
   
    helm repo add kubeforge https://wsadza.github.io/kubeforge && helm repo update

   </details>
   </li>   

   <!-- element [1] -->    
   <li>
   <details>
   <summary>$\color{#EEE8AA}{\textsf{Installation}}$</summary>
   <ul>
   <li>
   <p>Install the <code>Kubeforge</code> Helm chart with a customized source configuration</p>

    cat <<EOF | helm install kubeforge kubeforge/kubeforge -f -
    kubeforge:
      sourceConfiguration:
        Pod:
        - metadata:
            name: bannana-pod 
          spec:
            containers:
              - name: bannana 
                command: [ "tail", "-f", "/dev/null" ]
    EOF

   </ul>
   </details>
   </li>
   <!-- element [1] --> 
   
   <!-- element [2] -->    
   <li>
   <details>
   <summary>$\color{#F0E68C}{\textsf{Usage}}$</summary>
   <ul>
   <li>
   <p>Create a <code>Kubeforge</code> overlay resource to provision the "banana-pod"</p>
   
    cat <<EOF | kubectl apply -f -
    apiVersion: kubeforge.sh/v1
    kind: Overlay
    metadata:
      name: "bannana" 
    spec:
      data:
        Pod:
          - metadata:
              name: bannana-pod 
            spec:
              containers:
              - name: bannana 
                image: busybox 
    EOF
    
   </li>
   </ul>
   </details>
   </li>
   <!-- element [2] --> 

</ul>

<br>
<br>
<div align="center">
   <img src="../../.media/assets/badges/assets_badges_project_backgroundless.png" width="15%" height="auto"/>
</div>
