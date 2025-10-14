<h1 style="color: #1bf89cff; font-size: 48px; font-weight: bold;">Linux Based Hosting</h1>


## Kindergarten Registry — React + Go + MongoDB (Deployed with Nginx)

This project is a Kindergarten Registry Management System built using React (frontend), Go (backend), and MongoDB (database).  
It is deployed on an AWS EC2 (Ubuntu t3.micro) instance using Nginx as a reverse proxy.



## Step-by-Step Deployment Guide

### Launch EC2 Instance
1. Create an EC2 instance using **Ubuntu 22.04 (t3.micro)**.  
2.  Connect via SSH:
   ```bash
   ssh -i key.pem ubuntu@<EC2_PUBLIC_IP>
```

###  Clone the Repository
```
cd ~
git clone https://github.com/Nabil720/Task.git
```

---

## Project Structure
```
Task/
├── kindergarten-registry_linux/
│ ├── backend/ # Go server (API)
│ ├── frontend/ # React application
│ ├── docker-compose.yml # MongoDB service
```

---

### Install Required Packages
```
sudo apt update && sudo apt upgrade -y
sudo apt install -y nginx docker.io docker-compose nodejs npm golang-go

# Verify
nginx -v
docker -v
go version
node -v
```
### Set Up MongoDB
```
cd All_Project/kindergarten-registry
nano docker-compose.yml

version: '3.8'

services:
  mongo:
    image: mongo
    container_name: mongo
    restart: unless-stopped
    ports:
      - "27017:27017"
    volumes:
      - mongo_data:/data/db

volumes:
  mongo_data:

# Run MongoDB
docker-compose up -d

```

### Setup and Run Backend and Frontend
```
cd backend
go mod init backend
go mod tidy
go run main.go

cd ../frontend
npm install
npm run build
```

### Configure Nginx
```
# Backup default configuration
sudo mv /etc/nginx/sites-available/default /etc/nginx/sites-available/default.bak

# Create a new config file
sudo nano /etc/nginx/sites-available/kindergarten.conf

server {
    listen 80;
    server_name 34.236.155.125;

    root /home/ubuntu/All_Project/kindergarten-registry/frontend/build;
    index index.html;

    location / {
        try_files $uri /index.html;
    }

    location /api/ {
        proxy_pass http://localhost:5000/;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}


# Enable the configuration and restart Nginx
sudo ln -s /etc/nginx/sites-available/kindergarten.conf /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
sudo systemctl enable nginx
```

### Fix File Permissions
```
sudo chmod -R 755 /home/ubuntu/All_Project/kindergarten-registry/frontend/build
sudo chmod +x /home/ubuntu
sudo chmod +x /home/ubuntu/All_Project
sudo chmod +x /home/ubuntu/All_Project/kindergarten-registry
sudo chmod +x /home/ubuntu/All_Project/kindergarten-registry/frontend
sudo systemctl reload nginx
```

### Access the Application
```
http://34.236.155.125

```



## Applicatio architecture

```
                    +---------------------+
                    |     Web Browser     |
                    | (React Frontend UI) |
                    +---------+-----------+
                              |
                              v
                    +---------------------+
                    |       Nginx         |
                    | (Reverse Proxy)     |
                    +----+----------+-----+
                         |          |
          Static Files   |          |  API Requests
                         v          v
        +-------------------+   +-------------------+
        | React Build Files |   |    Go API Server  |
        | (frontend/build)  |   | (Handles logic &  |
        +-------------------+   |  connects to DB)  |
                                +---------+---------+
                                          |
                                          v
                                +-------------------+
                                |     MongoDB       |
                                |  (Student Records)|
                                +-------------------+

                  Entire stack hosted on:
                  +-------------------------+
                  |      AWS EC2 (Ubuntu)   |
                  | - Docker (MongoDB)      |
                  | - Node.js & Go Runtime  |
                  | - Nginx Web Server      |
                  +-------------------------+



```



![Website View](./Images/Screenshot%20from%202025-10-12%2017-57-05.png)



