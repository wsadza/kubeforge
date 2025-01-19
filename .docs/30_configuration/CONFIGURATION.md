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

It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).

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

It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).


<!-- list -->   
<ul>

   <!-- element[0] -->     
   <li>
   <details>
   <summary>$\color{#FAFAD2}{\textsf{Source}}$</summary>

   <ul>
   <!-- element [0][0] -->
   <li>
   <p>Prepare the <code>Kubeforge</code> source configuration as a foundation for the next steps.</p>
      
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
      
   </details>
   </li>

   <!-- element [1] -->      
   <li>
   <details>
   <summary>$\color{#FAFAD2}{\textsf{Overlay}}$</summary>

   <ul>
   <!-- element [1][0] -->     
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
      
   <!-- element [1][1] -->  
   <li>
   <details>
   <summary>Examples</summary>
   <br>
   <p>It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.</p>

   <ul>
   <!-- element [1][1][0] -->     
   <li>
   <details>
   <summary>$\color{#FAFAD2}{\textsf{Bannana}}$</summary>
   <br>
      
      dd
      dd
    
   </details>
   </li>  
   
   <!-- element [1][1][1] -->    
   <li>
   <details>
   <summary>$\color{#FAFAD2}{\textsf{Apple}}$</summary>
   <br>
      
      dd
      dd
    
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
