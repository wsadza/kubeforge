<div align="center">

   <!-- logo -->
   <div style="width: 100%; height: auto; background-color: black;">
      <img src="./.media/assets/badges/assets_badges_project.png" width="100%" height="auto"/>      
   </div>
   <br>
  
   <!-- labels -->
   <img src="https://labl.es/svg?text=Resource%20Provisoring&width=200&bgcolor=a93226" align="center" style="margin: 5px"/>
   <img src="https://labl.es/svg?text=Kubernetes&width=200&bgcolor=1e8449" align="center" style="margin: 5px"/>
   <img src="https://labl.es/svg?text=Helm&width=200&bgcolor=154360" align="center" style="margin: 5px"/>

   <div align="center" style="display: flex; gap: 5px; justify-content: center;">
      <img src="https://labl.es/svg?text=Controller&width=200&bgcolor=a50068" align="center"/>
      <img src="https://labl.es/svg?text=Resource%20Management&width=200&bgcolor=d35400" align="center"/>
   </div>
   
</div>

<!---
$$\   $$\          $$\                  $$$$$$\                                         
$$ | $$  |         $$ |                $$  __$$\                                        
$$ |$$  /$$\   $$\ $$$$$$$\   $$$$$$\  $$ /  \__|$$$$$$\   $$$$$$\   $$$$$$\   $$$$$$\  
$$$$$  / $$ |  $$ |$$  __$$\ $$  __$$\ $$$$\    $$  __$$\ $$  __$$\ $$  __$$\ $$  __$$\ 
$$  $$<  $$ |  $$ |$$ |  $$ |$$$$$$$$ |$$  _|   $$ /  $$ |$$ |  \__|$$ /  $$ |$$$$$$$$ |
$$ |\$$\ $$ |  $$ |$$ |  $$ |$$   ____|$$ |     $$ |  $$ |$$ |      $$ |  $$ |$$   ____|
$$ | \$$\\$$$$$$  |$$$$$$$  |\$$$$$$$\ $$ |     \$$$$$$  |$$ |      \$$$$$$$ |\$$$$$$$\ 
\__|  \__|\______/ \_______/  \_______|\__|      \______/ \__|       \____$$ | \_______|
                                                                    $$\   $$ |          
                                                                    \$$$$$$  |          
                                                                     \______/                         
--->
# Kubeforge
<img src="./.media/assets/sections/assets_sections_a.png" align="left" width="5%" height="auto"/>

The Kubeforge is a Kubernetes-native solution that addresses the limitations of dynamic resource provisioning. It uses the Kubernetes controller pattern to merge user-defined Custom Resource Definitions (CRDs) with source pre-definied configuration, enabling consistent and automated resource provisioning.

