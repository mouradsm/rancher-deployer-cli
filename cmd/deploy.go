/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mouradsm/rancher-deployer-cli/deployer"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a workload to rancher",
	Long:  `Deploy, creating or updating, a workload to a k8s cluster managed by rancher`,
	Run: func(cmd *cobra.Command, args []string) {
		deployer.Deploy(IgnoreVerifySSL)
	},
}

//flags

var RancherUrl string
var RancherKey string
var RancherSecret string
var Cluster string
var Project string
var Namespace string
var Deployment string
var Image string
var IgnoreVerifySSL bool
var ServiceName string
var ServiceListeningPort string
var ServiceTargetPort string
var ServiceSelectorLabel string
var ServiceSelectorValue string

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&RancherUrl, "rancher-url", "u", "", "rancher server url (required)")
	deployCmd.Flags().StringVarP(&RancherKey, "rancher-key", "k", "", "rancher server key (required)")
	deployCmd.Flags().StringVarP(&RancherSecret, "rancher-secret", "s", "", "rancher server secret (required)")
	deployCmd.Flags().StringVarP(&Cluster, "cluster", "c", "", "rancher cluster name (required)")
	deployCmd.Flags().StringVarP(&Project, "project", "p", "", "rancher project name (required)")
	deployCmd.Flags().StringVarP(&Namespace, "namespace", "n", "default", "kubernetes namespace name")
	deployCmd.Flags().StringVarP(&RancherUrl, "Deployment", "d", "", "kubernetes deployment name (required)")
	deployCmd.Flags().StringVarP(&RancherUrl, "Image", "i", "", "docker image (required)")

	deployCmd.Flags().BoolVar(&IgnoreVerifySSL, "no-ssl-verify", true, "Flag to disable ssl verify on self-signed certs")

	// deployCmd.MarkFlagRequired("rancher-url")
	// deployCmd.MarkFlagRequired("rancher-key")
	// deployCmd.MarkFlagRequired("rancher-secret")
	// deployCmd.MarkFlagRequired("cluster")
	// deployCmd.MarkFlagRequired("project")
	// //deployCmd.MarkFlagRequired("namespace")
	// deployCmd.MarkFlagRequired("deployment")
	// deployCmd.MarkFlagRequired("image")
}
