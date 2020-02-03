package cmd

import (
  "fmt"
  "github.com/bukowa/verisignftp/pkg"
  "github.com/spf13/cobra"
  "log"
  "os"
)

var (
  verisignHost string
  verisignLogin string
  verisignPassword string
  verisignZone string

  verisignDownloadPath string

  verisignUnzip bool
  verisignUnzipPath string

  extractDomains bool
  extractDomainsOnly bool
  extractPath string
)

func ExtractDomains(){
  log.Printf("Extracting domains from %v to %v", verisignUnzipPath, extractPath)
  unzippedFile := pkg.FileOpenPanic(verisignUnzipPath)
  extractFile := pkg.FileCreateOrOpenForAppend(extractPath)
  pkg.ExtractDomains(verisignZone, unzippedFile, extractFile)
  log.Print("Done extracting domains!")
}

var rootCmd = &cobra.Command{
  Use:   "verisignftp",
  Short: "Download and unzip from verisign ftp servers",
  Run: func(cmd *cobra.Command, args []string) {

    if extractDomainsOnly == true {
      ExtractDomains()
      return
    }
    downloadPath := pkg.FileCreatePanic(verisignDownloadPath)

    // check if we can create unzip path
    if verisignUnzip == true {
     f := pkg.FileCreatePanic(verisignUnzipPath)
     if err := f.Close(); err != nil {log.Fatal(err)}
    }

    pkg.VerisignDownload(
     verisignLogin,
     verisignPassword,
     verisignHost,
     verisignZone,
     downloadPath)

    if verisignUnzip == true {
      pkg.UnzipFile(pkg.FileOpenPanic(verisignDownloadPath), pkg.FileCreatePanic(verisignUnzipPath))
    }
    if extractDomains == true {
      ExtractDomains()
    }
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  // #TODO user = a, pass = b, downloadpath =c, unzippath = d ...
  rootCmd.PersistentFlags().BoolVarP(&verisignUnzip, "unzip", "x", false, "unzip after downloading?")

  rootCmd.PersistentFlags().StringVarP(&verisignLogin, "user", "u", "", "ftp user")
  rootCmd.PersistentFlags().StringVarP(&verisignPassword, "pass", "p", "", "ftp password")
  if err := rootCmd.MarkPersistentFlagRequired("user"); err != nil {panic(err)}
  if err := rootCmd.MarkPersistentFlagRequired("pass"); err != nil {panic(err)}

  rootCmd.PersistentFlags().StringVarP(&verisignHost, "address", "a", "rz.verisign-grs.com:21", "ftp host with port")
  rootCmd.PersistentFlags().StringVarP(&verisignZone, "zone", "z", "com.zone.gz", "what zone")

  rootCmd.PersistentFlags().StringVarP(&verisignDownloadPath, "downloadpath", "d", "com.zone.gz", "where to download")
  rootCmd.PersistentFlags().StringVarP(&verisignUnzipPath, "unzippath", "i", "com.zone", "where to unzip")
  rootCmd.PersistentFlags().StringVarP(&extractPath, "extractpath", "k", "com.zone.domains", "where to extract domains")

  rootCmd.PersistentFlags().BoolVarP(&extractDomains, "extract", "e", false, "extract domain names after unzipping")
  rootCmd.PersistentFlags().BoolVarP(&extractDomainsOnly, "extractonly", "o", false, "only extract domains from unzipppath")

}