> [!NOTE] 
> Imagine a scenario where each pod within your scope needs to be slightly different (e.g. difrent amount of resource; difrent runtime-class; difrent variables), you have the option to install them separately or use Kubeforge. With Kubeforge, you can define the [source configuration](charts/kubeforge/values.yaml#L107-L126) and only operates with [overlays](charts/kubeforge/templates/tests/bannana.yml#L14-L62) resources.

##
<!---
#####################################################
# TL;DR
#####################################################
--->
<h3 id="tldr">
   $\large\color{Goldenrod}{\textbf{TL;DR}}$
</h3>


<details>
<summary>$\color{#FAFAD2}{\textsf{Preparation}}$</summary>
<br>
<p>Add the Helm chart repository.</p>
   
    helm repo add kubeforge https://wsadza.github.io/kubeforge && helm repo update

</details>   

<details>
<summary>$\color{#EEE8AA}{\textsf{Installation}}$</summary>
<br>
<p>Install the <code>Kubeforge</code> Helm chart with a customized source configuration.</p>

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
    
</details>

<details>
<summary>$\color{#F0E68C}{\textsf{Usage}}$</summary>
<br>
<p>Create a <code>Kubeforge</code> overlay resource to provision the "banana-pod".</p>
   
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

</details>   


<!---
$$$$$$$\  $$$$$$$\  $$$$$$$$\ $$\    $$\ $$$$$$\ $$$$$$$$\ $$\      $$\ 
$$  __$$\ $$  __$$\ $$  _____|$$ |   $$ |\_$$  _|$$  _____|$$ | $\  $$ |
$$ |  $$ |$$ |  $$ |$$ |      $$ |   $$ |  $$ |  $$ |      $$ |$$$\ $$ |
$$$$$$$  |$$$$$$$  |$$$$$\    \$$\  $$  |  $$ |  $$$$$\    $$ $$ $$\$$ |
$$  ____/ $$  __$$< $$  __|    \$$\$$  /   $$ |  $$  __|   $$$$  _$$$$ |
$$ |      $$ |  $$ |$$ |        \$$$  /    $$ |  $$ |      $$$  / \$$$ |
$$ |      $$ |  $$ |$$$$$$$$\    \$  /   $$$$$$\ $$$$$$$$\ $$  /   \$$ |
\__|      \__|  \__|\________|    \_/    \______|\________|\__/     \__|
--->
## Preview
<div align="center">
   <sup><code>It was easy, right?</code></sup>
   <br>
   <br>
   <div style="width: 800; height: auto; background-color: black;">
   <img src="./.media/previews/previews_installation.gif" width="800" height="auto"/>
   </div>      
</div>


<!---
$$$$$$$$\  $$$$$$\   $$$$$$\  
\__$$  __|$$  __$$\ $$  __$$\ 
   $$ |   $$ /  $$ |$$ /  \__|
   $$ |   $$ |  $$ |$$ |      
   $$ |   $$ |  $$ |$$ |      
   $$ |   $$ |  $$ |$$ |  $$\ 
   $$ |    $$$$$$  |\$$$$$$  |
   \__|    \______/  \______/
--->
## Table Of Contents:
- [Installation](#installation)
- [Configuration](#configuration)
- [Miscellaneous](#miscellaneous)

<!---
$$$$$$\                       $$\               $$\ $$\            $$\     $$\                     
\_$$  _|                      $$ |              $$ |$$ |           $$ |    \__|                    
  $$ |  $$$$$$$\   $$$$$$$\ $$$$$$\    $$$$$$\  $$ |$$ | $$$$$$\ $$$$$$\   $$\  $$$$$$\  $$$$$$$\  
  $$ |  $$  __$$\ $$  _____|\_$$  _|   \____$$\ $$ |$$ | \____$$\\_$$  _|  $$ |$$  __$$\ $$  __$$\ 
  $$ |  $$ |  $$ |\$$$$$$\    $$ |     $$$$$$$ |$$ |$$ | $$$$$$$ | $$ |    $$ |$$ /  $$ |$$ |  $$ |
  $$ |  $$ |  $$ | \____$$\   $$ |$$\ $$  __$$ |$$ |$$ |$$  __$$ | $$ |$$\ $$ |$$ |  $$ |$$ |  $$ |
$$$$$$\ $$ |  $$ |$$$$$$$  |  \$$$$  |\$$$$$$$ |$$ |$$ |\$$$$$$$ | \$$$$  |$$ |\$$$$$$  |$$ |  $$ |
\______|\__|  \__|\_______/    \____/  \_______|\__|\__| \_______|  \____/ \__| \______/ \__|  \__|
--->

## Installation
<sup>[(Back to Top)](#table-of-contents)</sup><br>

This section offers instructions for installing a Kubeforge controller. Whether you're looking to test it or deploy it on your Kubernetes cluster, here are two methods you can follow.

### Table Of Contents:
  - $\large\color{Goldenrod}{\textbf{Installation}}$
     - [Installation `Standalone`](./.docs/10_installation/INSTALLATION.md#installation---docker) 
     - [Installation `Kubernetes`](./.docs/10_installation/INSTALLATION.md#installation---kubernetes)

<!---
$$$$$$$\  $$$$$$$\  $$$$$$$$\ $$\    $$\ $$$$$$\ $$$$$$$$\ $$\      $$\ 
$$  __$$\ $$  __$$\ $$  _____|$$ |   $$ |\_$$  _|$$  _____|$$ | $\  $$ |
$$ |  $$ |$$ |  $$ |$$ |      $$ |   $$ |  $$ |  $$ |      $$ |$$$\ $$ |
$$$$$$$  |$$$$$$$  |$$$$$\    \$$\  $$  |  $$ |  $$$$$\    $$ $$ $$\$$ |
$$  ____/ $$  __$$< $$  __|    \$$\$$  /   $$ |  $$  __|   $$$$  _$$$$ |
$$ |      $$ |  $$ |$$ |        \$$$  /    $$ |  $$ |      $$$  / \$$$ |
$$ |      $$ |  $$ |$$$$$$$$\    \$  /   $$$$$$\ $$$$$$$$\ $$  /   \$$ |
\__|      \__|  \__|\________|    \_/    \______|\________|\__/     \__|
--->
<h2>Preview</h2>
<div align="center">
   <sup><code>Take a piece of CRD, mix it with the flavor of CM, and voilà, you have your pod!</code></sup>
   <br>
   <br>
   <div style="width: 600; height: auto; background-color: black;">
      <img src="./.media/previews/previews_concept.png" align="center" width="600" height="auto"/>   
   </div>
</div>

<!---
 $$$$$$\   $$$$$$\  $$\   $$\ $$$$$$$$\ $$$$$$\  $$$$$$\  $$\   $$\ $$$$$$$\   $$$$$$\ $$$$$$$$\ $$$$$$\  $$$$$$\  $$\   $$\ 
$$  __$$\ $$  __$$\ $$$\  $$ |$$  _____|\_$$  _|$$  __$$\ $$ |  $$ |$$  __$$\ $$  __$$\\__$$  __|\_$$  _|$$  __$$\ $$$\  $$ |
$$ /  \__|$$ /  $$ |$$$$\ $$ |$$ |        $$ |  $$ /  \__|$$ |  $$ |$$ |  $$ |$$ /  $$ |  $$ |     $$ |  $$ /  $$ |$$$$\ $$ |
$$ |      $$ |  $$ |$$ $$\$$ |$$$$$\      $$ |  $$ |$$$$\ $$ |  $$ |$$$$$$$  |$$$$$$$$ |  $$ |     $$ |  $$ |  $$ |$$ $$\$$ |
$$ |      $$ |  $$ |$$ \$$$$ |$$  __|     $$ |  $$ |\_$$ |$$ |  $$ |$$  __$$< $$  __$$ |  $$ |     $$ |  $$ |  $$ |$$ \$$$$ |
$$ |  $$\ $$ |  $$ |$$ |\$$$ |$$ |        $$ |  $$ |  $$ |$$ |  $$ |$$ |  $$ |$$ |  $$ |  $$ |     $$ |  $$ |  $$ |$$ |\$$$ |
\$$$$$$  | $$$$$$  |$$ | \$$ |$$ |      $$$$$$\ \$$$$$$  |\$$$$$$  |$$ |  $$ |$$ |  $$ |  $$ |   $$$$$$\  $$$$$$  |$$ | \$$ |
 \______/  \______/ \__|  \__|\__|      \______| \______/  \______/ \__|  \__|\__|  \__|  \__|   \______| \______/ \__|  \__|
--->

## Configuration
<sup>[(Back to Top)](#table-of-contents)</sup><br>

<img src=".media/assets/sections/assets_sections_d.png" align="left" width="5%" height="auto"/>

It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).

### Table Of Contents:
  - $\large\color{Goldenrod}{\textbf{Configuration}}$
     - [Configuration - `Helm`](./.docs/30_configuration/CONFIGURATION.md#configuration---helm)
     - [Configuration - `Overlay`](./.docs/30_configuration/CONFIGURATION.md#configuration---overlay)

<!---
$$$$$$$\  $$$$$$$\  $$$$$$$$\ $$\    $$\ $$$$$$\ $$$$$$$$\ $$\      $$\ 
$$  __$$\ $$  __$$\ $$  _____|$$ |   $$ |\_$$  _|$$  _____|$$ | $\  $$ |
$$ |  $$ |$$ |  $$ |$$ |      $$ |   $$ |  $$ |  $$ |      $$ |$$$\ $$ |
$$$$$$$  |$$$$$$$  |$$$$$\    \$$\  $$  |  $$ |  $$$$$\    $$ $$ $$\$$ |
$$  ____/ $$  __$$< $$  __|    \$$\$$  /   $$ |  $$  __|   $$$$  _$$$$ |
$$ |      $$ |  $$ |$$ |        \$$$  /    $$ |  $$ |      $$$  / \$$$ |
$$ |      $$ |  $$ |$$$$$$$$\    \$  /   $$$$$$\ $$$$$$$$\ $$  /   \$$ |
\__|      \__|  \__|\________|    \_/    \______|\________|\__/     \__|
--->
<h2>Preview</h2>
<div align="center">
   <sup><code>Sequences! We love sequences, right?</code></sup>
   <br>
   <br>
   <div style="width: 600; height: auto; background-color: black;">
      <img src="./.media/previews/previews_sequence.png" align="center" width="600" height="auto"/>   
   </div>
</div>

<!---
$$\      $$\ $$$$$$\  $$$$$$\   $$$$$$\  
$$$\    $$$ |\_$$  _|$$  __$$\ $$  __$$\ 
$$$$\  $$$$ |  $$ |  $$ /  \__|$$ /  \__|
$$\$$\$$ $$ |  $$ |  \$$$$$$\  $$ |      
$$ \$$$  $$ |  $$ |   \____$$\ $$ |      
$$ |\$  /$$ |  $$ |  $$\   $$ |$$ |  $$\ 
$$ | \_/ $$ |$$$$$$\ \$$$$$$  |\$$$$$$  |
\__|     \__|\______| \______/  \______/
--->
## Miscellaneous
<sup>[(Back to top)](#table-of-contents)</sup>

<img src="./.media/assets/sections/assets_sections_f.png" align="left" width="5%" height="auto"/>

The "Miscellaneous" section gathers various resources and content that may not belong to a specific category but are still valuable and worth referencing. It's a place for extra tools, tips, and information that support a wide range of needs.

### Table Of Contents:
- $\large\color{Goldenrod}{\textbf{Helpful Resources}}$
   - [Helpful Resources - Setup](./.docs/50_miscellaneous/MISCELLANEOUS.md#helpful-resources---setup)
   - [Helpful Resources - Questions / Answers](./.docs/50_miscellaneous/MISCELLANEOUS.md#helpful-resources---questions---answers)
- [Document Template](./.docs/50_miscellaneous/DOCUMENT_TEMPLATE.md)


<br>
<br>
<div align="center">
   <img src="./.media/assets/badges/assets_badges_project_backgroundless.png" width="15%" height="auto"/>
</div>
