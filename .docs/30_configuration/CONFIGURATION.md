<!--- HEADER --->
<div align="center">
   <img src="../../.media/assets/badges/assets_badges_project_backgroundless.png" width="15%" height="auto"/>
</div>

<!---
#####################################################
# Configuration - Helm
#####################################################
--->
### Configuration - Helm
<sup>[(Back to Readme)](../../README.md#configuration)</sup>
<br>
<!--- CONTENT --->

The Helm chart follows the <code>library</code> pattern, containing all reusable functions. All variables are parameterized, ensuring there are no hardcoded values or unnecessary volumes within the deployment.

- [Helm - `Library`](../../charts/kubeforge/templates/kubeforge/_helpers.tpl)
- [Helm - `Values`](../../charts/kubeforge/values.yaml)

##

<!---
#####################################################
# Configuration - Overlay
#####################################################
--->
### Configuration - Overlay
<sup>[(Back to Readme)](../../README.md#configuration)</sup>
<br>
<!--- CONTENT --->

The overlay configuration works in conjunction with the source configuration. The source configuration serves as the foundation, while the overlay configuration provides the final adjustments overriding or adding elements to the source configuration. This approach allows you to retain 99% of your original configuration and only modify the 1% that differs from the base.

<!-- list -->   
<ul>

   <!-- element[0] -->     
   <li>
   <details>
   <summary>$\color{#FAFAD2}{\textsf{Source}}$</summary>

   <ul>
   <!-- element [0][0] -->
   <li>
   <p>The source configuration defines a template to be used for the overlays.</p>
      
    # @kubernetes pod(s) configuration
    Pod:
    - metadata:
        name: bannana-pod 
      spec:
        containers:
        - name: bannana 
          command: [ "tail", "-f", "/dev/null" ]
    
    # @kubernetes pvc(s) configurations
    #PersistentVolumeClaim:
    
    # @kubernetes cm(s) configurations
    #ConfigMap:

    # @kubernetes pv(s) configurations
    #PersistentVolume:    
      
   </details>
   </li>

   <!-- element [1] -->      
   <li>
   <details>
   <summary>$\color{#EEE8AA}{\textsf{Overlay}}$</summary>

   <ul>
   <!-- element [1][0] -->     
   <li>
   <p>The overlay configuration overrides the source configuration using the <code>names</code> field as the reference point during the merging process.</p>
      
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
      
   <!-- element [1][1] -->  
   <li>
   <details>
   <summary>$\color{#EE82EE}{\textsf{Examples}}$</summary>
   <br>
   <p>Here are two examples of overlays, <code>banana</code> and <code>apple</code>, each utilizing the source configuration defined above. <br>The resulting outputs will be the pods named <code>mybanana-pod</code> and <code>myapple-pod</code></p>

   <ul>
   <!-- element [1][1][0] -->     
   <li>
   <details>
   <summary>$\color{#DA70D6}{\textsf{Bannana}}$</summary>
   <br>
      
    ---
    apiVersion: kubeforge.sh/v1
    kind: Overlay
    metadata:
      name: "bannana" 
      annotations:
        "helm.sh/hook": test
    spec:
      data:
    # @kubernetes pod(s) configurations
        Pod:
          - metadata:
              # name should match with sourceConfiguration
              name: bannana-pod 
              annotations:
                # "metal.io/override-name" annotation allows overriding the final Pod's name
                # during the manifest rendering phase. Since metadata.name is immutable and used
                # for merging configurations, the annotation is parsed by the rendering tool
                # to generate the final Pod name before deployment.
                kubeforge.sh/override-name: "mybannana-pod"
            spec:
              containers:
                - name: bannana 
                  image: busybox 
                  volumeMounts:
                    - name: bannana-pvc 
                      mountPath: /opt/config
                      subPath: config
              volumes:
                - name: bannana-pvc 
                  persistentVolumeClaim:
                    claimName: bannana-pvc 
    
    # @kubernetes pvc(s) configurations
        PersistentVolumeClaim:
          - metadata:
              name: bannana-pvc
              annotations:
                magic: "my-bannana-pvc"
            spec:
              accessModes:
                - ReadWriteMany
              resources:
                requests:
                  storage: 20Gi
              storageClassName: storage-local-retain
    
    # @kubernetes cm(s) configurations
        ConfigMap:
          - metadata:
              name: bannana-cm 
              annotations:
                magic: "my-bannana-cm"
            data:
              config: |
                lorem-ipsum
    ...
    
   </details>
   </li>  
   
   <!-- element [1][1][1] -->    
   <li>
   <details>
   <summary>$\color{#BA55D3}{\textsf{Apple}}$</summary>
   <br>
      
    ---
    apiVersion: kubeforge.sh/v1
    kind: Overlay
    metadata:
      name: "apple"
      annotations:
        "helm.sh/hook": test
    spec:
      data:
     
    # @kubernetes pod(s) configurations
        Pod:
          - metadata:
              # name should match with sourceConfiguration
              name: bannana-pod 
              annotations:
                # "metal.io/override-name" annotation allows overriding the final Pod's name
                # during the manifest rendering phase. Since metadata.name is immutable and used
                # for merging configurations, the annotation is parsed by the rendering tool
                # to generate the final Pod name before deployment.
                kubeforge.sh/override-name: "myapple-pod"
            spec:
              containers:
                - name: bannana 
                  image: busybox 
                  volumeMounts:
                    - name: apple-pvc 
                      mountPath: /opt/config
                      subPath: config
              volumes:
                - name: apple-pvc 
                  persistentVolumeClaim:
                    claimName: apple-pvc 
    
    # @kubernetes pvc(s) configurations
        PersistentVolumeClaim:
          - metadata:
              name: apple-pvc
              annotations:
                magic: "my-apple-pvc"
            spec:
              accessModes:
                - ReadWriteMany
              resources:
                requests:
                  storage: 20Gi
              storageClassName: storage-local-retain
    
    # @kubernetes cm(s) configurations
        ConfigMap:
          - metadata:
              name: apple-cm 
              annotations:
                magic: "my-apple-cm"
            data:
              config: |
                lorem-ipsum
    ...
    
   </details>
   </li>  
   
   </ul>
   
   </details>
   </li>
   
   </details>
   </li>
   
</ul>


<!--- FOOTER --->
<br>
<br>
<div align="center">
   <img src="../../.media/assets/badges/assets_badges_project_backgroundless.png" width="15%" height="auto"/>
</div>
