package config

import (
    "fmt"
    "os"

    "go.yaml.in/yaml/v4"
)

type Config struct {
    Database          DatabaseConfig    `yaml:"database"`
    Kafka             KafkaConfig       `yaml:"kafka"`
    Redis             RedisConfig       `yaml:"redis"`
    GameSettings      GameSettings      `yaml:"game_settings"`
}

type DatabaseConfig struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    DBName   string `yaml:"name"`
    SSLMode  string `yaml:"ssl_mode"`
}

type KafkaConfig struct {
    Host               string `yaml:"host"`
    Port               int    `yaml:"port"`
    PlayerEventsTopic  string `yaml:"player_events_topic"`
}

type RedisConfig struct {
    Host string `yaml:"host"`
    Port int    `yaml:"port"`
    DB   int    `yaml:"db"`
}

type GameSettings struct {
    KillPoints   int `yaml:"kill_points"`   // default: 1
    DeathPenalty int `yaml:"death_penalty"` // default: -1
}

func LoadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }

    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
    }

    return &config, nil
}