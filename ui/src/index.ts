import { app, BrowserWindow } from 'electron';
import { spawn } from 'child_process';
import path from 'path';

declare const MAIN_WINDOW_WEBPACK_ENTRY: string;
declare const MAIN_WINDOW_PRELOAD_WEBPACK_ENTRY: string;

let mainWindow: BrowserWindow;
let goProcess: ReturnType<typeof spawn>;

const createWindow = () => {
  mainWindow = new BrowserWindow({
    width: 1024,
    height: 768,
    webPreferences: {
      contextIsolation: true,
      nodeIntegration: true,
        preload: MAIN_WINDOW_PRELOAD_WEBPACK_ENTRY
    },  
  });

  

  mainWindow.loadURL(MAIN_WINDOW_WEBPACK_ENTRY);

  mainWindow.webContents.openDevTools();

  // Resolve the Go binary path
  const sendStats = (data: any) => {
    mainWindow.webContents.send('system-stats', {
      data: data.toString(),
      status: "success"
    });
  }

  const sendError = () => {
    mainWindow.webContents.send('system-stats', {
      data: {},
      status: "error"
    });
  }

  const goPath =
    process.env.NODE_ENV === 'development'
      ? '../core/sysmetrics' // Development binary path
      : path.join(process.resourcesPath, 'sysmetrics'); // Production binary path in Resources

  goProcess = spawn(goPath);

  goProcess.stdout.on('data', (data: Buffer) => {
    sendStats(data);
  });

  goProcess.on('error', (error) => {
    console.error(`Go error: ${error.message}`);
    sendError();
  });

  goProcess.stderr.on('data', (error: Buffer) => {
    console.error(`Go error: ${error.toString()}`);
    sendError();
  });

  goProcess.on('close', (code) => {
    console.log(`Go process exited with code ${code}`);
  });

  app.on('before-quit', () => {
    if (goProcess) goProcess.kill();
  });

  app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
      app.quit();
    }
  });

  app.on('activate', () => {
    // On OS X it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow();
    }
});
};

app.on('ready', createWindow);


