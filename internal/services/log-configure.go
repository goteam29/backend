package services

import (
	"api-repository/internal/config"
	"fmt"
	"strings"
	"time"
)

func GetServerStartedLogString(time time.Time, port int, name string) string {
	return fmt.Sprintf("SERVER %s started by PORT: %d at the TIME: %v", name, port, time)
}

func GetBeautifulConfigurationString(c *config.MainConfig) string {
	var sb strings.Builder

	sb.WriteString("╔══════════════════════════════════════════════════╗\n")
	sb.WriteString("║                🔧 SERVICES CONFIG                ║\n")
	sb.WriteString("╚══════════════════════════════════════════════════╝\n")
	sb.WriteString(fmt.Sprintf("  ➤ Gateway Port        : %d\n", c.GatewayPort))
	sb.WriteString(fmt.Sprintf("  ➤ User Service Port   : %d\n", c.UserServicePort))
	sb.WriteString(fmt.Sprintf("  ➤ File Service Port   : %d\n", c.FileServicePort))
	sb.WriteString(fmt.Sprintf("  ➤ Text Service Port   : %d\n", c.TextServicePort))
	sb.WriteString("\n")

	sb.WriteString("╔══════════════════════════════════════════════════╗\n")
	sb.WriteString("║               🛢️ POSTGRES CONFIG                ║\n")
	sb.WriteString("╚══════════════════════════════════════════════════╝\n")
	sb.WriteString(fmt.Sprintf("  ➤ Host                : %s\n", c.POSTGRES.Host))
	sb.WriteString(fmt.Sprintf("  ➤ Port                : %d\n", c.POSTGRES.Port))
	sb.WriteString(fmt.Sprintf("  ➤ Username            : %s\n", c.POSTGRES.Username))
	sb.WriteString(fmt.Sprintf("  ➤ Password            : %s\n", c.POSTGRES.Password))
	sb.WriteString(fmt.Sprintf("  ➤ Database            : %s\n", c.POSTGRES.Database))
	sb.WriteString(fmt.Sprintf("  ➤ Min Connections     : %d\n", c.POSTGRES.MinConns))
	sb.WriteString(fmt.Sprintf("  ➤ Max Connections     : %d\n", c.POSTGRES.MaxConns))
	sb.WriteString(fmt.Sprintf("  ➤ Search Schema       : %s\n", c.POSTGRES.SearchSchema))
	sb.WriteString("\n")

	sb.WriteString("╔══════════════════════════════════════════════════╗\n")
	sb.WriteString("║                 🚀 REDIS CONFIG                  ║\n")
	sb.WriteString("╚══════════════════════════════════════════════════╝\n")
	sb.WriteString(fmt.Sprintf("  ➤ Address             : %s\n", c.REDIS.Address))
	sb.WriteString(fmt.Sprintf("  ➤ Password            : %s\n", c.REDIS.Password))
	sb.WriteString(fmt.Sprintf("  ➤ DB                  : %d\n", c.REDIS.DB))
	sb.WriteString(fmt.Sprintf("  ➤ Protocol            : %d\n", c.REDIS.Protocol))
	sb.WriteString("\n")

	sb.WriteString("╔══════════════════════════════════════════════════╗\n")
	sb.WriteString("║                 🪣 MINIO CONFIG                  ║\n")
	sb.WriteString("╚══════════════════════════════════════════════════╝\n")
	sb.WriteString(fmt.Sprintf("  ➤ Endpoint            : %s\n", c.MinIO.VideoEndpoint))
	sb.WriteString(fmt.Sprintf("  ➤ Access Key          : %s\n", c.MinIO.AccessKey))
	sb.WriteString(fmt.Sprintf("  ➤ Secret Key          : %s\n", c.MinIO.SecretKey))
	sb.WriteString(fmt.Sprintf("  ➤ Region              : %s\n", c.MinIO.Region))
	sb.WriteString(fmt.Sprintf("  ➤ Use SSL             : %t\n", c.MinIO.UseSSL))
	sb.WriteString(fmt.Sprintf("  ➤ Force Path Style    : %t\n", c.MinIO.ForcePathStyle))

	return sb.String()
}
