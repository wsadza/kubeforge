<div align="center">
   <img src="../../.media/assets/badges/assets_badges_project_backgroundless.png" width="15%" height="auto"/>
</div>

<!---
#####################################################
# Helpful Resources
#####################################################
--->
## Helpful Resources

<!--- CONTENT --->

<!---
#####################################################
# Helpful Resources - Controller Patterns
#####################################################
--->
<h3 id="helpful-resources---controller---patterns">
$\color{yellow}{\textsf{Helpful Resources - Controller Patterns}}$    
</h3>

<sup>[(Back to Miscellaneous)](../../README.md#miscellaneous)</sup>
<br>

<!--- CONTENT --->

#### Controller Patterns

> [!TIP]
> <ul>
>    <li><a href="https://github.com/kubernetes/sample-controller">Simple Kubernetes Controller</a></li>
>    <li><a href="https://s3.amazonaws.com/ebooks.syncfusion.com/LiveReadOnlineFiles/kubernetes-succinctly/Images/pod-creation-sequence-source.png">Sequence of Pod creation example.</a></li>  
> </ul>
> </table>

##

<!---
#####################################################
# Helpful Resources - TODO List
#####################################################
--->
<h3 id="helpful-resources---todo---list">
$\color{yellow}{\textsf{Helpful Resources - TODO List }}$    
</h3>

<sup>[(Back to Miscellaneous)](../../README.md#miscellaneous)</sup>
<br>
<!--- CONTENT --->

#### Question: Why don't we need to install the NVIDIA driver in the finall container?
<details>
<summary>Answer</summary>
<br>
<!--- ANSWER --->
   
The NVIDIA driver toolkit ships all the necessary libraries to the final container.

<!--- ANSWER --->
</details>

##

#### Question: Why don't we need to use VirtualGL?
<details>
<summary>Answer</summary>
<br>
<!--- ANSWER --->   
   
Think of containers like laptops. On laptops, we often deal with NVIDIA Prime and Bumblebee to manage GPU usage. However, in containers, the primary GPU is usually LLVMPipe, instead of the Intel integrated GPU. To ensure we use the NVIDIA GPU, we can set the following environment variables:
<br><br>
   
   - `__NV_PRIME_RENDER_OFFLOAD=1`
   - `__GLX_VENDOR_LIBRARY_NAME=NVIDIA`
   
   These settings force graphic applications to utilize the NVIDIA GPU. For more details, you can refer to the official NVIDIA documentation [here](https://download.nvidia.com/XFree86/Linux-x86_64/435.17/README/primerenderoffload.html).
   
<!--- ANSWER --->
</details>

##

#### Question: Why can't I fully utilize WSL?
<details>
<summary>Answer</summary>
<br>
<!--- ANSWER --->      
   
Unfortunately, Microsoft started implementing [D3D12](https://www.phoronix.com/news/Mesa-24.1-Zink-D3D12-Default) in Mesa around year or two ago? - which limits our ability to use all Vulkan functions, resulting in significantly poor performance.

<!--- ANSWER --->   
</details>

##

#### Question: Are there other solutions to share GPU?
<details>
<summary>Answer</summary>
<br>
<!--- ANSWER --->   
   
- **Intel**: Offers various solutions for virtual GPU sharing, particularly through their integrated graphics.
- **AMD**: Provides solutions like MxGPU, which allows multiple virtual machines to share a single GPU.
- **Dual-Coder Crack for RTX Series 20***: Enables the use of NVIDIA's vGPU technology on consumer-grade RTX 20 series cards.
- **Older NVIDIA GPUs**: Models like the K2 support vGPU technology, which can be leveraged for virtualization and resource sharing.
- **Fractional-GPU by ClearML**: This appears to be the first Docker-related project to offer a sharing mechanism publicly.
   - <details>
     <summary>$\color{green}{\textsf{Details}}$</summary>
      
     - The project utilizes a customized libnvidia-ml.so.1, which includes non-NVIDIA functions such as get_max_gpu_mem.
     - They effectively reduce the size of this library by half (from 2MB original to 0.9MB after modification).
     - The get_max_gpu_mem function is located at f30f1efab800000080c3.
       
     </details>
   
Some resources:
- [Intel vGPU KubeVirt](https://kubevirt.io/2021/intel-vgpu-kubevirt.html)
- [NVIDIA vGPU on Proxmox VE](https://pve.proxmox.com/wiki/NVIDIA_vGPU_on_Proxmox_VE)
- [DualCoder vGPU Unlock](https://github.com/DualCoder/vgpu_unlock)
- [Using vGPU Unlock with Proxmox 7](https://www.michaelstinkerings.org/using-vgpu-unlock-with-proxmox-7/)
- [vgpu-proxmox GitLab](https://gitlab.com/polloloco/vgpu-proxmox)
- [clearml-fractional-gpu](https://github.com/allegroai/clearml-fractional-gpu)
- [HAMi](https://github.com/Project-HAMi/HAMi)
- [gpu-partitioning-guide](https://github.com/rh-aiservices-bu/gpu-partitioning-guide)
- [mortalgpu](https://gitlab.com/MaxIV/kubernetes/mortalgpu)
  
<!--- ANSWER --->  
</details>


##

#### Question: Are there plans to configure and discover solutions for Intel/AMD?
<details>
<summary>Answer</summary>
<br>
<!--- ANSWER --->   
   
It could happen in the future — who knows? But for now, NVIDIA is the dominant player in the market. Honestly, if you're looking to use your PC GPU unit for computing, you're pretty much compelled to go with NVIDIA

<!--- ANSWER ---> 
</details>

<br>
<br>
<div align="center">
   <img src="../../.media/assets/badges/assets_badges_project_backgroundless.png" width="15%" height="auto"/>
</div>
