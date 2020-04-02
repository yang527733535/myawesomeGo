package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"os/exec"
	"time"
)

func main() {
	GetInitData()
	rsyncCmd:=viper.GetString("rsyncCmd")
	origin_host:=viper.GetString("origin_host")
	origin_path:=viper.GetString("origin_path")
	local_secrets:=viper.GetString("local_secrets")
	local_path:=viper.GetString("local_path")
	codestring := rsyncCmd+" -azvrt --delete --password-file="+local_secrets+"  ecTeamSync@"+origin_host+"::web/"+origin_path+"/ " +local_path

	time.Sleep(time.Second*3)
	fmt.Println(codestring)
	var cmd *exec.Cmd
	cmd = exec.Command("/bin/bash","-C",codestring)
    output ,err:=cmd.CombinedOutput();
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))

}

func GetInitData() {
	workDir, _ := os.Getwd()
	fmt.Println(workDir)
	viper.AddConfigPath(workDir)
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	err:=viper.ReadInConfig()
	if err != nil {
		return
	}
}
