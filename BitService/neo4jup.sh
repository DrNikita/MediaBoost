set -e

sudo docker compose down --volumes --remove-orphans

sudo docker system prune -af --volumes

sudo docker compose up -d --build
