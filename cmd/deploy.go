/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a workload to rancher",
	Long:  `Deploy, creating or updating, a workload to a k8s cluster managed by rancher`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deploy called")
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
var VerifySSL bool
var ServiceName string
var ServiceListeningPort string
var ServiceTargetPort string
var ServiceSelectorLabel string
var ServiceSelectorValue string

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&RancherUrl, "rancher-url", "u", "", "rancher server url (required)")
	deployCmd.Flags().StringVarP(&RancherUrl, "rancher-key", "k", "", "rancher server key (required)")
	deployCmd.Flags().StringVarP(&RancherUrl, "rancher-secret", "s", "", "rancher server secret (required)")
	deployCmd.Flags().StringVarP(&RancherUrl, "cluster", "c", "", "rancher cluster name (required)")
	deployCmd.Flags().StringVarP(&RancherUrl, "project", "p", "", "rancher project name (required)")
	deployCmd.Flags().StringVarP(&RancherUrl, "namespace", "n", "default", "kubernetes namespace name")
	deployCmd.Flags().StringVarP(&RancherUrl, "deployment", "d", "", "kubernetes deployment name (required)")
	deployCmd.Flags().StringVarP(&RancherUrl, "image", "i", "", "docker image (required)")

	//deployCmd.Flags().BoolVar(&VerifySSL, "no-ssl-verify", true, "Flag to disable ssl verify on self-signed certs")

	deployCmd.MarkFlagRequired("rancher-url")
	deployCmd.MarkFlagRequired("rancher-key")
	deployCmd.MarkFlagRequired("rancher-secret")
	deployCmd.MarkFlagRequired("cluster")
	deployCmd.MarkFlagRequired("project")
	//deployCmd.MarkFlagRequired("namespace")
	deployCmd.MarkFlagRequired("deployment")
	deployCmd.MarkFlagRequired("image")
}
