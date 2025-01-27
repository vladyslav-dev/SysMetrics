# System Metrics

System Metrics is a cross-platform desktop application designed for real-time monitoring of system performance. Built using modern technologies like Electron, React, and Golang, it provides detailed insights into CPU, Memory, Disk, and Network usage with sleek visualizations.

<img width="717" alt="Screenshot 2025-01-27 at 9 55 49 PM" src="https://github.com/user-attachments/assets/d5f3e16a-496d-4727-b88b-a18cd5e84f42" />


---

## Features


- **CPU Monitoring**: Real-time CPU usage percentage and the number of cores.
- **Memory Insights**: Total, used, and free memory in GB, with percentage usage.
- **Disk Statistics**: Total, used, and free disk space in GB.
- **Network Activity**: Real-time upload and download speeds.
- **System Information**: Operating system details, uptime, and hostname.

---

## Tech Stack

- **Frontend**: React + TypeScript
- **Backend**: Golang (using gopsutil for system metrics)
- **Desktop Environment**: Electron

---

## Installation

### Prerequisites

- **Node.js** (>= 16.x)
- **Go** (>= 1.18)
- **Git**

### Clone the Repository

```bash
git clone https://github.com/yourusername/system-metrics.git
cd system-metrics
```

### Setup and Run in Development Mode

1. Navigate to the `src` folder:
   ```bash
   cd ./ui/src
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the application:
   ```bash
   npm run start
   ```

---

## Packaging for Distribution

### macOS

1. Build the macOS package:
   ```bash
   npm run make
   ```
2. The `.dmg` and `.zip` files will be available in the `out/make` directory.

### Windows

1. Build the Windows installer:
   ```bash
   npm run make
   ```
2. The `.exe` file will be available in the `out/make/squirrel.windows` directory.

---
