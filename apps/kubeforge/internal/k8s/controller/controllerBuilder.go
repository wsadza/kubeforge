// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
// Package controller provides a builder pattern to configure 
// and create a Kubernetes controller with options for address, 
// controller name, working context, workers, and namespace filter.
//
// Example usage:
//
// 	builder := controller.NewControllerBuilder().
//    SetControllerName("my-controller").
//    SetWorkingContext(context.Background()).
//    SetWorkingWorkers(5).
//    SetDefaultDefinitionFilePath("/path/to/file").
//    SetKubernetesAddress("https://k8s-cluster.com").
//
// ############################################################

package controller

import "context"

type controllerBuilder struct {
  kubernetesConfig    string          `mandatory:"false"`
  kubernetesAddress   string          `mandatory:"false"`
  controllerName      string          `mandatory:"true"`
  workingContext      context.Context `mandatory:"true"`
	workingWorkers		  int             `mandatory:"true"`
  sourceConfiguration string          `mandatory:"true"`
  namespaceFilter     string          `mandatory:"false"`
	updateReadyz        func(bool)      `mandatory:"true"`
	updateHealthz       func(bool)      `mandatory:"true"`
}
func NewControllerBuilder() *controllerBuilder {
  return &controllerBuilder{}
}
func (controller *controllerBuilder) SetKubernetesConfig(config string) *controllerBuilder {
  controller.kubernetesConfig = config
  return controller
}
func (controller *controllerBuilder) SetKubernetesAddress(address string) *controllerBuilder {
  controller.kubernetesAddress = address
  return controller
}
func (controller *controllerBuilder) SetControllerName(name string) *controllerBuilder {
  controller.controllerName = name
  return controller
}
func (controller *controllerBuilder) SetWorkingContext(ctx context.Context) *controllerBuilder {
  controller.workingContext = ctx
  return controller
}
func (controller *controllerBuilder) SetWorkingWorkers(workersCount int) *controllerBuilder {
  controller.workingWorkers = workersCount 
  return controller
}
func (controller *controllerBuilder) SetSourceConfiguration(sourceConfiguration string) *controllerBuilder {
  controller.sourceConfiguration = sourceConfiguration 
  return controller
}
func (controller *controllerBuilder) SetNamespaceFilter(namespaceFilter string) *controllerBuilder {
  controller.namespaceFilter = namespaceFilter
  return controller
}
func (controller *controllerBuilder) SetUpdateReadyz(updateReadyz func(bool)) *controllerBuilder {
	controller.updateReadyz = updateReadyz
	return controller
}
func (controller *controllerBuilder) SetUpdateHealthz(updateHealthz func(bool)) *controllerBuilder {
	controller.updateHealthz = updateHealthz 
	return controller
}
