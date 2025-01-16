<div align="center">

   <!-- logo -->
   <div style="width: 100%; height: auto; background-color: black;">
      <img src="./.media/assets/badges/assets_badges_project.png" width="100%" height="auto"/>      
   </div>
   <br>
  
   <!-- labels -->
   <img src="https://labl.es/svg?text=Dynamic%20Provisoring&width=200&bgcolor=a93226" align="center" style="margin: 5px"/>
   <img src="https://labl.es/svg?text=Kubernetes%20Controller&width=200&bgcolor=1e8449" align="center" style="margin: 5px"/>
   <img src="https://labl.es/svg?text=Helm%20Chart&width=200&bgcolor=154360" align="center" style="margin: 5px"/>
   
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

Lorem Ipsum is simply dummy text of the printing and typesetting industry.

##
<!---
#####################################################
# TL;DR
#####################################################
--->
<h3 id="tldr">
   $\large\color{Goldenrod}{\textbf{TL;DR}}$
</h3>

```sh

```

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
   <div style="width: 800; height: auto; background-color: black;">
      <img src="./.media/native/preview/gif/native_preview_gif_tldr.gif" width="800" height="auto"/>   
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
- [Usage](#usage)
- [Development](#development)
- [Miscellaneous](#miscellaneous)

<!---
$$\   $$\  $$$$$$\   $$$$$$\   $$$$$$\  $$$$$$$$\ 
$$ |  $$ |$$  __$$\ $$  __$$\ $$  __$$\ $$  _____|
$$ |  $$ |$$ /  \__|$$ /  $$ |$$ /  \__|$$ |      
$$ |  $$ |\$$$$$$\  $$$$$$$$ |$$ |$$$$\ $$$$$\    
$$ |  $$ | \____$$\ $$  __$$ |$$ |\_$$ |$$  __|   
$$ |  $$ |$$\   $$ |$$ |  $$ |$$ |  $$ |$$ |      
\$$$$$$  |\$$$$$$  |$$ |  $$ |\$$$$$$  |$$$$$$$$\ 
 \______/  \______/ \__|  \__| \______/ \________|
--->

## Usage
<sup>[(Back to Top)](#table-of-contents)</sup><br>

This section provides guidance on deploying and configuring streaming instances using Docker, Docker Compose, and Kubernetes (K8S) manifests. It includes specific instructions for different Linux distributions and GPU acceleration.

### Table Of Contents:
  - $\large\color{Goldenrod}{\textbf{Usage}}$
     - [Usage `Standalone`](./.docs/10_usage/USAGE.md#usage---docker) 
     - [Usage `Kubernetes`](./.docs/10_usage/USAGE.md#usage---docker-compose)

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
   <sup><code>Steam? No Problem!</code></sup>
   <br>
   <div style="width: 800; height: auto; background-color: black;">
      <img src="./.media/native/preview/gif/native_preview_gif_steam.gif" width="800" height="auto"/>
   </div>   
</div>

<!---
$$$$$$$\  $$$$$$$$\ $$\    $$\ $$$$$$$$\ $$\       $$$$$$\  $$$$$$$\  $$\      $$\ $$$$$$$$\ $$\   $$\ $$$$$$$$\ 
$$  __$$\ $$  _____|$$ |   $$ |$$  _____|$$ |     $$  __$$\ $$  __$$\ $$$\    $$$ |$$  _____|$$$\  $$ |\__$$  __|
$$ |  $$ |$$ |      $$ |   $$ |$$ |      $$ |     $$ /  $$ |$$ |  $$ |$$$$\  $$$$ |$$ |      $$$$\ $$ |   $$ |   
$$ |  $$ |$$$$$\    \$$\  $$  |$$$$$\    $$ |     $$ |  $$ |$$$$$$$  |$$\$$\$$ $$ |$$$$$\    $$ $$\$$ |   $$ |   
$$ |  $$ |$$  __|    \$$\$$  / $$  __|   $$ |     $$ |  $$ |$$  ____/ $$ \$$$  $$ |$$  __|   $$ \$$$$ |   $$ |   
$$ |  $$ |$$ |        \$$$  /  $$ |      $$ |     $$ |  $$ |$$ |      $$ |\$  /$$ |$$ |      $$ |\$$$ |   $$ |   
$$$$$$$  |$$$$$$$$\    \$  /   $$$$$$$$\ $$$$$$$$\ $$$$$$  |$$ |      $$ | \_/ $$ |$$$$$$$$\ $$ | \$$ |   $$ |   
\_______/ \________|    \_/    \________|\________|\______/ \__|      \__|     \__|\________|\__|  \__|   \__|
 --->
## Development
<sup>[(Back to top)](#table-of-contents)</sup>

<img src="./.media/assets/sections/assets_sections_e.png" align="left" width="5%" height="auto"/>

This section explains how we build our software, focusing on different structures like monolithic and distributed systems. You will also find information about our development workflows, including continuous integration and delivery.

### Table Of Contents:
  - $\large\color{Goldenrod}{\textbf{Development - Structure}}$
     - [Development - Structure - Monolithic](./.docs/40_development/structure/MONOLITHIC.md#development---structure---monolithic)
   <sup><img src="https://labl.es/svg?text=WIP&bgcolor=F7DC6F" align="center"/></sup>
     - Development - Structure - Distributed
   <sup><img src="https://labl.es/svg?text=WIP&bgcolor=F7DC6F" align="center"/></sup>
     - Development - Structure - Repository
   <sup><img src="https://labl.es/svg?text=WIP&bgcolor=F7DC6F" align="center"/></sup> 
  - $\large\color{Goldenrod}{\textbf{Development - Workflow}}$
     - Development - Workflow - CI
   <sup><img src="https://labl.es/svg?text=WIP&bgcolor=F7DC6F" align="center"/></sup>
     - Development - Workflow - CD
   <sup><img src="https://labl.es/svg?text=WIP&bgcolor=F7DC6F" align="center"/></sup>


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
