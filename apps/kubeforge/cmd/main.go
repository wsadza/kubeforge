package main

import (
	"fmt"
	"kubeforge/internal/k8s/controller"
	"kubeforge/pkg/signals"
	"net/http"
	"sync"

  "github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
)

// Health state variables
var (
	mu         sync.RWMutex
	isReady    = false 
  isHealthy  = false
)

// HTTP readyz / healthz server check
func startHealthCheckServer(serverPort string) {

  // Serve the /readyz endpoint
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		defer mu.RUnlock()
		if !isReady {
			http.Error(w, "Not Ready!", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

  // Serve the /helahtz endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		defer mu.RUnlock()
		if !isHealthy {
			http.Error(w, "Not Healthy!", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

  // Serve the prometheus /metrics endpoint
  http.Handle("/metrics", promhttp.Handler())

	go func() {
    if err := http.ListenAndServe(fmt.Sprintf(":" + serverPort), nil); err != nil {
			klog.Fatalf("Health check server failed: %v", err)
		}
	}()
}

func setReadyz(state bool) {
	mu.Lock()
	defer mu.Unlock()
	isReady = state
}

func setHealhtz(state bool) {
	mu.Lock()
	defer mu.Unlock()
	isHealthy = state
}

func main() {

	// Create the root command for Cobra
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run CRD controller",
		Run: func(cmd *cobra.Command, args []string) {

      // Initialize flags / parameters
			kubernetesConfig, _ := cmd.Flags().GetString("kubernetesConfig")
      if kubernetesConfig == "" { kubernetesConfig = viper.GetString("KUBERNETES_CONFIG") }

			kubernetesAddress, _ := cmd.Flags().GetString("kubernetesAddress")
      if kubernetesAddress == "" { kubernetesAddress = viper.GetString("KUBERNETES_ADDRESS") }

      sourceConfiguration, _  := cmd.Flags().GetString("sourceConfiguration")
      if sourceConfiguration == "" { sourceConfiguration = viper.GetString("SOURCE_CONFIGURATION") } 

      namespaceFilter, _ := cmd.Flags().GetString("namespaceFilter") 
      if namespaceFilter == "" { namespaceFilter = viper.GetString("NAMESPACE_FILTER") }

      controllerName, _ := cmd.Flags().GetString("controllerName")
      if controllerName == "" { controllerName  = viper.GetString("CONTROLLER_NAME") }

      metricsServerPort, _ := cmd.Flags().GetString("metricsServerPort")
      if controllerName == "" { controllerName  = viper.GetString("METRICS_SERVER_PORT") }

			// Initialize klog
			klog.InitFlags(nil)

			// Start the health check server
			startHealthCheckServer(metricsServerPort)

			// Setup signal handler and context
			ctx := signals.SetupSignalHandler()
			logger := klog.FromContext(ctx)

			// Build the controller (configure)
			controllerBuilder := controller.NewControllerBuilder().
				SetWorkingContext(ctx).
				SetWorkingWorkers(2).
				SetControllerName(controllerName).
				SetKubernetesConfig(kubernetesConfig).
				SetKubernetesAddress(kubernetesAddress).
				SetNamespaceFilter(namespaceFilter).
				SetSourceConfiguration(sourceConfiguration).
        SetUpdateReadyz(setReadyz).
        SetUpdateHealthz(setHealhtz)

			// Construct the controller
			controllerClient, err := 
        controller.NewControllerDirector(controllerBuilder).ConstructController()

			if err != nil {
				logger.Error(err, "Error during controller build")
				klog.FlushAndExit(klog.ExitFlushTimeout, 1)
			}

			// Run the controller
			err = controllerClient.Run()
			if err != nil {
				logger.Error(err, "Error during controller run")
				klog.FlushAndExit(klog.ExitFlushTimeout, 1)
			}
		},
	}

	// Initialize configuration
	cobra.OnInitialize(initConfig)

  // Flags 
  runCmd.Flags().String(
    "kubernetesConfig",        
    "",                                       
    "Path to the Kubernetes configuration file (optional)",
  )
  runCmd.Flags().String(
    "kubernetesAddress",       
    "",                                       
    "Address of the Kubernetes API server (optional)",
  )
  runCmd.Flags().String(
    "sourceConfiguration", 
    "/opt/kubeforge/sourceConfiguration.yaml", 
    "Path to the source configuration file (defaults to '/opt/kubeforge/sourceConfiguration.yaml')",
  )
  runCmd.Flags().String(
    "namespaceFilter",
    "default",
    "Namespace to monitor (defaults to 'default')",
  )
  runCmd.Flags().String(
    "controllerName",
    "kubeforge",
    "Name of the controller (defaults to 'kubeforge')",
  )
  runCmd.Flags().String(
    "metricsServerPort",
    "8080",
    "Healthz server port (defaults to '8080')",
  )

  var rootCmd = &cobra.Command{Use: "kubeforge"}
  rootCmd.AddCommand(runCmd)
  rootCmd.Execute()
}

// initConfig reads in the environment variables
func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("KUBEFORGE")
}
