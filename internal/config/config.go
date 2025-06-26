// internal/config/config.go
package config

import (
    "os"
    "path/filepath"
    "github.com/spf13/viper"
)

func Init() error {
    
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
   
    configDir := filepath.Join(os.Getenv("HOME"), ".goforge")
    
    viper.AddConfigPath(configDir)
    
    viper.SetDefault("directories.home", configDir)
    viper.SetDefault("directories.bin", filepath.Join(configDir, "bin"))
    viper.SetDefault("directories.packages", filepath.Join(configDir, "packages"))
    viper.SetDefault("apt.auto_update", true)
    viper.SetDefault("apt.use_sudo", true)
    viper.SetDefault("logging.level", "info")
    viper.SetDefault("logging.console", true)
    

    dirs := []string{
        configDir,
        filepath.Join(configDir, "bin"),
        filepath.Join(configDir, "packages"),
        filepath.Join(configDir, "logs"),
    }
    
    for _, dir := range dirs {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return err
        }
    }
    
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            return viper.SafeWriteConfig()
        }
        return err
    }

    return nil
}